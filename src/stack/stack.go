package stack

import "fmt"

type Stack struct {
	contents []int
	index    int
}

func (s *Stack) Push(v int) bool {
	s.contents = append(s.contents, v)
	s.index += 1
	return true
}

func (s *Stack) Pop() int {
	s.index -= 1
	i := s.contents[s.index]
	slice := make([]int, len(s.contents)-1)
	copy(slice, s.contents)
	s.contents = slice
	return i
}

func (s *Stack) Size() int {
	return len(s.contents)
}
