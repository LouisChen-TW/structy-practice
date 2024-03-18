// Write a function, sumList, that takes in the head of a linked list containing numbers as an argument. The function should return the total sum of all values in the linked list.

class Node {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

const sumList = head => {
    let sum = 0
    while(head !== null) {
        sum += head.val;
        head = head.next;
    }
    return sum
};


const x = new Node(38);
const y = new Node(4);

x.next = y;

// 38 -> 4

console.log(sumList(x)); // 42