package record

import (
	"context"
)

type repository interface {
	AddRecord(ctx context.Context, userID int, record any) error
	GetAllRecords(ctx context.Context, userID int, recordType any) (any, error)
}

type decoder interface {
	Encrypt(data []byte) (string, error)
	Decrypt(encrypted string) ([]byte, error)
}
