package main

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)

	i, j, res := 3, len(nums)-1, math.MaxInt32

	for k := 0; k < 4; k++ {
		if nums[j-k]-nums[i-k] < res {
			res = nums[j-k] - nums[i-k]
		}
	}

	return res
}
