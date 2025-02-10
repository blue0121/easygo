package number

import (
	"github.com/blue0121/easygo/misc/util"
)

func MaskForInt32(len int) int32 {
	util.AssertIsTrue(len >= 0 && len <= 31, "length must be between 0 and 31")
	return ^(-1 << len)
}

func MaskForInt64(len int) int64 {
	util.AssertIsTrue(len >= 0 && len <= 63, "length must be between 0 and 63")
	return ^(-1 << len)
}
