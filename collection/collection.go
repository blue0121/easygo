package collection

type Collection[E comparable] interface {
	Add(e ...E) bool
	AddAll(c Collection[E]) bool
	Contains(e E) bool
	Size() int
	IsEmpty() bool
	Delete(e E) bool
	Clear()
	Range(func(e E) bool)
}
