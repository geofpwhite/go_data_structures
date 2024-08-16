package queue

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	s := make(Queue[T], 0)
	return &s
}

func (s *Queue[T]) Peek() (T, int) {
	return (*s)[0], len(*s)
}
func (s *Queue[T]) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Queue[T]) Enqueue(value T) {
	*s = append(*s, value)
}

/*
returns nil if empty
*/
func (s *Queue[T]) Dequeue() *T {
	if s.IsEmpty() {
		return nil
	}
	poppedNode := (*s)[0]
	*s = (*s)[1:]
	return &poppedNode
}
