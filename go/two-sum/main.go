package main

// brute force
// time: o(n^2)
// space: o(n)
func twoSum1(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-num {
				return []int{i, j}
			}
		}
	}

	return nil
}

// one pass hash table
// time: o(n)
// space: o(n)
func twoSum2(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if val, ok := hashmap[complement]; ok {
			return []int{val, i}
		}
		hashmap[num] = i
	}

	return nil
}
