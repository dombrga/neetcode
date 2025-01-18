/*
https://neetcode.io/problems/two-integer-sum
Two Sum
Given an array of integers nums and an integer target, 
return the indices i and j such that nums[i] + nums[j] == target and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first.
*/

class Solution {
  /**
   * @param {number[]} nums
   * @param {number} target
   * @return {number[]}
   */
  twoSum(nums, target) {
    const indices = [];
    nums.forEach((firstNum, i) => {
      const secondNumIndex = nums.findIndex(n => {
        return (n + firstNum) === target
      })

      if(secondNumIndex !== -1) {
        indices[0] = i < secondNumIndex ? i : secondNumIndex;
        // indices[1] = i > secondNumIndex ? i : secondNumIndex;
        indices[1] = i < secondNumIndex ? secondNumIndex : i;
        // console.log('secondnum', secondNumIndex, indices)
        return;
      }
    })
    
    console.log('twosum', indices)
    return indices;
  }

  /**
   * @param {number[]} nums
   * @param {number} target
   * @return {number[]}
   */
  twoSum2(nums, target) {
    const indices = [];
    nums.forEach((firstNum, i) => {
      const secondNum = target - firstNum;

      const secondNumIndex = nums.findIndex((n, i2) => {
        return n === secondNum && i !== i2;
      })

      if(secondNumIndex !== -1) {
        indices[0] = i < secondNumIndex ? i : secondNumIndex;
        indices[1] = i < secondNumIndex ? secondNumIndex : i;
        return;
      }

    })

    console.log('twosum2', indices);
    return indices;
  
  }
  /**
   * @param {number[]} nums
   * @param {number} target
   * @return {number[]}
   */
  twoSum3(nums, target) {
    const indices = [];
    const diffs = {};
    nums.forEach((n, i) => {
      const secondNum = target - n;
      if(diffs[secondNum]) {
        indices[0] = i;
        indices[1] = diffs[secondNum];
        return;
      }

      diffs[secondNum] = i;
    })

    console.log('twosum3', indices);
    return indices;
  }
}

function assert(got, wanted) {
  return (got[0] === wanted[0]) && (got[1] === wanted[1]) ? 'passed' : 'failed';
}
const solution = new Solution()

const testCases = [
  {
    nums: [3, 4, 5, 6],
    target: 7,
    wanted: [0, 1],
  },
  {
    nums: [4, 5, 6],
    target: 10,
    wanted: [0, 2],
  },
  {
    nums: [5, 5],
    target: 10,
    wanted: [0, 1],
  },
]

testCases.forEach((t, i) => {
  console.log(`two-sum test #${i+1}: ${assert(solution.twoSum(t.nums, t.target), t.wanted)}`);
  console.log(`two-sum2 test #${i+1}: ${assert(solution.twoSum2(t.nums, t.target), t.wanted)}`);
  console.log(`two-sum3 test #${i+1}: ${assert(solution.twoSum2(t.nums, t.target), t.wanted)}`);
  // console.log(`hasDuplicate2 test #${i+1}: ${solution.assert(solution.hasDuplicate2(t.nums), t.wanted)}`);
  // console.log(`hasDuplicate3 test #${i+1}: ${solution.assert(solution.hasDuplicate3(t.nums), t.wanted)}`);
  console.log('------------');
})

