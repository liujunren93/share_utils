package helper

import (
	"fmt"
	"slices"
	"strings"
)

type sortMap map[string]any

func NewSortMapFromMap(src map[string]any) sortMap {
	newMap := sortMap{}
	for key, val := range src {
		newMap[key] = val
	}
	return newMap
}
func NewSortMap() *sortMap {
	return &sortMap{}
}

func (s sortMap) Add(key string, val any) {
	s[key] = val
}
func (s sortMap) Del(key string) {
	delete(s, key)
}

// eof=true 最后一个元素 需要处理好拼接的结尾
func (s sortMap) Encode(encodeFunc func(key string, val any, eof bool) string) string {
	mapLen := len(s)
	keys := make([]string, 0, mapLen)
	for k := range s {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	if encodeFunc == nil {
		encodeFunc = func(key string, val any, eof bool) string {
			str := fmt.Sprintf("%v=%v", key, val)
			if !eof {
				str += "&"
			}
			return str
		}
	}
	encodeStr := strings.Builder{}
	for k, v := range keys {
		eof := k == mapLen-1
		encodeStr.WriteString(encodeFunc(v, s[v], eof))
	}
	return encodeStr.String()
}
