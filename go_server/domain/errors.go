package domain

import "errors"

var (
	ErrUserWithEmailAlreadyExist    = errors.New("User with Email already Exist")
	ErrUserWithUsernameAlreadyExist = errors.New("User with Username already Exist")
	ErrNoResult                     = errors.New("no result")
)
