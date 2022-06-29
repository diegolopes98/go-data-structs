package bst

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Value     T
	Frequency uint
	Left      *Node[T]
	Right     *Node[T]
}

func NewNode[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{value, 1, nil, nil}
}

type BST[T constraints.Ordered] interface {
	Root() *Node[T]
	Insert(value T) BST[T]
}

func New[T constraints.Ordered]() BST[T] {
	return &tree[T]{nil}
}

type tree[T constraints.Ordered] struct {
	root *Node[T]
}

func (t *tree[T]) Root() *Node[T] {
	return t.root
}

func insert[T constraints.Ordered](root, node *Node[T]) {
	if node.Value > root.Value {
		if root.Right != nil {
			insert(root.Right, node)
		} else {
			root.Right = node
		}
	} else if node.Value < root.Value {
		if root.Left != nil {
			insert(root.Left, node)
		} else {
			root.Left = node
		}
	} else {
		root.Frequency += 1
	}
}

func (t *tree[T]) Insert(value T) BST[T] {
	node := NewNode(value)
	if t.root == nil {
		t.root = node
	} else {
		insert(t.root, node)
	}
	return t
}
