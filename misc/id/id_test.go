package id

import (
	"github.com/blue0121/easygo/misc/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEpoch(t *testing.T) {
	logger.Info("epoch: %d", epoch)
	assert.Equal(t, int64(1735660800000), epoch)
}
