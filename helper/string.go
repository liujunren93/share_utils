package helper

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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

func RandRange(min, max int) int {
	if min > max {
		panic("min must be less than max")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

//String2Int 字符串转int 忽略错误
func String2Int(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return integer

}

func Any2String(data any) string {
	return fmt.Sprintf("%v", data)
}

func SnakeString(s string) string {
	if s == "ID" {
		s = strings.ReplaceAll(s, "ID", "Id")
	}
	if s == "PK" {
		s = strings.ReplaceAll(s, "PK", "Pk")
	}

	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写

	return strings.ToLower(string(data[:]))
}

// keep substr left
func SubstrLeft(s, substr string) string {
	strLen := len(s)
	if strLen == 0 {
		return s
	}
	index := strings.Index(s, substr)
	if index < 0 {
		return s
	}
	return s[:strings.Index(s, substr)]
}

// keep substr right
func SubstrRight(s, substr string) string {
	strLen := len(s)
	if strLen == 0 {
		return s
	}
	index := strings.Index(s, substr) + len(substr)
	if strLen < index {
		return s
	}
	return s[index:]
}
