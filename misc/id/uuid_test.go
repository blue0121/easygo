package id

import (
	"github.com/blue0121/easygo/misc/logger"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestUuidV7(t *testing.T) {
	c := 10
	uidMap := make(map[string]struct{}, c)
	uidList := make([]string, 0, c)
	for i := 0; i < c; i++ {
		id := NewUuid()
		uidMap[id.String()] = struct{}{}
		uidList = append(uidList, id.String())
		logger.Info(id.String())

		id2, err := FromUuidString(id.String())
		assert.NoError(t, err)
		assert.Equal(t, id, id2)

		id3, err := FromUuidBytes(id.Bytes())
		assert.NoError(t, err)
		assert.Equal(t, id, id3)
	}
	assert.Equal(t, c, len(uidMap))
	assert.True(t, sort.StringsAreSorted(uidList))
}
