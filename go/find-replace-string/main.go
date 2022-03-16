package main

import "fmt"

func main() {
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(findReplaceStringSlow("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	// this map essentially holds: replace [0] with [1]
	// eg [2]string{"a", "eee"} -> replace "a" with "eee"
	hashmap := make(map[int][2]string)

	// populate the hashmap
	for i, val := range indices {
		add := [2]string{sources[i], targets[i]}
		hashmap[val] = add
	}

	// starting from the back of s
	for i := len(s); i >= 0; i-- {
		if tuple, ok := hashmap[i]; ok {
			substr := s[i : i+len(tuple[0])]

			if substr == tuple[0] {
				s = s[:i] + tuple[1] + s[i+len(tuple[0]):]
			}
		}
	}

	return s
}

func findReplaceStringSlow(s string, indices []int, sources []string, targets []string) string {
	hashmap := make(map[int][]string)

	for i := 0; i < len(indices); i++ {
		index := indices[i]
		source := sources[i]
		target := targets[i]

		if s[index:index+len(source)] == source {
			hashmap[index] = []string{source, target}
		}
	}

	output := ""

	for j := 0; j < len(s); {
		if val, ok := hashmap[j]; ok {
			output += val[1]
			j += len(val[0])
		} else {
			output += string(s[j])
			j++
		}
	}

	return output
}
