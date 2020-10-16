package helper

import (
	"crypto/sha1"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// sha1
func Sha1Str(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum(nil))
}


func NewPassword(secret, password string, cost int) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(secret+password), cost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}

func CheckPassword(secret, hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(secret+password))
}
