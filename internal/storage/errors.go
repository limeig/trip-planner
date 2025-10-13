package storage

import "errors"

var (
	// User related errors
	ErrUserNotFound = errors.New("user does not exist")
	ErrUserExists   = errors.New("user exists")

	// Trip related errors
	ErrTripInvalid = errors.New("trip is invalid")
)
