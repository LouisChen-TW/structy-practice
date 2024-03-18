// Write a function, insertNode, that takes in the head of a linked list, a value, and an index. The function should insert a new node with the value into the list at the specified index. Consider the head of the linked list as index 0. The function should return the head of the resulting linked list.

// Do this in-place.

// You may assume that the input list is non-empty and the index is not greater than the length of the input list.

class Node {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

const insertNode = (head, value, index) => {
  const newNode = new Node(value);
  if (index <= 0) {
    newNode.next = head;
    return newNode;
  }
  let curIndex = 1;
  let prev = head;
  let cur = head.next;
  while (cur !== null && curIndex <= index) {
    if (curIndex === index) {
      prev.next = newNode;
      newNode.next = cur;
    } else {
      prev = cur;
      cur = cur.next;
    }
    curIndex++;
  }
  if (cur === null) {
    prev.next = newNode
  }

  return head;
};

const a = new Node("a");
const b = new Node("b");
const c = new Node("c");
const d = new Node("d");

a.next = b;
b.next = c;
c.next = d;

// a -> b -> c -> d

console.log(insertNode(a, 'x', -1));
// a -> b -> x -> c -> d
