package stack

type Stack struct {
	list []int
}

func (s *Stack) Pop() int {
	res := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return res
}

func (s *Stack) Push(value int) {
	s.list = append(s.list, value)
}

func (s Stack) GetSize() int {
	return len(s.list)
}
