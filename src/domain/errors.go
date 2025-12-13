package domain

import "errors"

var (
	ErrUserIDRequired  = errors.New("user_id is required")
	ErrContentRequired = errors.New("content is required")
)
