package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("ac"))
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("level"))
	fmt.Println(isPalindrome("eve"))
	fmt.Println(isPalindrome("radr"))
	fmt.Println(isPalindrome("radar"))
	fmt.Println(isPalindrome("noon"))
}

func isPalindrome(s string) bool {
	for i, c := range s {
		j := len(s) - i - 1

		if byte(c) != s[j] {
			return false
		}
	}

	return true
}
