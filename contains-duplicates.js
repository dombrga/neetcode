/*
https://neetcode.io/problems/duplicate-integer, easy
Contains Duplicate
Given an integer array nums, return true if any value appears more than once in the array, 
otherwise return false.
*/

class Solution {
  /**
   * @param {number[]} nums
   * @return {boolean}
   */
  hasDuplicate(nums) {
    for(let i = 0; i < nums.length; i++) {
      const index = nums.indexOf(nums[i]);
      // If index is greater than 0 and not equal to i, then there is a duplicate
      if (index >= 0 && index !== i) {
        return true;
      }
    }
    return false;
  }
  
  /**
   * @param {number[]} nums
   * @return {boolean}
   */
  hasDuplicate2(nums) {
    const numSet = new Set(nums);
    return numSet.size === nums.length ? false : true;
  }

  /**
   * @param {number[]} nums
   * @return {boolean}
   */
  hasDuplicate3(nums) {
    const obj = {};
    for(const n of nums) {
      // if n is already in obj, then there is duplicate
      if(obj[n]) {
        return true;
      }
      obj[n] = 1;
    }
    return false;
  }
}

function assert(got, wanted) {
  return got === wanted ? 'passed' : 'failed';
}

const solution = new Solution()

const testCases = [
  {
    nums: [1, 2, 3, 4],
    wanted: false,
  },
  {
    nums: [1, 2, 3, 2],
    wanted: true,
  },
  {
    nums: [1, 1, 2, 2],
    wanted: true,
  },
  {
    nums: [1],
    wanted: false,
  },
  {
    nums: [],
    wanted: false,
  },
]

testCases.forEach((t, i) => {
  console.log(`hasDuplicate test #${i+1}: ${assert(solution.hasDuplicate(t.nums), t.wanted)}`);
  console.log(`hasDuplicate2 test #${i+1}: ${assert(solution.hasDuplicate2(t.nums), t.wanted)}`);
  console.log(`hasDuplicate3 test #${i+1}: ${assert(solution.hasDuplicate3(t.nums), t.wanted)}`);
  console.log('------------');
})

