// queue.go
package queue

type Queue[T any] struct {
  items []T
}

func NewQueue[T any]() *Queue[T] {
  return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
  q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
  if len(q.items) == 0 {
    panic("Empty Queue")
  }
  item := q.items[0]
  q.items = q.items[1:]
  return item
}

func (q *Queue[T]) IsEmpty() bool {
  return len(q.items) == 0
}