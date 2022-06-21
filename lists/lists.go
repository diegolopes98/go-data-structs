package lists

type List[T any] interface {
	Head() (T, error)
	Tail() (T, error)
	Length() uint
	Push(value T) List[T]
	Pop() (T, error)
	Shift() (T, error)
	Unshift(value T) List[T]
	Get(index uint) (T, error)
	Set(index uint, value T)
	Insert(index uint, value T) List[T]
	Remove(index uint) (T, error)
	Reverse()
}
