package main

import "fmt"

func main() {
	var ans bool
	ans = solve4(12345)
	fmt.Println(ans)
	ans = solve4(-101)
	fmt.Println(ans)
	ans = solve4(11111)
	fmt.Println(ans)
	ans = solve4(12421)
	fmt.Println(ans)
	ans = solve4(1221)
	fmt.Println(ans)
	ans = solve4(1251)
	fmt.Println(ans)
	ans = solve4(1111)
	fmt.Println(ans)
}

func solve4(n int) bool {
	var (
		q int
		r int
	)
	q = n
	r = 0
	for q > r {
		nextQ := q / 10
		r = r * 10 + q % 10
		if q == r || nextQ == r {
			return true
		}
		q = nextQ
	}
	return false
}
