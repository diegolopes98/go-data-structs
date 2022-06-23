package stack

import "errors"

const (
	emptyerr = "empty stack"
)

type Stack[T any] interface {
	Size() uint
	Push(value T) Stack[T]
	Pop() (T, error)
}

type stack[T any] struct {
	first node[T]
	last  node[T]
	size  uint
}

func New[T any]() Stack[T] {
	return &stack[T]{nil, nil, 0}
}

func (s *stack[T]) Size() uint {
	return s.size
}

func (s *stack[T]) Push(value T) Stack[T] {
	node := newnode(value)
	if s.size == 0 {
		s.first = node
		s.last = node
	} else {
		oldfirst := s.first
		node.setnext(oldfirst)
		s.first = node
	}
	s.size++
	return s
}

func (s *stack[T]) Pop() (T, error) {
	if s.size == 0 {
		return *new(T), errors.New(emptyerr)
	}
	popped := s.first
	s.first = popped.getnext()
	s.size--
	if s.size == 0 {
		s.last = nil
	}
	value := popped.getvalue()
	return value, nil
}
