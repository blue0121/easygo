package id

import (
	"github.com/blue0121/easygo/misc/logger"
	"github.com/blue0121/easygo/misc/number"
	"sync"
	"time"
)

var (
	once  sync.Once
	epoch int64
)

func init() {
	value := "2025-01-01"
	t, err := time.ParseInLocation(time.DateOnly, value, time.Local)
	if err != nil {
		logger.Fatal("Failed to parse date, date: %s, err: %v", value, err)
	}
	epoch = t.UnixMilli()
}

type epochId struct {
	seqMask, seq, lastTs int64
}

func newEpochId(seqBits int) *epochId {
	seqMask := number.MaskForInt64(seqBits)
	return &epochId{
		seqMask: seqMask,
		seq:     0,
		lastTs:  0,
	}
}

func (id *epochId) genSeq() {
	ts := time.Now().UnixMilli()
	id.checkTs(ts)
	if id.lastTs == ts {
		id.seq = (id.seq + 1) & id.seqMask
		if id.seq == 0 {
			time.Sleep(time.Millisecond)
			id.lastTs = time.Now().UnixMilli()
		}
	} else {
		id.lastTs = ts
		id.seq = 0
	}
	// logger.Debug("seqMask: %d, seq: %d, lastTs: %d, ts: %d", id.seqMask, id.seq, id.lastTs, ts)
}

func (id *epochId) checkTs(ts int64) {
	if ts >= id.lastTs {
		return
	}
	interval := ts - id.lastTs
	if interval <= 50 {
		time.Sleep(time.Duration(interval) * time.Millisecond)
	} else {
		logger.Panic("System clock rollback, reject ID generator within %d ms", interval)
	}
}
