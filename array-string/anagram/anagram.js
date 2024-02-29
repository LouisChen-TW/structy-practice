// Write a function, anagrams, that takes in two strings as arguments. The function should return a boolean indicating whether or not the strings are anagrams. Anagrams are strings that contain the same characters, but in any order.

const anagrams = (s1, s2) => {
    if (s1.length !== s2.length) return false
    const map = {}

    for(let i of s1) {
        if (map[i]) {
            map[i]++;
        } else {
            map[i] = 1;
        }
    }

    for (let i of s2) {
        if (map[i]) {
            map[i]--
        } else {
            return false
        }
    }

    for (let i in map) {
        if(map[i] !== 0) {
            return false
        }
    }
    return true
};


anagrams('restful', 'fluster'); // -> true
anagrams('cats', 'tocs'); // -> false
anagrams('monkeyswrite', 'newyorktimes'); // -> true
console.log(anagrams('monkeyswrite', 'newyorktimes'))