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
	ForEach(f func(*T))
}

func ForEach[T any](l List[T], f func(*T)) {
	l.ForEach(f)
}

func Filter[T any](l List[T], new func() List[T], f func(T) bool) List[T] {
	nl := new()
	l.ForEach(func(value *T) {
		if f(*value) {
			nl.Push(*value)
		}
	})
	return nl
}

func Map[T, N any](l List[T], new func() List[N], f func(T) N) List[N] {
	nl := new()
	l.ForEach(func(value *T) {
		nl.Push(f(*value))
	})
	return nl
}
