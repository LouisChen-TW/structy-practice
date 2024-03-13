// Write a function, undirectedPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return a boolean indicating whether or not there exists a path between nodeA and nodeB.


const undirectedPath = (edges, nodeA, nodeB) => {
    const graph = edgesToGraph(edges);
    const newSet = new Set();
    const stack = [nodeA];

    while (stack.length > 0) {
        const cur = stack.pop();
        if (newSet.has(cur)) {
            continue;
        }
        if (cur === nodeB) {
            return true;
        }
        newSet.add(cur);
        for (let node of graph[cur]) {
            stack.push(node);
        }
    }
    return false;
};

const edgesToGraph = edges => {
    const graph = {};
    for (let edge of edges) {
        for (let i = 0; i < edge.length; i++) {
            if (graph[edge[i]]) {
                if (i === 0) {
                    graph[edge[i]].push(edge[1]);
                } else {
                    graph[edge[i]].push(edge[0]);
                }
            } else {
                if (i === 0) {
                    graph[edge[i]] = [edge[1]];
                } else {
                    graph[edge[i]] = [edge[0]];
                }
            }
        }
    }
    return graph;
};


const edges = [
    ['b', 'a'],
    ['c', 'a'],
    ['b', 'c'],
    ['q', 'r'],
    ['q', 's'],
    ['q', 'u'],
    ['q', 't'],
];

console.log(undirectedPath(edges, 'r', 'b'));
// -> true
