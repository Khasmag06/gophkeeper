package api

import (
	"context"
	"github.com/Khasmag06/gophkeeper/internal/models"
)

type authService interface {
	SignUp(ctx context.Context, user models.User) error
	Login(ctx context.Context, loginData models.User) (*models.TokensResponse, error)

	GenerateToken(userHash string) (string, error)
	ParseToken(accessToken string) (*models.TokenClaims, error)
}

type recordService interface {
	AddRecord(ctx context.Context, userID int, record any) error
	GetAllRecords(ctx context.Context, userID int, recordType any) (string, error)
}

type logger interface {
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
}

type decoder interface {
	Encrypt(data []byte) (string, error)
	Decrypt(encrypted string) ([]byte, error)
}
