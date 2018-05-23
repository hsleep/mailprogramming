package main

import (
	"fmt"
)

func main() {
	fmt.Println(solve8([]int{10, 5, 4, 3, -1}))
	fmt.Println(solve8([]int{3, 3, 3}))
}

func solve8(input []int) string {
	var max1, max2 int
	max1 = input[0]
	max2 = max1
	for _, n := range input {
		if max1 < n {
			max2 = max1
			max1 = n
		} else if max1 == max2 || max2 < n {
			max2 = n
		}
	}
	if max1 == max2 {
		return "Does not exist."
	}
	return fmt.Sprint(max2)
}
