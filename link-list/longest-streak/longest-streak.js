// Write a function, longestStreak, that takes in the head of a linked list as an argument. The function should return the length of the longest consecutive streak of the same value within the list.

class Node {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

const longestStreak = head => {
    if(head === null) return 0
    let maxStreak = 1
    let curStreak = 1
    let curValue = head.val
    let pointer = head.next
    while (pointer !== null) {
        if (curValue === pointer.val) {
            curStreak++
        } else {
            curStreak = 1
            curValue = pointer.val
        }
        maxStreak = curStreak > maxStreak ? curStreak : maxStreak;
        pointer = pointer.next
    }
    return  maxStreak
};

const a = new Node(5);
const b = new Node(5);
const c = new Node(7);
const d = new Node(7);
const e = new Node(7);
const f = new Node(6);

a.next = b;
b.next = c;
c.next = d;
d.next = e;
e.next = f;

console.log(longestStreak(a)); // 3