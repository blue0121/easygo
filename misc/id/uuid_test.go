package id

import (
	"github.com/blue0121/easygo/misc/logger"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestUuidV7(t *testing.T) {
	c := 10
	uuidMap := make(map[string]struct{}, c)
	uuidList := make([]string, 0, c)
	for i := 0; i < c; i++ {
		uuid := GenUUID()
		uuidMap[uuid] = struct{}{}
		uuidList = append(uuidList, uuid)
		logger.Info(uuid)
	}
	assert.Equal(t, c, len(uuidMap))
	assert.True(t, sort.StringsAreSorted(uuidList))
}
