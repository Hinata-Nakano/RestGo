package domain

import "errors"

var (
	ErrUserIDRequired  = errors.New("user_id is required")
	ErrContentRequired = errors.New("content is required")
	ErrNameRequired    = errors.New("name is required")
	ErrEmailRequired   = errors.New("email is required")
)
