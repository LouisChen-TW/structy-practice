// Write a function, pairProduct, that takes in an array and a target product as arguments. The function should return an array containing a pair of indices whose elements multiply to the given target. The indices returned must be unique.

// Be sure to return the indices, not the elements themselves.

// There is guaranteed to be one such pair whose product is the target.

const pairProduct = (numbers, targetProduct) => {
    const store = {}
    for(let i = 0; i < numbers.length; i++) {
        const complement = targetProduct / numbers[i]
        if (complement in store) {
            return [i, store[complement]]
        } else {
            store[numbers[i]] = i
        }
    }
};



pairProduct([3, 2, 5, 4, 1], 8); // -> [1, 3]
pairProduct([3, 2, 5, 4, 1], 10); // -> [1, 2]
pairProduct([4, 7, 9, 2, 5, 1], 5); // -> [4, 5]
pairProduct([4, 7, 9, 2, 5, 1], 35); // -> [1, 4]
console.log(pairProduct([4, 7, 9, 2, 5, 1], 35));