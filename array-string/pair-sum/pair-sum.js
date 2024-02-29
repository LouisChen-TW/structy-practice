// Write a function, pairSum, that takes in an array and a target sum as arguments. The function should return an array containing a pair of indices whose elements sum to the given target. The indices returned must be unique.

// Be sure to return the indices, not the elements themselves.

// There is guaranteed to be one such pair that sums to the target.


const pairSum = (numbers, targetSum) => {
    const store = {}

    for (let i = 0; i < numbers.length; i++) {
        const complement = targetSum - numbers[i]
        if (complement in store) {
            return [i, store[complement]]
        } else {
            store[numbers[i]] = i
        }
    }
};

pairSum([3, 2, 5, 4, 1], 8); // -> [0, 2]
pairSum([4, 7, 9, 2, 5, 1], 5); // -> [0, 5]
pairSum([4, 7, 9, 2, 5, 1], 3); // -> [3, 5]

console.log(pairSum([3, 2, 5, 4, 1], 8));