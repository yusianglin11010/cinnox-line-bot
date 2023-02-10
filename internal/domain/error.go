package domain

import "errors"

var (
	ErrUnexpected       = errors.New("unexpected error")
	ErrInvalidParameter = errors.New("invalid parameter")

	ErrMongoCreateFail = errors.New("failed to create mongo data")
	ErrMongoGetFail    = errors.New("failed to get mongo data")

	ErrUserNotExisted = errors.New("user not exist")
	ErrNoDocuments = errors.New("no document found")
)
