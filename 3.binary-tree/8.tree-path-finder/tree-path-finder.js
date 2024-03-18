// Write a function, pathFinder, that takes in the root of a binary tree and a target value. The function should return an array representing a path to the target value. If the target value is not found in the tree, then return null.

// You may assume that the tree contains unique values.

class Node {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

const pathFinderHelper = (root, target) => {
    if (root === null) return null;
    if (root.val === target) return [root.val];
    const left = pathFinderHelper(root.left, target);
    const right = pathFinderHelper(root.right, target);
    if (left !== null) {
        left.push(root.val);
        return left;
    }
    if (right !== null) {
        right.push(root.val);
        return right;
    }
    return null;
};

const pathFinder = (root, target) => {
    const result = pathFinderHelper(root, target);
    if (result === null) {
        return null;
    } else {
        return result.reverse();
    }
};

// not efficient since it create a new array in each call
// const pathFinder = (root, target) => {
//     if (root === null) return null;
//     if (root.val === target) return [root.val];

//     const leftPath = pathFinder(root.left, target);
//     if (leftPath !== null) return [root.val, ...leftPath];

//     const rightPath = pathFinder(root.right, target);
//     if (rightPath !== null) return [root.val, ...rightPath];

//     return null;
// };


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


const root = new Node(0);
let curr = root;
for (let i = 1; i <= 6000; i += 1) {
    curr.right = new Node(i);
    curr = curr.right;
}
//      a
//    /   \
//   b     c
//  / \     \
// d   e     f

console.time("start")
pathFinder(root, 5800); // -> [ 'a', 'b', 'e' ]
console.timeEnd("start")