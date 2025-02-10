package id

import (
	"github.com/blue0121/easygo/misc/logger"
	"github.com/gofrs/uuid/v5"
)

func GenUUID() string {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		logger.Panic("Failed to generate uuid: %v", err)
	}
	return uuidV7.String()
}
