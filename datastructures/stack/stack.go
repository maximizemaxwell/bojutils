// stack.go
package stack

type Stack[T any] struct {
  items []T
}

func NewStack[T any]() *Stack[T] {
  return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
  s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
  if len(s.items) == 0 {
    panic("Empty Stack")
  }
  item := s.items[len(s.items) - 1]
  s.items = s.items[:len(s.items) -1]
  return item
}

func (s *Stack[T]) IsEmpty() bool {
  return len(s.items) == 0
}
