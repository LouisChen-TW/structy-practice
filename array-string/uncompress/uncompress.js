// Write a function, uncompress, that takes in a string as an argument. The input string will be formatted into multiple groups according to the following pattern:

// <number><char>

// for example, '2c' or '3a'.
// The function should return an uncompressed version of the string where each 'char' of a group is repeated 'number' times consecutively. You may assume that the input string is well-formed according to the previously mentioned pattern.

// uncompress("2c3a1t"); // -> 'ccaaat'
// uncompress("4s2b"); // -> 'ssssbb'

const uncompress = s => {
    let result = [];
    let jPointer = 0;
    let iPointer = 0;

    for (let i = 0; i <= s.length - 1; i++) {
        if (!isNaN(Number(s[jPointer]))) {
            jPointer++;
        } else {
            let number = Number(s.slice(iPointer, jPointer));
            for (let i = 0; i < number; i++) {
                result.push(s[jPointer])
            }
            jPointer++;
            iPointer = jPointer;
        }
    }
    return result.join('');
};

console.log(uncompress('4s2b'));
