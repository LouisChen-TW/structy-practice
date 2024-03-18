// Write a function, hasPath, that takes in an object representing the adjacency list of a directed acyclic graph and two nodes (src, dst). The function should return a boolean indicating whether or not there exists a directed path between the source and destination nodes.

// Hey. This is our first graph problem, so you should be liberal with watching the Approach and Walkthrough. Be productive, not stubborn. -AZ

// iterative
const hasPath = (graph, src, dst) => {
    const stack = [src]
    while(stack.length > 0) {
        const cur = stack.pop()
        if (cur === dst) return true
        for (let neighbor of graph[cur]) {
            stack.push(neighbor)
        }
    }
    return false
};

// recursive
// const hasPath = (graph, src, dst) => {
//     if (src === dst) return true;

//     for (let neighbor of graph[src]) {
//         if (hasPath(graph, neighbor, dst) === true) {
//             return true;
//         }
//     }
//     return false;
// };

const graph = {
    f: ['g', 'i'],
    g: ['h'],
    h: [],
    i: ['g', 'k'],
    j: ['i'],
    k: [],
};

console.log(hasPath(graph, 'f', 'k')); // true
