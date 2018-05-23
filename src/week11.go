package main

import "fmt"

func main() {
	fmt.Println(solve11("EGG", "FOO"))
	fmt.Println(solve11("ABBCD", "APPLE"))
	fmt.Println(solve11("AAB", "FOO"))
}

func solve11(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	encMap := make(map[byte]byte)

	for i := range a {
		ca := a[i]
		cb := b[i]
		if c, ok := encMap[ca]; ok {
			if cb != c {
				return false
			}
		} else {
			encMap[ca] = cb
		}
	}
	return true
}
