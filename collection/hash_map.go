package collection

type hashMap[K comparable, V any] struct {
	initSize int
	m        map[K]V
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &hashMap[K, V]{
		initSize: initMapSize,
		m:        make(map[K]V, initMapSize),
	}
}

func NewHashMapFrom[K comparable, V any](m Map[K, V]) Map[K, V] {
	newM := make(map[K]V, m.Size())
	m.Range(func(key K, value V) bool {
		newM[key] = value
		return true
	})
	return &hashMap[K, V]{
		initSize: initMapSize,
		m:        newM,
	}
}

func (m *hashMap[K, V]) Load(key K) (V, bool) {
	v, ok := m.m[key]
	return v, ok
}

func (m *hashMap[K, V]) LoadOrDefault(key K, defaultValue V) V {
	v, ok := m.m[key]
	if ok {
		return v
	}
	return defaultValue
}

func (m *hashMap[K, V]) LoadIfAbsent(key K, f func(K) V) V {
	v, ok := m.m[key]
	if ok {
		return v
	}

	newV := f(key)
	m.m[key] = newV
	return newV
}

func (m *hashMap[K, V]) LoadAndStore(key K, value V) (V, bool) {
	v, ok := m.m[key]
	m.m[key] = value
	return v, !ok
}

func (m *hashMap[K, V]) Store(key K, value V) {
	m.m[key] = value
}

func (h *hashMap[K, V]) StoreAll(m Map[K, V]) {
	m.Range(func(key K, value V) bool {
		h.m[key] = value
		return true
	})
}

func (m *hashMap[K, V]) Size() int {
	return len(m.m)
}

func (m *hashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *hashMap[K, V]) Delete(key K) {
	delete(m.m, key)
}

func (m *hashMap[K, V]) LoadAndDelete(key K) (V, bool) {
	v, ok := m.m[key]
	if ok {
		delete(m.m, key)
	}
	return v, ok
}

func (m *hashMap[K, V]) Clear() {
	clear(m.m)
	m.m = make(map[K]V, m.initSize)
}

func (m *hashMap[K, V]) Range(f func(K, V) bool) {
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
