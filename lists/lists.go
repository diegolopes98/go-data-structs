package lists

type Node[T any] interface {
	SetValue(value T)
	GetValue() T
	SetNext(node Node[T])
	GetNext() Node[T]
	SetPrevious(node Node[T])
	GetPrevious() Node[T]
}

type List[T any] interface {
	Head() Node[T]
	Tail() Node[T]
	Length() uint
	Push(value T) List[T]
	Pop() Node[T]
	Shift() Node[T]
	Unshift(value T) List[T]
	Get(index uint) Node[T]
	Set(index uint, value T)
	Insert(index uint, value T) List[T]
	Remove(index uint)
	Reverse()
}
