// Write a function, depthFirstValues, that takes in the root of a binary tree. The function should return an array containing all values of the tree in depth-first order.

// Hey. This is our first binary tree problem, so be extra sure to check out the approach video! -AZ

class Node {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

// const depthFirstValues = root => {
//     if (root === null) return []
//     const result = [];
//     const stack = [root];
//     while (stack.length > 0) {
//         const cur = stack.pop();
//         result.push(cur.val);

//         if (cur.right) stack.push(cur.right);
//         if (cur.left) stack.push(cur.left);
//     }
//     return result;
// };

const depthFirstValues = root => {
    if (root === null) return [];
    const leftValues = depthFirstValues(root.left)
    const rightValues = depthFirstValues(root.right)
    return [root.val, ...leftValues, ...rightValues]
};

const a = new Node('a');
const b = new Node('b');
const c = new Node('c');
const d = new Node('d');
const e = new Node('e');
const f = new Node('f');

a.left = b;
a.right = c;
b.left = d;
b.right = e;
c.right = f;

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f
console.log(depthFirstValues(a));
//    -> ['a', 'b', 'd', 'e', 'c', 'f']
