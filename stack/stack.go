package stack

type StackNode[T any] struct {
	value T
	next  *StackNode[T]
}

type Stack[T any] struct {
	top *StackNode[T]
}

func NewStack[T any]() *Stack[T] {
	bottom := &StackNode[T]{}
	bottom.next = bottom
	s := Stack[T]{}
	s.top = bottom
	return &s
}

func (s *Stack[T]) Peek() (T, bool) {
	return s.top.value, s.isEmpty()
}
func (s *Stack[T]) isEmpty() bool {
	return s.top.next == s.top
}
func (s *Stack[T]) Push(value T) {
	newNode := StackNode[T]{value: value, next: s.top}
	s.top = &newNode
}

/*

 */
func (s *Stack[T]) Pop() (T, bool) {
	poppedNode := s.top
	s.top = s.top.next
	return poppedNode.value, s.isEmpty()
}
