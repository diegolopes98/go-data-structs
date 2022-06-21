package stack

type node[T any] interface {
	getvalue() T
	setvalue(value T)
	getnext() node[T]
	setnext(value node[T])
}

type snode[T any] struct {
	value T
	next  node[T]
}

func newnode[T any](value T) node[T] {
	return &snode[T]{value, nil}
}

func (n *snode[T]) setvalue(value T) {
	n.value = value
}

func (n *snode[T]) getvalue() T {
	return n.value
}

func (n *snode[T]) setnext(snode node[T]) {
	n.next = snode
}

func (n *snode[T]) getnext() node[T] {
	return n.next
}
