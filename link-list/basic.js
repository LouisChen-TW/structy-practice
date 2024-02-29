class Node {
    constructor(value) {
        this.value = value;
        this.next = null;
    }
}

const a = new Node('a');
const b = new Node('b');
const c = new Node('c');
const d = new Node('d');
a.next = b;
b.next = c;
c.next = d;

// const printLinkList = head => {
//     let node = head;
//     while (node) {
//         console.log(node.value);
//         node = node.next;
//     }
// };

const printLinkList = head => {
    if (!head) return;
    console.log(head.value);
    printLinkList(head.next);
}


printLinkList(a);
