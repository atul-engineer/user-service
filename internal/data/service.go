package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"user-service/internal/api"
)

type UserService struct {
	db *sql.DB
}


func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}


func (usrv *UserService) CreateUser(ctx context.Context, usr *User) (*User, error) {
	if usr.Email == "" || usr.Name == "" {
		return nil, api.ErrInputRequired
	}
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel()

	user, err := CreateUser(ctx, usrv.db, usr)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, fmt.Errorf("db timeout: %w", err)
		default:
			return nil, fmt.Errorf("create user: %w", err)
		}
	}
	return user, nil
}

func (usrv *UserService) UpdateUserStatus(isActive bool, id int) error {
	// err := UpdateUserStatus(usrv.db, isActive, id)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }
	return nil
}