package utils

import (
	"fmt"
	t2 "github.com/shareChina/utils/t"
	"net/http"
	"testing"
)

func TestGetConfig(t *testing.T) {
	fmt.Println(t2.MysqlConf)
	http.ListenAndServe(":10000", nil)

}
