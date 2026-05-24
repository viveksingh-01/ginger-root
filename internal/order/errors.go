package order

import "errors"

var (
	ErrOrderNotFound    = errors.New("order not found")
	ErrOrderForbidden   = errors.New("order does not belong to user")
)
