package calculation

// Stack наш стек для работы с постфиксами
type Stack struct {
	Items []float64
}

// Push пушит числа в стек
func (s *Stack) Push(item float64) {
	s.Items = append(s.Items, item)
}

// Pop забирает (буквально) элемент со стека
func (s *Stack) Pop() float64 {
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}