package data

import (
	"context"
	"database/sql"
)

func CreateUser(ctx context.Context, db *sql.DB, usr *User) (*User, error) {
	var user User
	query := `INSERT INTO users (name, email, is_active) VALUES ($1, $2, $3) RETURNING id, name, email, is_active`
	err := db.QueryRowContext(ctx, query, usr.Name, usr.Email, usr.IsActive).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}