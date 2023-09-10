package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct{}

func NewPassword() *Password {
	return &Password{}
}

func (p *Password) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p *Password) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
