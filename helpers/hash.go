package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(hash_password string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
