package carver

import (
	"errors"
)

var (
	// ErrInvalid is thrown for invalid argument
	ErrInvalid = errors.New("invalid argument")
	// ErrFormat is thrown for invalid input format
	ErrFormat = errors.New("invalid input format")
)
