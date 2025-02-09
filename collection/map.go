package collection

const initMapSize = 16

type Map[K comparable, V any] interface {
	Load(key K) (V, bool)
	LoadOrDefault(key K, defaultValue V) V
	LoadIfAbsent(key K, f func(K) V) V
	LoadAndStore(key K, value V) (V, bool)
	Store(key K, value V)
	StoreAll(m Map[K, V])
	Size() int
	IsEmpty() bool
	Delete(key K)
	LoadAndDelete(key K) (V, bool)
	Clear()
	Range(f func(key K, value V) bool)
}
