package number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMaskForInt64(t *testing.T) {
	lens := []int{0, 1, 10, 31, 50, 63}

	for _, len := range lens {
		mask := MaskForInt64(len)
		fmt.Println(strconv.FormatInt(mask, 2))
		for i := 0; i < 64; i++ {
			if i < len {
				assert.Equal(t, int64(1), (mask>>i)&1, fmt.Sprintf("第 %d 位不为1", i))
			} else {
				assert.Equal(t, int64(0), (mask>>i)&1, fmt.Sprintf("第 %d 位不为0", i))
			}
		}
	}
}

func TestMaskForInt32(t *testing.T) {
	lens := []int{0, 1, 10, 31}

	for _, len := range lens {
		mask := MaskForInt32(len)
		fmt.Println(strconv.FormatInt(int64(mask), 2))
		for i := 0; i < 32; i++ {
			if i < len {
				assert.Equal(t, int32(1), (mask>>i)&1, fmt.Sprintf("第 %d 位不为1", i))
			} else {
				assert.Equal(t, int32(0), (mask>>i)&1, fmt.Sprintf("第 %d 位不为0", i))
			}
		}
	}
}
