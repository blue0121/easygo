package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSet_NewHashSetFrom(t *testing.T) {
	s1 := NewHashSet[int]()
	s1.Add(1, 2, 3)
	s2 := NewHashSetFrom(s1)

	assert.Equal(t, 3, s2.Size())
	assert.True(t, s2.Contains(1))
	assert.True(t, s2.Contains(2))
	assert.True(t, s2.Contains(3))
}

func TestHashSet_NewSyncHashSetFrom(t *testing.T) {
	s1 := NewSyncHashSet[int]()
	s1.Add(1, 2, 3)
	s2 := NewSyncHashSetFrom(s1)

	assert.Equal(t, 3, s2.Size())
	assert.True(t, s2.Contains(1))
	assert.True(t, s2.Contains(2))
	assert.True(t, s2.Contains(3))
}

func TestHashSet_Add(t *testing.T) {
	s := NewHashSet[int]()
	assert.True(t, s.Add(1, 2, 3))
	assert.Equal(t, 3, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))

	assert.False(t, s.Add(1))
}

func TestHashSet_AddAll(t *testing.T) {
	s1 := NewHashSet[int]()
	s1.Add(1, 2, 3)
	s2 := NewHashSet[int]()
	assert.True(t, s2.IsEmpty())
	s2.AddAll(s1)

	assert.Equal(t, 3, s2.Size())
	assert.False(t, s2.IsEmpty())
	assert.True(t, s2.Contains(1))
	assert.True(t, s2.Contains(2))
	assert.True(t, s2.Contains(3))
}

func TestHashSet_Delete(t *testing.T) {
	s1 := NewHashSet[int]()
	s1.Add(1, 2, 3)
	assert.True(t, s1.Delete(1))
	assert.False(t, s1.Delete(1))
	assert.True(t, s1.Delete(2))
	assert.True(t, s1.Delete(3))
	assert.False(t, s1.Delete(4))
	assert.True(t, s1.IsEmpty())
	assert.Equal(t, 0, s1.Size())
}
