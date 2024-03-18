// Write a function, levelAverages, that takes in the root of a binary tree that contains number values. The function should return an array containing the average value of each level.

class Node {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

const levelAverages = root => {
    if (root === null) return [];
    const stack = [{ node: root, level: 0 }];
    const result = [];
    while (stack.length > 0) {
        const cur = stack.pop();
        const curLevel = result[cur?.level] ?? [];
        curLevel.push(cur.node.val);
        result[cur.level] = curLevel;

        if (cur.node.right !== null) stack.push({ node: cur.node.right, level: cur.level + 1 });
        if (cur.node.left !== null) stack.push({ node: cur.node.left, level: cur.level + 1 });
    }
    const levelAverages = []
    for (let sub of result) {
        let total = 0
        for (let item of sub) {
            total += item
        }
        levelAverages.push(total / sub.length)
    }
    return levelAverages;
};

const a = new Node(3);
const b = new Node(11);
const c = new Node(4);
const d = new Node(4);
const e = new Node(-2);
const f = new Node(1);

a.left = b;
a.right = c;
b.left = d;
b.right = e;
c.right = f;

//       3
//    /    \
//   11     4
//  / \      \
// 4   -2     1

console.log(levelAverages(a)); // -> [ 3, 7.5, 1 ]

