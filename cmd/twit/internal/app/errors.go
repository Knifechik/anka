package app

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrAccessDenied = errors.New("access denied")
	ErrInvalidAuth  = errors.New("invalid auth")
)
