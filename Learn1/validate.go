package main

import "strings"

func (u *User) Validate() error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	if len(u.Name) < 2 {
		return ErrNameTooShort
	}

	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}

	return nil
}
