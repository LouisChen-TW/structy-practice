// Write a function, maxValue, that takes in array of numbers as an argument. The function should return the largest number in the array.

// Solve this without using any built-in array methods.

// You can assume that the array is non-empty.

const maxValue = nums => {
    let maxNumber = -Infinity;
    for (let i = 0; i <= nums.length; i++) {
        if (maxNumber < nums[i]) {
            maxNumber = nums[i];
        }
    }
    return maxNumber;
};

maxValue([4, 7, 2, 8, 10, 9]); // -> 10
maxValue([10, 5, 40, 40.3]); // -> 40.3
maxValue([-5, -2, -1, -11]); // -> -1
