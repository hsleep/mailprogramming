package main

import (
	"common"
	"fmt"
)

func main() {
	var ans int
	ans = solve([]int{-1, 3, -1, 5})
	fmt.Println(ans)
	ans = solve([]int{-5, -3, -1})
	fmt.Println(ans)
	ans = solve([]int{2, 4, -2, -3, 8})
	fmt.Println(ans)
}

func solve(inputs []int) int {
	ans := inputs[0]
	currentSum := inputs[0]
	for _, n := range inputs[1:] {
		currentSum = util.Max(currentSum+n, n)
		if currentSum > ans {
			ans = currentSum
		}
	}
	return ans
}
