// Given an array, rotate the array to the right by k steps. K is a natural number
// eg.
// Input:  [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

import { benchmark } from "./benchmarks/speed";

const rotate = (nums: number[], k: number) => {
  for (let i = 0; i < k; i++) {
    const removed = nums.splice(nums.length - 1, 1);
    nums.unshift(...removed);
  }

  return nums;
};

benchmark(() => rotate([7, 1, 2, 3, 4, 5, 6], 5));
