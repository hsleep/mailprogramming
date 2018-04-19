package main

import (
	"bytes"
	"fmt"
)

type Stack struct {
	data []byte
}

func (s *Stack) Push(c byte) {
	s.data = append(s.data, c)
}

func (s *Stack) PopAll() string {
	var b bytes.Buffer
	for i := len(s.data) - 1; i >= 0; i-- {
		b.WriteByte(s.data[i])
	}
	rtn := b.String()
	s.Reset()
	return rtn
}

func (s *Stack) Reset() {
	s.data = make([]byte, 0)
}

func (s *Stack) Len() int {
	return len(s.data)
}

func main() {
	fmt.Println(solve7("abc def ghi"))
}

func solve7(s string) string {
	var rtn bytes.Buffer
	bag := Stack{}
	for _, c := range []byte(s) {
		if c == ' ' {
			rtn.WriteString(bag.PopAll())
			rtn.WriteByte(c)
		} else {
			bag.Push(c)
		}
	}
	if bag.Len() > 0 {
		rtn.WriteString(bag.PopAll())
	}
	return rtn.String()
}
