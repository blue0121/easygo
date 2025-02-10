package id

import (
	"github.com/blue0121/easygo/collection"
	"github.com/blue0121/easygo/misc/logger"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestSnowflakeId_Single(t *testing.T) {
	idSet := collection.NewHashSet[int64]()
	len := 1000000
	start := time.Now().UnixMilli()
	snowflakeGen(len, idSet)
	used := time.Now().UnixMilli() - start
	logger.Info("单线程, 个数: %d, 用时: %d ms, 速度: %g/ms.", len, used, float64(len)/float64(used))
	assert.Equal(t, len, idSet.Size())
}

func TestSnowflakeId_Multi(t *testing.T) {
	idSet := collection.NewSyncHashSet[int64]()
	var wg sync.WaitGroup
	c := 50
	len := 100000
	start := time.Now().UnixMilli()
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			snowflakeGen(len, idSet)
			wg.Done()
		}()
	}
	wg.Wait()
	used := time.Now().UnixMilli() - start
	logger.Info("多线程, 个数: %d, 用时: %d ms, 速度: %g/ms.", c*len, used, float64(c*len)/float64(used))
	assert.Equal(t, len*c, idSet.Size())
}

var (
	options = &EpochOptions{
		MachineId:     1,
		MachineIdBits: 10,
		SeqBits:       12,
	}
	snowflake = NewSnowflakeId(options)
)

func snowflakeGen(len int, idSet collection.Set[int64]) {
	for i := 0; i < len; i++ {
		idSet.Add(snowflake.gen())
	}
}
