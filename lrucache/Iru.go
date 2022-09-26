package lrucache

type item[T any] struct {
	value T
	uses  int
}
