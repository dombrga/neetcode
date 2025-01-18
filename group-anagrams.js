/*
https://neetcode.io/problems/anagram-groups
Group Anagrams
Given an array of strings strs, group all anagrams together into sublists.
You may return the output in any order.

An anagram is a string that contains the exact same characters as another string,
but the order of the characters can be different.
*/

class Solution {
  /**
   * @param {string[]} strs
   * @return {string[][]}
   */
  groupAnagrams(strs) {

  }

  /**
 * @param {string[]} strs
 * @return {string[][]}
 */
  groupAnagrams2(strs) {
    const res = new Map();

    for (const s of strs) {
      // Create a count array of size 26 (for 'a' to 'z') initialized with 0
      const count = new Array(26).fill(0);

      // Count the occurrences of each character in the string
      for (const c of s) {
        count[c.charCodeAt(0) - "a".charCodeAt(0)]++;
      }

      // Use the count array as a key (converted to a string)
      const key = count.join("#"); // Join with a delimiter to avoid collisions
      console.log('key', key);
      if (res.has(key)) {
        res.get(key).push(s);
      } else {
        res.set(key, [s]);
      }
    }

    // Return all the grouped anagrams
    return Array.from(res.values());
  }
}

// Input: strs = ["act","pots","tops","cat","stop","hat"]
// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

const solution = new Solution()
console.log(solution.groupAnagrams2(["act","pots","tops","cat","stop","hat"]))
