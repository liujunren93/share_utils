package helper

import (
	"fmt"
	"testing"
	"time"
)

func TestDateTime(t *testing.T) {

	unix, err := String2Unix("2006-01-02", "2006-01-02", nil)
	fmt.Println(time.Now().Unix(),	unix,err)
}
