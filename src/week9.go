package main

import "fmt"

func main() {
	fmt.Println(solve9([]int{0, 5, 0, 3, -1}))
	fmt.Println(solve9([]int{3, 0, 3}))
}

func solve9(input []int) []int {
	var cnt int
	for i, n := range input {
		if n == 0 {
			cnt++
		} else {
			input[i - cnt] = n
		}
	}
	for i := len(input) - cnt; i < len(input); i++ {
		input[i] = 0
	}
	return input
}
