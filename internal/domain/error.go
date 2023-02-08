package domain

import "errors"

var (
	ErrUnexpected = errors.New("unexpected error")

	ErrMongoCreateFail = errors.New("failed to create mongo data")
	ErrMongoGetFail    = errors.New("failed to get mongo data")

	ErrUserNotExisted = errors.New("user not exist")
)
