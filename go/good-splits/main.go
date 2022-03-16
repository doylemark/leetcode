package main

import "fmt"

func main() {
	numSplits("aacaba")
}

func numSplits(s string) int {
	right := make(map[rune]int)
	left := make(map[rune]int)
	ans := 0

	for _, r := range s {
		right[r]++
	}

	fmt.Printf("%#v\n", right)

	for _, r := range s {
		left[r]++
		right[r]--

		if right[r] == 0 {
			delete(right, r)
		}

		if len(left) == len(right) {
			ans++
		}
	}

	return ans
}
