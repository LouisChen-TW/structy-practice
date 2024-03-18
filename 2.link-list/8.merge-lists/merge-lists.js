// Write a function, mergeLists, that takes in the head of two sorted linked lists as arguments. The function should merge the two lists together into single sorted linked list. The function should return the head of the merged linked list.

// Do this in-place, by mutating the original Nodes.

// You may assume that both input lists are non-empty and contain increasing sorted numbers.

class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

const mergeLists = (head1, head2) => {
    let head = head1.val <= head2.val ? head1 : head2
    let tail = head;
    let pointer1 = head.val === head1.val ? head1.next : head2.next
    let pointer2 = pointer1 === head1.next ? head2 : head1

    while (pointer1 !== null && pointer2 !== null) {
        if (pointer1.val <= pointer2.val) {
            tail.next = pointer1;
            pointer1 = pointer1.next;
        } else {
            tail.next = pointer2;
            pointer2 = pointer2.next;
        }
        tail = tail.next;
    }
    if (pointer1 !== null) tail.next = pointer1;
    if (pointer2 !== null) tail.next = pointer2;

    return head;
};

const a = new Node(5);
const b = new Node(7);
const c = new Node(10);
const d = new Node(12);
const e = new Node(20);
const f = new Node(28);
a.next = b;
b.next = c;
c.next = d;
d.next = e;
e.next = f;
// 5 -> 7 -> 10 -> 12 -> 20 -> 28

const q = new Node(6);
const r = new Node(8);
const s = new Node(9);
const t = new Node(25);
q.next = r;
r.next = s;
s.next = t;
// 6 -> 8 -> 9 -> 25

console.log(mergeLists(a, q));
// 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 12 -> 20 -> 25 -> 28
