// Write a function, compress, that takes in a string as an argument. The function should return a compressed version of the string where consecutive occurrences of the same characters are compressed into the number of occurrences followed by the character. Single character occurrences should not be changed.

// 'aaa' compresses to '3a'
// 'cc' compresses to '2c'
// 't' should remain as 't'

const compress = s => {
    let result = ''
    let i = 0
    let j = 0
    while(j <= s.length) {
        if (s[i] === s[j]) {
            j++
        } else {
            let num = j - i
            let addString = num === 1 ? s[i] : `${String(num)}${s[i]}`;
            result += addString
            i = j
        }
    }
    return result
};


compress('ccaaatsss'); // -> '2c3at3s'
compress('ssssbbz'); // -> '4s2bz'
compress('ppoppppp'); // -> '2po5p'

console.log(compress('ppoppppp'))