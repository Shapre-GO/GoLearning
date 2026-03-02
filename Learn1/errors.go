package main

import "errors"

var (
	ErrNameTooShort = errors.New("Failed: Name is too short")
	ErrInvalidEmail = errors.New("Failed: Email is invalid")
	ErrEmailTaken   = errors.New("Failed: This email already exists")
)
