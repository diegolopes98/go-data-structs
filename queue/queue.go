package queue

import (
	"errors"
)

const (
	emptyerr = "empty queue"
)

type node[T any] struct {
	value T
	next  *node[T]
}

func newnode[T any](value T) *node[T] {
	return &node[T]{value, nil}
}

type Queue[T any] interface {
	Size() uint
	Enqueue(value T) Queue[T]
	Dequeue() (T, error)
}

type queue[T any] struct {
	first *node[T]
	last  *node[T]
	size  uint
}

func New[T any]() Queue[T] {
	return &queue[T]{nil, nil, 0}
}

func (q *queue[T]) Size() uint {
	return q.size
}

func (q *queue[T]) Enqueue(value T) Queue[T] {
	node := newnode(value)
	if q.size == 0 {
		q.first = node
		q.last = node
	} else {
		q.last.next = node
		q.last = node
	}
	q.size++
	return q
}

func (q *queue[T]) Dequeue() (T, error) {
	if q.size == 0 {
		return *new(T), errors.New(emptyerr)
	}
	dequeued := q.first
	q.first = q.first.next
	q.size--
	if q.size == 0 {
		q.last = nil
	}
	return dequeued.value, nil
}
