package collection

const initMapSize = 16

type Map[K comparable, V any] interface {
	Load(key K) (V, bool)
	LoadOrDefault(key K, defaultValue V) V
	LoadIfAbsent(key K, f func(K) V) V
	Store(key K, value V)
	Size() int
	IsEmpty() bool
	Delete(key K)
	LoadAndDelete(key K) (V, bool)
	Clear()
	Range(f func(key K, value V) bool)
}
