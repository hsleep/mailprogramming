package main

import "fmt"

func main() {
	fmt.Println(solve10("aabcbcbc"))
	fmt.Println(solve10("aaaaaaaa"))
	fmt.Println(solve10("abbbcedd"))
}

func solve10(s string) int {
	bag := make(map[byte]bool)
	var maxLen int
	for _, b := range []byte(s) {
		if bag[b] {
			if len(bag) > maxLen {
				maxLen = len(bag)
			}
			bag = make(map[byte]bool)
		}
		bag[b] = true
	}
	if len(bag) > maxLen {
		maxLen = len(bag)
	}
	return maxLen
}
