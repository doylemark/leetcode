// Given an integer array nums
// return true if any value appears at least twice in the array
// return false if every element is distinct
// avoid using includes() etc
import { benchmark } from "./benchmarks/benchmark";

// very slow runtime, good mem usage
const containsDuplicate = (nums: number[]) => {
  const reduced = nums.reduce((ctx, curr) => {
    if (!ctx.includes(curr)) {
      ctx.push(curr);
    }
    return ctx;
  }, [] as number[]);

  return nums.length !== reduced.length;
};

benchmark(() => containsDuplicate([1, 2, 3, 1]));

// slightly inferior mem usage, excellent runtime (top 15%)
const containsDuplicate2 = (nums: number[]) => {
  return Array.from(new Set(nums)).length !== nums.length;
};

benchmark(() => containsDuplicate2([1, 2, 3, 1]));
