package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

// TODO work with database
func (q *Queries) Login(ctx context.Context, login, password string) (*domain.User, error) {
	sqlStatement := `SELECT * FROM user WHERE login = $1`
	row := q.pool.QueryRow(ctx, sqlStatement, login)
	user := &domain.User{}
	err := row.Scan(
		&user.Login,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("can't found the user: %w", err)
	}
	return user, nil
}
func (q *Queries) Register(ctx context.Context, email, login, password string) error {
	sqlStatement := `INSERT INTO user (login, email, password) VALUES ($1, $2, $3)`
	_, err := q.pool.Exec(ctx, sqlStatement, login, email, password)
	return err
}
