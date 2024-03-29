// Write a function, connectedComponentsCount, that takes in the adjacency list of an undirected graph. The function should return the number of connected components within the graph.

// iterative
// const connectedComponentsCount = graph => {
//     let count = 0;
//     const visited = new Set();
//     for (let key in graph) {
//         if (visited.has(Number(key))) continue;
//         const stack = [key];
//         while (stack.length > 0) {
//             const cur = stack.pop();
//             if (visited.has(Number(cur))) continue;
//             visited.add(Number(cur));
//             for (let node of graph[cur]) {
//                 stack.push(node);
//             }
//         }
//         count++
//     }
//     return count;
// };


// recursive
const connectedComponentsCount = graph => {
    const visited = new Set();
    let count = 0;
    for (let node in graph) {
        if (explore(graph, node, visited) === true) {
            count += 1;
        }
    }
    return count;
};

const explore = (graph, current, visited) => {
    if (visited.has(String(current))) return false;

    visited.add(String(current));

    for (let neighbor of graph[current]) {
        explore(graph, neighbor, visited);
    }

    return true;
};

// recursive

console.log(
    connectedComponentsCount({
        0: [8, 1, 5],
        1: [0],
        5: [0, 8],
        8: [0, 5],
        2: [3, 4],
        3: [2, 4],
        4: [3, 2],
    })
); // -> 2
