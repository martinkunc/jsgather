Using openshift nodejs client
https://www.npmjs.com/package/openshift-rest-client

The client has to be updated to add new 4.3 swagger.
mv node_modules/openshift-rest-client/lib/specs/openshift-openapi-spec.json.gz \
  node_modules/openshift-rest-client/lib/specs/openshift-openapi-spec.json.old.gz
cp swagger.json node_modules/openshift-rest-client/lib/specs/openshift-openapi-spec.json
gzip node_modules/openshift-rest-client/lib/specs/openshift-openapi-spec.json 

Running:

JS test in Node
```
$ node test.js > output-js.txt
```

Go
```
KUBECONFIG=~/.crc/machines/crc/kubeconfig go run main.go > output-go.txt
```
