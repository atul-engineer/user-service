package api

import "errors"

var (
	ErrInputRequired = errors.New("all fields are required")
)