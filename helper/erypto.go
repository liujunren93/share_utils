package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// sha1
func Sha1Str(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

//Md5Str
func Md5Str(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum(nil))
}


//GetUuidV3 获得
func GetUuidV3(name string,keepLine bool) string {
	u := uuid.New()
	uuidv3 := uuid.NewMD5(u, []byte(name))
	if keepLine {
		return uuidv3.String()
	}
	all := strings.ReplaceAll(uuidv3.String(), "-", "")
	return all
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
