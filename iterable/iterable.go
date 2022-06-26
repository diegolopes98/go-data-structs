package iterable

type Iterable[T any] interface {
	ForEach(f func(*T))
}
