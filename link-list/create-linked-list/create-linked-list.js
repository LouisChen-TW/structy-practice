// Write a function, createLinkedList, that takes in an array of values as an argument. The function should create a linked list containing each element of the array as the values of the nodes. The function should return the head of the linked list.

class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

const createLinkedList = values => {
    let head = new Node(null);
    let cur = head;
    for (let i = 0; i < values.length; i++) {
        cur.next = new Node(values[i]);
        cur = cur.next;
    }
    return head.next;
};

console.log(createLinkedList(['H', 'E', 'Y']));
// h -> e -> y
