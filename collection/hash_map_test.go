package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap_Load(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Store(1, 11)
	assert.Equal(t, 1, m.Size())
	assert.False(t, m.IsEmpty())
	v, ok := m.Load(1)
	assert.True(t, ok)
	assert.Equal(t, 11, v)
}

func TestHashMap_LoadIfAbsent(t *testing.T) {
	f1 := func(k int) int {
		return 10 + k
	}
	f2 := func(k int) int {
		return 20 + k
	}
	m := NewHashMap[int, int]()
	assert.Equal(t, 11, m.LoadIfAbsent(1, f1))
	assert.Equal(t, 12, m.LoadIfAbsent(2, f1))
	assert.Equal(t, 11, m.LoadIfAbsent(1, f2))
	assert.Equal(t, 12, m.LoadIfAbsent(2, f2))
}

func TestHashMap_NewHashMapFrom(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Store(1, 11)
	m.Store(2, 22)
	newM := NewHashMapFrom(m)
	assert.Equal(t, 2, newM.Size())
	newM.Range(func(key int, value int) bool {
		assert.Equal(t, value, m.LoadOrDefault(key, 0))
		return true
	})
}
