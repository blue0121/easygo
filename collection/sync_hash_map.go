package collection

import (
	"sync"
)

type SyncHashMap[K comparable, V any] struct {
	initSize int
	mux      sync.RWMutex
	m        map[K]V
}

func NewSyncHashMap[K comparable, V any]() Map[K, V] {
	return &SyncHashMap[K, V]{
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
	return &SyncHashMap[K, V]{
		initSize: initMapSize,
		m:        newM,
	}
}

func (m *SyncHashMap[K, V]) Load(key K) (V, bool) {
	m.mux.RLock()
	v, ok := m.m[key]
	m.mux.RUnlock()
	return v, ok
}

func (m *SyncHashMap[K, V]) LoadOrDefault(key K, defaultValue V) V {
	m.mux.RLock()
	v, ok := m.m[key]
	m.mux.RUnlock()
	if ok {
		return v
	}
	return defaultValue
}

func (m *SyncHashMap[K, V]) LoadIfAbsent(key K, f func(K) V) V {
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

func (m *SyncHashMap[K, V]) Store(key K, value V) {
	m.mux.Lock()
	m.m[key] = value
	m.mux.Unlock()
}

func (m *SyncHashMap[K, V]) Size() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.m)
}

func (m *SyncHashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *SyncHashMap[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

func (m *SyncHashMap[K, V]) LoadAndDelete(key K) (V, bool) {
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

func (m *SyncHashMap[K, V]) Clear() {
	m.mux.Lock()
	clear(m.m)
	m.m = make(map[K]V, m.initSize)
	m.mux.Unlock()
}

func (m *SyncHashMap[K, V]) Range(f func(K, V) bool) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
