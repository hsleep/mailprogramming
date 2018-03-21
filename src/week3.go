package main

import "fmt"

func main() {
	n := 4
	openCnt, closeCnt := n, n
	var bag []string
	bag = solve3(bag, "", openCnt, closeCnt)
	fmt.Println(bag)
}

func solve3(bag []string, answer string, open, close int) []string{
	if open == 0 && close == 0 {
		bag = append(bag, answer)
	} else {
		if open > 0 {
			bag = solve3(bag, answer + "(", open - 1, close)
		}
		if open < close {
			bag = solve3(bag, answer + ")", open, close - 1)
		}
	}
	return bag
}
