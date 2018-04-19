package main

import (
	"common"
	"fmt"
	"sort"
)

type Range struct {
	Start int
	End   int
}

type RangeList []Range

func (l RangeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l RangeList) Len() int {
	return len(l)
}

func (l RangeList) Less(i, j int) bool {
	if l[i].Start < l[j].Start {
		return true
	} else if l[i].Start == l[j].Start {
		if l[i].End < l[j].End {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(solve6(RangeList{Range{2, 4}, Range{1, 5}, Range{7, 9}}))
	fmt.Println(solve6(RangeList{Range{3, 6}, Range{1, 3}, Range{2, 4}}))
}

func solve6(inputs RangeList) RangeList {
	sort.Sort(inputs)
	rtn := RangeList{}
	var cur *Range
	for _, r := range inputs {
		if cur != nil && r.Start <= cur.End {
			cur.End = util.Max(r.End, cur.End)
		} else {
			rtn = append(rtn, r)
			cur = &rtn[len(rtn)-1]
		}
	}
	return rtn
}
