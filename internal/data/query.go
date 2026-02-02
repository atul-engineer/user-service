package data

import (
	"context"
	"database/sql"
	"fmt"
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

func UpdateUser(ctx context.Context, db *sql.DB, isActive bool, id int) error {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println("rec from panic", r)
		}
	}()
	_ = ctx
	query := `UPDATE users SET is_active=$1 WHERE id=$2`
	_, err := db.Exec(query, isActive, id)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(ctx context.Context, db *sql.DB, limit int, offset int) (*[]User, error) {
	var users []User
	query := `SELECT * FROM users LIMIT $1 OFFSET $2`
	row, _ := db.QueryContext(ctx, query, limit, offset)
	for row.Next() {
		var user User
		row.Scan(&user.Id, &user.Name, &user.Email, &user.IsActive)
		//user.Id
		users = append(users, user)
	}
	return &users, nil
}