package helper

import (
	"crypto/sha1"
	"fmt"
)

// sha1
func Sha1Str(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
