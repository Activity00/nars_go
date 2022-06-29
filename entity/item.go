package entity

type Item[T any] interface {
	Name() T
}
