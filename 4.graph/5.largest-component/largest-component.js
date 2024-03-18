// Write a function, largestComponent, that takes in the adjacency list of an undirected graph. The function should return the size of the largest connected component in the graph.

// iterative
// const largestComponent = graph => {
//     let max = 0
//     const visited = new Set();
//     for (let key in graph) {
//         if (visited.has(key)) continue;
//         const stack = [key];
//         let componentLength = 0
//         while (stack.length > 0) {
//             const cur = stack.pop();
//             if (visited.has(cur)) continue;
//             visited.add(cur);
//             componentLength++
//             for (let node of graph[cur]) {
//                 stack.push(node);
//             }
//         }
//         if (componentLength > max) max = componentLength
//     }
//     return max
// };

// recursive
const largestComponent = graph => {
    const visited = new Set();
    let max = 0;
    for (let node in graph) {
        const componentLength = explore(graph, node, visited);
        if (componentLength > max) max = componentLength
    }
    return max;
};

const explore = (graph, current, visited) => {
    if (visited.has(current)) return 0;

    visited.add(current);
    let size = 1;

    for (let neighbor of graph[current]) {
        size += explore(graph, neighbor, visited);
    }

    return size;
};


console.log(
    largestComponent({
        0: ['8', '1', '5'],
        1: ['0'],
        5: ['0', '8'],
        8: ['0', '5'],
        2: ['3', '4'],
        3: ['2', '4'],
        4: ['3', '2'],
    })
); // -> 4