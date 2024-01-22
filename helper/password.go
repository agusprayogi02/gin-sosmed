package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyPassword(hashPass string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
}
