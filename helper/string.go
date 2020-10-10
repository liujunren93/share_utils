package helper

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

func GetUuidV3(name string) string {
	v1, _ := uuid.NewV4()
	variant := uuid.NewV3(v1, name)
	all := strings.ReplaceAll(variant.String(), "-", "")
	return all
}

func RandString(length int) string {
	baseStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var randStr []byte
	buf := []byte(baseStr)
	for i := 0; i < length; i++ {
		randStr = append(randStr, buf[r.Intn(len(baseStr))])
	}
	return string(randStr)
}

func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}


