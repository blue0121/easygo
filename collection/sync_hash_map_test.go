package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyncHashMap_Load(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Store(1, 11)
	assert.Equal(t, 1, m.Size())
	assert.False(t, m.IsEmpty())
	v, ok := m.Load(1)
	assert.True(t, ok)
	assert.Equal(t, 11, v)
}

func TestSyncHashMap_StoreAll(t *testing.T) {
	m1 := NewSyncHashMap[int, int]()
	m1.Store(1, 11)

	m2 := NewSyncHashMap[int, int]()
	m2.Store(2, 22)

	m3 := NewSyncHashMap[int, int]()
	m3.StoreAll(m1)
	m3.StoreAll(m2)
	assert.Equal(t, 2, m3.Size())
	assert.Equal(t, 11, m3.LoadOrDefault(1, 100))
	assert.Equal(t, 22, m3.LoadOrDefault(2, 100))
}

func TestSyncHashMap_LoadIfAbsent(t *testing.T) {
	f1 := func(k int) int {
		return 10 + k
	}
	f2 := func(k int) int {
		return 20 + k
	}
	m := NewSyncHashMap[int, int]()
	assert.Equal(t, 11, m.LoadIfAbsent(1, f1))
	assert.Equal(t, 12, m.LoadIfAbsent(2, f1))
	assert.Equal(t, 11, m.LoadIfAbsent(1, f2))
	assert.Equal(t, 12, m.LoadIfAbsent(2, f2))
}

func TestSyncHashMap_NewHashMapFrom(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Store(1, 11)
	m.Store(2, 22)
	newM := NewSyncHashMapFrom(m)
	assert.Equal(t, 2, newM.Size())
	newM.Range(func(key int, value int) bool {
		assert.Equal(t, value, m.LoadOrDefault(key, 0))
		return true
	})
}

func TestSyncHashMap_LoadAndStore(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	r1, ok := m.LoadAndStore(1, 11)
	assert.Equal(t, 0, r1)
	assert.True(t, ok)

	r2, ok := m.LoadAndStore(1, 12)
	assert.Equal(t, 11, r2)
	assert.False(t, ok)
}
