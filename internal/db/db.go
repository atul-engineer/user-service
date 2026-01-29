package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/users?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}