// Write a function, mostFrequentChar, that takes in a string as an argument. The function should return the most frequent character of the string. If there are ties, return the character that appears earlier in the string.

// You can assume that the input string is non-empty.

const mostFrequentChar = s => {
    const hash = {};
    let mostFrequentChar = '';

    for (let char of s) {
        if (char in hash) {
            hash[char]++;
        } else {
            hash[char] = 1;
        }
    }
    for (let char in hash) {
        if (mostFrequentChar === '' || hash[char] > hash[mostFrequentChar]) {
            mostFrequentChar = char;
        }
    }
    return mostFrequentChar;
};

mostFrequentChar('bookeeper'); // -> 'e'
mostFrequentChar('david'); // -> 'd'
mostFrequentChar('abby'); // -> 'b'
mostFrequentChar('mississippi'); // -> 'i'

console.log(mostFrequentChar('bookeeper'));
