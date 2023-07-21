package record

import (
	"context"
	"encoding/json"
)

type service struct {
	repo    repository
	decoder decoder
}

func New(repo repository, decoder decoder) *service {
	return &service{
		repo:    repo,
		decoder: decoder,
	}
}

func (s *service) AddRecord(ctx context.Context, userID int, record any) error {
	return s.repo.AddRecord(ctx, userID, record)
}

func (s *service) GetAllRecords(ctx context.Context, userID int, recordType any) (string, error) {
	records, err := s.repo.GetAllRecords(ctx, userID, recordType)
	if err != nil {
		return "", err
	}
	recordBytes, err := json.Marshal(records)
	if err != nil {
		return "", err
	}

	encryptedRecord, err := s.decoder.Encrypt(recordBytes)
	if err != nil {
		return "", err
	}

	return encryptedRecord, nil
}
