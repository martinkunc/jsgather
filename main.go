package main

import (
	"fmt"
	"io/ioutil"
	"os"

	configv1client "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err)
	}

	configClient, err := configv1client.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	vm := otto.New()
	err = Gather(vm, configClient)
	if err != nil {
		fmt.Println("Error", err)
	}
}

type Results struct {
	results map[string]string
}

func (r *Results) Report(name, data string) {
	if r.results == nil {
		r.results = map[string]string{}
	}
	r.results[name] = data
}

func Gather(vm *otto.Otto, client *configv1client.ConfigV1Client) error {

	dat, err := ioutil.ReadFile("gather.js")
	if err != nil {
		return err
	}
	results := &Results{}
	_, err = vm.Run(string(dat))
	if err != nil {
		return err
	}
	vm.Set("client", client)
	vm.Set("results", results)

	vm.Set("createMetav1_ListOptions", func(call otto.FunctionCall) otto.Value {
		result, _ := vm.ToValue(metav1.ListOptions{})
		return result
	})
	_, err = vm.Run("Gather()")
	if err != nil {
		return err
	}

	for n, d := range results.results {
		fmt.Printf("Collected result %s with data %s \n \n", n, d)
	}
	return nil
}
