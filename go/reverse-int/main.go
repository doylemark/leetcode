package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(456))
}

func reverse(x int) int {
	result := 0

	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	return result
}
