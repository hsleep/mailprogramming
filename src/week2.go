package main

import "fmt"

func main() {
	N := 20
	x := 1
	y := 2

	sum := 0
	for y < N {
		if y % 2 == 0 {
			sum += y
		}
		x, y = y, x + y
	}
	fmt.Println(sum)
}
