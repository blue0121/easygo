package id

import (
	"github.com/blue0121/easygo/misc/logger"
	"github.com/gofrs/uuid/v5"
)

func NewUuid() Id {
	u, err := uuid.NewV7()
	if err != nil {
		logger.Panic("Failed to generate uuid: %v", err)
	}
	return &uuidV7{
		id: u,
	}
}

func FromUuidString(text string) (Id, error) {
	u, err := uuid.FromString(text)
	if err != nil {
		return nil, err
	}
	return &uuidV7{
		id: u,
	}, nil
}

func FromUuidBytes(input []byte) (Id, error) {
	u, err := uuid.FromBytes(input)
	if err != nil {
		return nil, err
	}
	return &uuidV7{
		id: u,
	}, nil
}

type uuidV7 struct {
	id uuid.UUID
}

func (u *uuidV7) String() string {
	return u.id.String()
}

func (u *uuidV7) Bytes() []byte {
	return u.id.Bytes()
}
