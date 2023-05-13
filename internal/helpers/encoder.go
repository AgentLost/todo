package helpers

import "golang.org/x/crypto/bcrypt"

func Encode(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}
