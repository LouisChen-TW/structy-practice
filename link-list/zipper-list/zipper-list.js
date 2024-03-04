// Write a function, zipperLists, that takes in the head of two linked lists as arguments. The function should zipper the two lists together into single linked list by alternating nodes. If one of the linked lists is longer than the other, the resulting list should terminate with the remaining nodes. The function should return the head of the zippered linked list.

// Do this in-place, by mutating the original Nodes.

// You may assume that both input lists are non-empty.

class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

const zipperLists = (head1, head2) => {
    let count = 0;
    const head = head1;
    let tail = head
    let pointer1 = head1.next;
    let pointer2 = head2;
    while (pointer1 !== null && pointer2 !== null) {
        if (count % 2 === 0) {
            tail.next = pointer2;
            pointer2 = pointer2.next;
        } else {
            tail.next = pointer1;
            pointer1 = pointer1.next;
        }
        tail = tail.next;
        count++;
    }
    if (pointer1 !== null) tail.next = pointer1;
    if (pointer2 !== null) tail.next = pointer2;

    return head;
};

const a = new Node('a');
const b = new Node('b');
const c = new Node('c');
a.next = b;
b.next = c;
// a -> b -> c

const x = new Node('x');
const y = new Node('y');
const z = new Node('z');
x.next = y;
y.next = z;
// x -> y -> z

console.log(zipperLists(a, x));
// a -> x -> b -> y -> c -> z
