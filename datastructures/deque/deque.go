// deque.go
package deque

type Deque[T any] struct {
	items []T
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

func (d *Deque[T]) PopFront() T {
	if len(d.items) == 0 {
		panic("Empty Deque")
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item
}

func (d *Deque[T]) PopBack() T {
	if len(d.items) == 0 {
		panic("Empty Deque")
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item
}

func (d *Deque[T]) IsEmpty() bool {
	return len(d.items) == 0
}
