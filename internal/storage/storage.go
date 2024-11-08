package storage

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrAppNotFound  = errors.New("app not found")
	ErrUserNotFound = errors.New("user not found")
)
