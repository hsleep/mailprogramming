package main

import "fmt"

func main() {
	fmt.Println(solve12([]int{1, 2, 3, 4, 5}))
}

func solve12(input []int) []int {
	var rtn []int
	rtn = make([]int, len(input))
	buff := 1
	for i, n := range input {
		rtn[i] = buff
		buff *= n
	}
	buff = 1
	for i := len(input) - 1; i >= 0; i-- {
		rtn[i] *= buff
		buff *= input[i]
	}
	return rtn
}
