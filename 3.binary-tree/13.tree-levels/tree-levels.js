// Write a function, allTreePaths, that takes in the root of a binary tree. The function should return a 2-Dimensional array where each subarray represents a root-to-leaf path in the tree.

// The order within an individual path must start at the root and end at the leaf, but the relative order among paths in the outer array does not matter.

// You may assume that the input tree is non-empty.

class Node {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

// iterative
const treeLevels = root => {
    if (root === null) return []
    const stack = [{node: root, level: 0}]
    const result = []
    while (stack.length > 0) {
        const cur = stack.pop()
        const curLevel = result[cur?.level] ?? []
        curLevel.push(cur.node.val)
        result[cur.level] = curLevel
        
        if (cur.node.right !== null) stack.push({ node: cur.node.right, level: cur.level + 1 });
        if (cur.node.left !== null) stack.push({node: cur.node.left, level: cur.level + 1})
    }
    return result
};

// recursive
// const treeLevels = root => {
//     const levels = []
//     fillLevels(root, levels, 0)
//     return levels
// };

// const fillLevels = (root, levels, levelNum) => {
//     if (root === null) return [];

//     if (levels.length === levelNum) {
//         levels[levelNum] = [root.val];
//     } else {
//         levels[levelNum].push(root.val);
//     }

//     fillLevels(root.left, levels, levelNum + 1)
//     fillLevels(root.right, levels, levelNum + 1);
// }


const a = new Node("a");
const b = new Node("b");
const c = new Node("c");
const d = new Node("d");
const e = new Node("e");
const f = new Node("f");

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

console.log(treeLevels(a)); // ->
// [
//   ['a'],
//   ['b', 'c'],
//   ['d', 'e', 'f']
// ]