package convert

import (
	"encoding/json"
	"github.com/blue0121/easygo/misc/logger"
)

func ToJsonBytes(obj any) []byte {
	data, err := json.Marshal(obj)
	if err != nil {
		logger.Warn("Encode to json bytes/string error: %v", err)
		return nil
	}
	return data
}

func ToJsonString(obj any) string {
	data := ToJsonBytes(obj)
	if data == nil {
		return ""
	}
	return string(data)
}
