package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// sha1
func Sha1Str(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func Sha1Interface(src interface{}) (string, error) {
	fmt.Println(src)
	hash := sha1.New()
	data, err := json.Marshal(src)
	if err != nil {
		return "", err
	}
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Md5Str
func Md5Str(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

// GetUuidV3 获得
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

func AesEncrypt(orig string, key string) (str string, err error) {

	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blsize := block.BlockSize()
	origByte := PKCS7Padding([]byte(orig), blsize)
	bm := cipher.NewCBCEncrypter(block, keyByte[:blsize])
	tmp := make([]byte, len(origByte))
	bm.CryptBlocks(tmp, []byte(origByte))

	return base64.StdEncoding.EncodeToString(tmp), nil
}

func AesDecrypt(cryted string, key string) (str string, err error) {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	bm := cipher.NewCBCDecrypter(block, keyByte[:block.BlockSize()])
	tmp := make([]byte, len(crytedByte))
	bm.CryptBlocks(tmp, crytedByte)
	return string(PKCS7UnPadding(tmp)), nil
}

// 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
