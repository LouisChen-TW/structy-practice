// Write a function, leafList, that takes in the root of a binary tree and returns an array containing the values of all leaf nodes in left-to-right order.

class Node {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

// iterative
// const leafList = root => {
//     if (root === null) return [];
//     const list = [];
//     const stack = [root];

//     while (stack.length > 0) {
//         const cur = stack.pop();
//         if (cur.left === null && cur.right === null) {
//             list.push(cur.val);
//         }
//         if (cur.right !== null) {
//             stack.push(cur.right);
//         }
//         if (cur.left !== null) {
//             stack.push(cur.left);
//         }
//     }
//     return list;
// };


// recursive
const fillLeaves = (root, leaves) => {
    if (root === null) return [];
    if (root.left === null && root.right === null) {
        leaves.push(root.val);
    }
    fillLeaves(root.left, leaves);
    fillLeaves(root.right, leaves);
};

const leafList = root => {
    const leaves = [];
    fillLeaves(root, leaves)
    return leaves
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

console.log(leafList(a)); // -> [ 'd', 'e', 'f' ]
