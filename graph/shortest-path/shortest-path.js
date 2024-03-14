// Write a function, shortestPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return the length of the shortest path between A and B. Consider the length as the number of edges in the path, not the number of nodes. If there is no path between A and B, then return -1.

const shortestPath = (edges, nodeA, nodeB) => {
    const visited = new Set();
    const graph = buildGraph(edges);
    const queue = [{ node: nodeA, distance: 0 }];

    while (queue.length > 0) {
        const cur = queue.shift();
        if (visited.has(cur.node)) continue;
        if (cur.node === nodeB) return cur.distance;
        visited.add(cur.node);

        for (let node of graph[cur.node]) {
            queue.push({ node, distance: cur.distance + 1});
        }
    }
    return - 1;
};

const buildGraph = edges => {
    const graph = {};
    for (const [src, dst] of edges) {
        if (!graph[src]) {
            graph[src] = [];
        }
        if (!graph[dst]) {
            graph[dst] = [];
        }
        graph[src].push(dst);
        graph[dst].push(src);
    }
    return graph;
};

const edges = [
    ['w', 'x'],
    ['x', 'y'],
    ['z', 'y'],
    ['z', 'v'],
    ['w', 'v'],
];

console.log(shortestPath(edges, 'w', 'z')); // -> 2
