package collection

import (
	"sync"
)

type syncHashMap[K comparable, V any] struct {
	initSize int
	mux      sync.RWMutex
	m        map[K]V
}

func NewSyncHashMap[K comparable, V any]() Map[K, V] {
	return &syncHashMap[K, V]{
		initSize: initMapSize,
		m:        make(map[K]V, initMapSize),
	}
}

func NewSyncHashMapFrom[K comparable, V any](m Map[K, V]) Map[K, V] {
	newM := make(map[K]V, m.Size())
	m.Range(func(key K, value V) bool {
		newM[key] = value
		return true
	})
	return &syncHashMap[K, V]{
		initSize: initMapSize,
		m:        newM,
	}
}

func (m *syncHashMap[K, V]) Load(key K) (V, bool) {
	m.mux.RLock()
	v, ok := m.m[key]
	m.mux.RUnlock()
	return v, ok
}

func (m *syncHashMap[K, V]) LoadOrDefault(key K, defaultValue V) V {
	m.mux.RLock()
	v, ok := m.m[key]
	m.mux.RUnlock()
	if ok {
		return v
	}
	return defaultValue
}

func (m *syncHashMap[K, V]) LoadIfAbsent(key K, f func(K) V) V {
	m.mux.RLock()
	v, ok := m.m[key]
	m.mux.RUnlock()
	if ok {
		return v
	}

	m.mux.Lock()
	defer m.mux.Unlock()
	if v, ok = m.m[key]; ok {
		return v
	}
	newV := f(key)
	m.m[key] = newV
	return newV
}

func (m *syncHashMap[K, V]) LoadAndStore(key K, value V) (V, bool) {
	m.mux.Lock()
	defer m.mux.Unlock()
	v, ok := m.m[key]
	m.m[key] = value
	return v, !ok
}

func (m *syncHashMap[K, V]) Store(key K, value V) {
	m.mux.Lock()
	m.m[key] = value
	m.mux.Unlock()
}

func (h *syncHashMap[K, V]) StoreAll(m Map[K, V]) {
	h.mux.Lock()
	m.Range(func(key K, value V) bool {
		h.m[key] = value
		return true
	})
	h.mux.Unlock()
}

func (m *syncHashMap[K, V]) Size() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.m)
}

func (m *syncHashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *syncHashMap[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

func (m *syncHashMap[K, V]) LoadAndDelete(key K) (V, bool) {
	m.mux.RLock()
	v, ok := m.m[key]
	m.mux.RUnlock()
	if ok {
		m.mux.Lock()
		delete(m.m, key)
		m.mux.Unlock()
	}
	return v, ok
}

func (m *syncHashMap[K, V]) Clear() {
	m.mux.Lock()
	clear(m.m)
	m.m = make(map[K]V, m.initSize)
	m.mux.Unlock()
}

func (m *syncHashMap[K, V]) Range(f func(K, V) bool) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
