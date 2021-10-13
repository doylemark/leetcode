// Given a non-empty array of integers nums every element appears twice except for one
// Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

import { benchmark } from "./benchmarks/benchmark";

const singleNumber = (nums: number[]) => {
  for (let i = 0; i < nums.length; i++) {
    const curr = nums[i];
    nums.splice(i, 1);

    if (!nums.includes(curr)) {
      return curr;
    }
  }

  return undefined;
};

benchmark(() => singleNumber([1, 0, 1]));
