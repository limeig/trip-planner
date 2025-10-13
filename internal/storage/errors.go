package storage

import "errors"

var (
	// User related errors
	ErrUserNotFound = errors.New("user does not exist")
	ErrUserExists   = errors.New("user exists")

	// Location related errors
	ErrLocationNotFound = errors.New("location does not exist")
	ErrLocationExists   = errors.New("location exists")

	// Trip related errors
	ErrTripInvalid = errors.New("trip contains non-exixtent location")
)
