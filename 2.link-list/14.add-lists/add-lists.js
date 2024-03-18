// Write a function, addLists, that takes in the head of two linked lists, each representing a number. The nodes of the linked lists contain digits as values. The nodes in the input lists are reversed; this means that the least significant digit of the number is the head. The function should return the head of a new linked listed representing the sum of the input lists. The output list should have its digits reversed as well.

// You must do this by traversing through the linked lists once.

class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

const addLists = (head1, head2) => {
    let pointer1 = head1;
    let pointer2 = head2;
    let head = new Node(null);
    let tail = head
    let carry = 0;

    while (pointer1 !== null || pointer2 !== null) {
        const pointer1Val = pointer1?.val ?? 0
        const pointer2Val = pointer2?.val ?? 0;
        const remain = carry
        const sum = pointer1Val + pointer2Val + remain;
        carry = 0
        if (sum > 9) {
            tail.next = new Node(sum - 10);
            carry++
        } else {
            tail.next = new Node(sum);
        }
        tail = tail.next;
        pointer1 = pointer1?.next ?? null
        pointer2 = pointer2?.next ?? null
    }
    if (carry > 0) tail.next = new Node(carry)

    return head.next;
};

//  7541
// +  32
// -----
//  7573

// const a1 = new Node(1);
// const a2 = new Node(4);
// const a3 = new Node(5);
// const a4 = new Node(7);
// a1.next = a2;
// a2.next = a3;
// a3.next = a4;
// // 1 -> 4 -> 5 -> 7

// const b1 = new Node(2);
// const b2 = new Node(3);
// b1.next = b2;
// // 2 -> 3

// console.log(addLists(a1, b1));
// 3 -> 7 -> 5 -> 7

//   999
//  +  6
//  ----
//  1005

const a1 = new Node(9);
const a2 = new Node(9);
const a3 = new Node(9);
a1.next = a2;
a2.next = a3;
// 9 -> 9 -> 9

const b1 = new Node(6);
// 6

console.log(addLists(a1, b1))

// 5 -> 0 -> 0 -> 1