class Gatherer {
    client;
    results;

    constructor() { }

    Gather() {
        this.GatherClusterOperators()
        this.GatherProject()
    }

    GatherProject() {
        this.client.apis['project.openshift.io'].v1.project.get().then((response) => {
            this.results.report("projects", JSON.stringify(response.body))
        });
    }

    GatherClusterOperators() {
        this.client.apis['config.openshift.io'].v1.clusteroperators.get().then((response) => {
            this.results.report("clusteroperators", JSON.stringify(response.body))
        });
    }

}