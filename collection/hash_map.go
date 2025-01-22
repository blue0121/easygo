package collection

type HashMap[K comparable, V any] struct {
	initSize int
	m        map[K]V
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &HashMap[K, V]{
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
	return &HashMap[K, V]{
		initSize: initMapSize,
		m:        newM,
	}
}

func (m *HashMap[K, V]) Load(key K) (V, bool) {
	v, ok := m.m[key]
	return v, ok
}

func (m *HashMap[K, V]) LoadOrDefault(key K, defaultValue V) V {
	v, ok := m.m[key]
	if ok {
		return v
	}
	return defaultValue
}

func (m *HashMap[K, V]) LoadIfAbsent(key K, f func(K) V) V {
	v, ok := m.m[key]
	if ok {
		return v
	}

	newV := f(key)
	m.m[key] = newV
	return newV
}

func (m *HashMap[K, V]) Store(key K, value V) {
	m.m[key] = value
}

func (m *HashMap[K, V]) Size() int {
	return len(m.m)
}

func (m *HashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *HashMap[K, V]) Delete(key K) {
	delete(m.m, key)
}

func (m *HashMap[K, V]) LoadAndDelete(key K) (V, bool) {
	v, ok := m.m[key]
	if ok {
		delete(m.m, key)
	}
	return v, ok
}

func (m *HashMap[K, V]) Clear() {
	clear(m.m)
	m.m = make(map[K]V, m.initSize)
}

func (m *HashMap[K, V]) Range(f func(K, V) bool) {
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
