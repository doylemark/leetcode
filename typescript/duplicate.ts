// remove duplicates from sorted integer array
// cannot create another array in memory, must sort in place
// return length of array

import { benchmark } from "./benchmarks/benchmark";

const removeDuplicates = (nums: number[]) => {
  for (let i = 0; i < nums.length; i++) {
    if (nums[i + 1] === nums[i]) {
      while (nums[i + 1] === nums[i]) {
        nums.splice(i + 1, 1);
      }
    }
  }

  return nums.length;
};

benchmark(() => removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]));
