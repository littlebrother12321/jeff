package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {
	// Generate a hash
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

/*
* Returns true if validated.
 */
func CheckPassword(hashedPassword []byte, password []byte) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, password) == nil
}
