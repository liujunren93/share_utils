package helper

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GetUuidV3(name string) string {
	v1, _ := uuid.NewV4()
	variant := uuid.NewV3(v1, name)
	all := strings.ReplaceAll(variant.String(), "-", "")
	return all
}
