package id

import (
	"github.com/blue0121/easygo/collection"
	"github.com/blue0121/easygo/misc/number"
	"sync"
)

var (
	snowflakeIdMap = collection.NewSyncHashMap[int, *snowflakeId]()
)

type snowflakeId struct {
	mux     *sync.Mutex
	options *EpochOptions

	tsBits, ipShift, tsShift int
	ip                       int64
	epochId
}

func NewSnowflakeId(options *EpochOptions) *snowflakeId {
	total := 63
	tsBits := total - options.MachineIdBits - options.SeqBits
	ip := options.MachineId & int(number.MaskForInt32(options.MachineIdBits))
	epochId := newEpochId(options.SeqBits)
	return &snowflakeId{
		mux:     &sync.Mutex{},
		options: options,
		tsBits:  tsBits,
		ipShift: tsBits + options.SeqBits,
		tsShift: options.SeqBits,
		ip:      int64(ip),
		epochId: *epochId,
	}
}

func (id *snowflakeId) gen() int64 {
	id.mux.Lock()
	defer id.mux.Unlock()
	id.genSeq()
	return (id.ip << id.ipShift) | ((id.lastTs - epoch) << id.tsShift) | id.seq
}

func SingleLongId() int64 {
	return LongId(0)
}

func LongId(machineId int) int64 {
	id := snowflakeIdMap.LoadIfAbsent(machineId, func(i int) *snowflakeId {
		options := NewSingleEpoch(machineId)
		return NewSnowflakeId(options)
	})
	return id.gen()
}
