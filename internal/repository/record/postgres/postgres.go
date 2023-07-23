package repo

import (
	"context"
	"fmt"
	"github.com/Khasmag06/gophkeeper/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *repo {
	return &repo{
		pool: db,
	}
}

func (r *repo) AddRecord(ctx context.Context, userID int, record any) error {
	switch rec := (record).(type) {
	case *models.LoginCredentials:
		return r.AddLoginCredentials(ctx, userID, rec)
	case *models.TextData:
		return r.AddTextData(ctx, userID, rec)
	case *models.BinaryData:
		return r.AddBinaryData(ctx, userID, rec)
	case *models.BankCard:
		return r.AddBankCard(ctx, userID, rec)
	default:
		return fmt.Errorf("unsupported record type for add record")
	}
}

func (r *repo) GetAllRecords(ctx context.Context, userID int, recordType any) (any, error) {
	switch recordType.(type) {
	case models.LoginCredentials:
		loginCredentials, err := r.GetAllLoginCredentials(ctx, userID)
		return loginCredentials, err
	case models.TextData:
		textData, err := r.GetAllTextData(ctx, userID)
		return textData, err
	case models.BinaryData:
		binaryData, err := r.GetAllBinaryData(ctx, userID)
		return binaryData, err
	case models.BankCard:
		bankCards, err := r.GetAllBankCards(ctx, userID)
		return bankCards, err
	default:
		return nil, fmt.Errorf("unsupported record type for get all records")
	}
}

func (r *repo) AddLoginCredentials(ctx context.Context, userID int, loginCreds *models.LoginCredentials) error {
	query := "INSERT INTO login_credentials (user_id, login, password) VALUES ($1, $2, $3)"
	_, err := r.pool.Exec(ctx, query, userID, loginCreds.Login, loginCreds.Password)
	if err != nil {
		return fmt.Errorf("failed to add login credentials: %w", err)
	}
	return nil
}

func (r *repo) GetAllLoginCredentials(ctx context.Context, userID int) ([]*models.LoginCredentials, error) {
	query := "SELECT login, password FROM login_credentials WHERE user_id = $1"
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get login credentials: %w", err)
	}
	defer rows.Close()

	var loginCredentials []*models.LoginCredentials
	for rows.Next() {
		creds := &models.LoginCredentials{}
		err := rows.Scan(&creds.Login, &creds.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan login credentials row: %w", err)
		}
		loginCredentials = append(loginCredentials, creds)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating login credentials rows: %w", err)
	}

	return loginCredentials, nil
}

func (r *repo) AddTextData(ctx context.Context, userID int, textData *models.TextData) error {
	query := "INSERT INTO text_data (user_id, data) VALUES ($1, $2)"
	_, err := r.pool.Exec(ctx, query, userID, textData.Data)
	if err != nil {
		return fmt.Errorf("failed to add text data: %w", err)
	}
	return nil
}

func (r *repo) GetAllTextData(ctx context.Context, userID int) ([]*models.TextData, error) {
	query := "SELECT data FROM text_data WHERE user_id = $1"
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get text data: %w", err)
	}
	defer rows.Close()

	var textData []*models.TextData
	for rows.Next() {
		data := &models.TextData{}
		err := rows.Scan(&data.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to scan text data row: %w", err)
		}
		textData = append(textData, data)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating text data rows: %w", err)
	}

	return textData, nil
}

func (r *repo) AddBinaryData(ctx context.Context, userID int, binaryData *models.BinaryData) error {
	query := "INSERT INTO binary_data (user_id, data) VALUES ($1, $2)"
	_, err := r.pool.Exec(ctx, query, userID, binaryData.Data)
	if err != nil {
		return fmt.Errorf("failed to add binary data: %w", err)
	}
	return nil
}

func (r *repo) GetAllBinaryData(ctx context.Context, userID int) ([]*models.BinaryData, error) {
	query := "SELECT data FROM binary_data WHERE user_id = $1"
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get binary data: %w", err)
	}
	defer rows.Close()

	var binaryData []*models.BinaryData
	for rows.Next() {
		data := &models.BinaryData{}
		err := rows.Scan(&data.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to scan binary data row: %w", err)
		}
		binaryData = append(binaryData, data)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating binary data rows: %w", err)
	}

	return binaryData, nil
}

func (r *repo) AddBankCard(ctx context.Context, userID int, bankCard *models.BankCard) error {
	query := "INSERT INTO bank_cards (user_id, card_number, expiration_date, cvv) VALUES ($1, $2, $3, $4)"
	_, err := r.pool.Exec(ctx, query, userID, bankCard.CardNumber, bankCard.ExpirationDate, bankCard.CVV)
	if err != nil {
		return fmt.Errorf("failed to add bank card: %w", err)
	}
	return nil
}

func (r *repo) GetAllBankCards(ctx context.Context, userID int) ([]*models.BankCard, error) {
	query := "SELECT card_number, expiration_date, cvv FROM bank_cards WHERE user_id = $1"
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bank cards: %w", err)
	}
	defer rows.Close()

	var bankCards []*models.BankCard
	for rows.Next() {
		card := &models.BankCard{}
		err := rows.Scan(&card.CardNumber, &card.ExpirationDate, &card.CVV)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bank card row: %w", err)
		}
		bankCards = append(bankCards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating bank card rows: %w", err)
	}

	return bankCards, nil
}
