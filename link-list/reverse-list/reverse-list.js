// Write a function, reverseList, that takes in the head of a linked list as an argument. The function should reverse the order of the nodes in the linked list in-place and return the new head of the reversed linked list.

class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

const reverseList = head => {
    let prev = null
    let cur = head
    while (cur !== null) {
        const next = cur.next
        cur.next = prev
        prev = cur
        cur = next
    }
    return prev
};


const a = new Node("a");
const b = new Node("b");
const c = new Node("c");
const d = new Node("d");
const e = new Node("e");
const f = new Node("f");

a.next = b;
b.next = c;
c.next = d;
d.next = e;
e.next = f;

// a -> b -> c -> d -> e -> f

console.log(reverseList(a)); // f -> e -> d -> c -> b -> a