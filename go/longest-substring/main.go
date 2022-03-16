package main

func lengthOfLongestSubstring(s string) int {
	var longest, curStart int

	for i, c := range s {
		for j, c2 := range s[curStart:i] {
			if c == c2 {
				if i-curStart > longest {
					longest = i - curStart
				}
				curStart += j + 1
				break
			}
		}
	}

	if len(s)-curStart > longest {
		longest = len(s) - curStart
	}

	return longest
}
