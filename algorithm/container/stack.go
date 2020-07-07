package container

type Stack struct {
	size  int
	items []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(element int) {
	s.size++
	s.items = append(s.items, element)
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		return -1
	}
	s.size--
	first := s.items[s.size]
	s.items = s.items[:s.size]
	return first
}

func (s *Stack) Size() int {
	return s.size
}
