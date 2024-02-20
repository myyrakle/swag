package array

type Array[T any] []T

func (a Array[T]) IsEmpty() bool {
	return len(a) == 0
}

func (a Array[T]) IsNotEmpty() bool {
	return !a.IsEmpty()
}
