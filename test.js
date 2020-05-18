const openshiftRestClient = require('openshift-rest-client').OpenshiftClient;
var fs = require('fs');

createGatherer = function (a,b) {
    let g = eval(fs.readFileSync('./gather.js') + ' new Gatherer()');
    return g
}


class Results {
    report(name, data) {
        console.log("\n\nReturned result \n"+name+" data " + data)
    }
}

const config = '/home/mkunc/.crc/machines/crc/kubeconfig';

results = new Results()


openshiftRestClient({ config }).then((client) => {
  gatherer = createGatherer()
  gatherer.client = client
  gatherer.results = results

  gatherer.Gather();
  
},(r) => {
    console.log("rejected because of "+r)
});

