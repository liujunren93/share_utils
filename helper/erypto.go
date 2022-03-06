package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/url"
	"sort"
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
func GetUuidV3(name string, keepLine bool) string {
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

func MapEncodeEscape(data map[string]interface{}) string {
	if data == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := data[k]
		keyEscaped := url.QueryEscape(k)
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(fmt.Sprintf("%v", vs)))
	}
	return buf.String()
}
