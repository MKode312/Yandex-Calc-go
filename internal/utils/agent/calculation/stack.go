package calculation

type Stack struct {
	Items []float64
}

func (s *Stack) Push(item float64) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() float64 {
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}
