package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string
	PasswordHash string
	IsAdmin      bool
}

func NewUser(ID string, password string, isAdmin bool) (*User, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &User{ID, string(h), isAdmin}
	return user, nil
}

type UserRepository interface {
	Save(ID string, passwordHash string, isAdmin bool) error
	Delete(ID string) error
	FindByID(ID string) (*User, error)
}
