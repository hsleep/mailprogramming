package main

import "fmt"

func main() {
	fmt.Println(solve5([]int{2, 5, 6, 1, 10}, 8))
}

func solve5(inputs []int, target int) (int, int) {
	m := make(map[int]int)
	for i, a := range inputs {
		b := target - a
		if j, ok := m[b]; ok {
			return j, i
		}
		m[a] = i
	}
	return -1, -1
}
