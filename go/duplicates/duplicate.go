package main

import "fmt"

func main() {
	// fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})) // 2, 3
	fmt.Println(findDuplicatesFaster([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

// func findDuplicates(nums []int) []int {
// 	result := []int{}
// 	hashmap := make(map[int]byte)

// 	for _, num := range nums {
// 		if val, ok := hashmap[num]; ok {
// 			if val == 0 {
// 				result = append(result, num)
// 				hashmap[num] = byte(1)
// 			}
// 		} else {
// 			hashmap[num] = 0
// 		}
// 	}

// 	return result
// }

func findDuplicatesFaster(nums []int) []int {
	result := []int{}

	for _, num := range nums {
		index := abs(num) - 1
		fmt.Println(nums[index])
		if nums[index] < 0 {
			result = append(result, index+1)
		} else {
			nums[index] = nums[index] * -1
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
