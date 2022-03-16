package main

import "fmt"

func main() {
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	// hashmap stores a tuple consisting of:
	// [0]: the string to be replace
	// [1]: the string to replace with
	hashmap := make(map[int][2]string)

	// populate the hashmap
	for i, index := range indices {
		hashmap[index] = [2]string{sources[i], targets[i]}
	}

	// start at the end of s
	for i := len(s); i >= 0; i-- {
		if tuple, ok := hashmap[i]; ok {
			substr := s[i : i+len(tuple[0])]

			if tuple[0] == substr {
				s = s[:i] + tuple[1] + s[i+len(tuple[0]):]
			}
		}
	}

	return s
}
