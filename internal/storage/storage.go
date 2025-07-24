package storage

import "errors"

var (
	ErrUserExists   = errors.New("stub already exists")
	ErrUserNotFound = errors.New("stub not found")
	ErrAppNotFound  = errors.New("project not found")
)
