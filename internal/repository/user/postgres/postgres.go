package postgres

import (
	"context"
	"fmt"
	"github.com/Khasmag06/gophkeeper/internal/models"
	"github.com/Khasmag06/gophkeeper/internal/repository/repo_errs"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *repo {
	return &repo{
		pool: db,
	}
}

func (r *repo) CreateUser(ctx context.Context, user models.User) error {
	row := r.pool.QueryRow(ctx,
		` INSERT INTO users (login, password) 
                   VALUES($1, $2) RETURNING id`, user.Login, user.Password)

	var id int

	err := row.Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return repo_errs.ErrAlreadyExists
			}
			return err
		}
	}

	return nil
}

func (r *repo) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	row := r.pool.QueryRow(ctx,
		`	SELECT id, login, password, created_at
			      FROM users
			     WHERE login = $1`, login)

	var user models.User

	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan user by login: %w", err)
	}

	return &user, nil
}
