// Write a function, isUnivalueList, that takes in the head of a linked list as an argument. The function should return a boolean indicating whether or not the linked list contains exactly one unique value.

// You may assume that the input list is non-empty.


class Node {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}


const isUniValueList = head => {
    let next = head.next
    while (next !== null) {
        if (head.val === next.val) {
            next = next.next;
        } else {
            return false;
        }
    }
    return true
};

const a = new Node(7);
const b = new Node(7);
const c = new Node(4);

a.next = b;
b.next = c;

// 7 -> 7 -> 4

console.log(isUniValueList(a)); // false

const u = new Node(2);
const v = new Node(2);
const w = new Node(2);
const x = new Node(2);
const y = new Node(2);

u.next = v;
v.next = w;
w.next = x;
x.next = y;

// 2 -> 2 -> 2 -> 2 -> 2

console.log(isUniValueList(u)); // true