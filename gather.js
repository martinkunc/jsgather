var client, results;


function Gather() {
    GatherClusterOperators();
}

function GatherClusterOperators() {
    clusterOperators = client.ClusterOperators().List(createMetav1_ListOptions())
    results.Report("clusteroperators", JSON.stringify(clusterOperators))
}