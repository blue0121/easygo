package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteFormat(t *testing.T) {
	assert.Equal(t, "5B", ByteFormat(5))
	assert.Equal(t, "2K", ByteFormat(2048))
}

func TestParseByteFormat(t *testing.T) {
	b1, err := ParseByteFormat("2K")
	assert.Nil(t, err)
	assert.Equal(t, uint64(2048), b1)
}
