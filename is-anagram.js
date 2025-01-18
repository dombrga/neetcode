/*
https://neetcode.io/problems/is-anagram
Valid Anagram
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
*/
class Solution {
  /**
   * @param {string} s
   * @param {string} t
   * @return {boolean}
   */
  isAnagram(s, t) {
    if(s.length !== t.length) {
      return false;
    }

    const firstWordLetters = {};
    const secondWordLetters = {};

    s.split('').forEach(letter => {
      if(firstWordLetters[letter]) {
        firstWordLetters[letter] += 1;
      } else {
        firstWordLetters[letter] = 1;
      }
    })

    t.split('').forEach(letter => {
      if(secondWordLetters[letter]) {
        secondWordLetters[letter] += 1;
      } else {
        secondWordLetters[letter] = 1;
      }
    })

    for(const letter in firstWordLetters) {
      // if letter in firstWordLetters is not in secondWordLetters, not anagram
      if(!secondWordLetters[letter]) {
        return false;
      }

      // if count of letter not same, not anagram
      if(firstWordLetters[letter] !== secondWordLetters[letter]) {
        return false;
      }
    }

    return true;
  }
}

function assert(got, wanted) {
  return got === wanted ? 'passed' : 'failed';
}
const solution = new Solution()

const testCases = [
  {
    firstWord: 'racecar',
    secondWord: 'carrace',
    wanted: true,
  },
  {
    firstWord: 'jam',
    secondWord: 'jar',
    wanted: false,
  },
  {
    firstWord: 'jjj',
    secondWord: 'jjj',
    wanted: true,
  },
]

testCases.forEach((t, i) => {
  console.log(`is-anagram test #${i+1}: ${assert(solution.isAnagram(t.firstWord, t.secondWord), t.wanted)}`);
  // console.log(`hasDuplicate2 test #${i+1}: ${solution.assert(solution.hasDuplicate2(t.nums), t.wanted)}`);
  // console.log(`hasDuplicate3 test #${i+1}: ${solution.assert(solution.hasDuplicate3(t.nums), t.wanted)}`);
  // console.log('------------');
})

