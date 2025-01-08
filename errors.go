package gonius

import "errors"

var (
	ErrInvalidToken = errors.New("invalid access token")
	ErrNotFound     = errors.New("not found")
	ErrApiError     = errors.New("api error")
)
