package stack

type Stack struct {
	contents [5]int
	index    int
}

func (s *Stack) Push(v int) bool {
	if s.index == len(s.contents) {
		return false
	} else {
		s.contents[s.index] = v
		s.index += 1
		return true
	}
}

func (s *Stack) Pop() int {
	s.index -= 1
	i := s.contents[s.index]
	s.contents[s.index] = 0
	return i
}
