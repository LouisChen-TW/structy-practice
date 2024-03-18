// basic graph traverse, use depth and breadth.

const graph = {
    a: ['b', 'c'],
    b: ['d'],
    c: ['e'],
    d: ['f'],
    e: [],
    f: [],
};

// iterative depth
// const depthFirstPrint = (graph, source) => {
//     const stack = [source];
//     while (stack.length > 0) {
//         const cur = stack.pop();
//         console.log(cur);
//         for (let item of graph[cur]) {
//             stack.push(item);
//         }
        
//     }
// };

// recursive depth
const depthFirstPrint = (graph, source) => {
    console.log(source)
    for (let item of graph[source]) {
        depthFirstPrint(graph, item)
    }
};

const breadthFirstPrint = (graph, source) => {
    const queue = [source];
    while (queue.length > 0) {
        const cur = queue.shift();
        console.log(cur);
        for (let item of graph[cur]) {
            queue.push(item);
        }
    }
};

// depthFirstPrint(graph, 'a');
breadthFirstPrint(graph, 'a');
