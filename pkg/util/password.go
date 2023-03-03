package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashResult, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashResult), nil
}

func ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
