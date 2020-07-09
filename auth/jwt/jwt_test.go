package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	token := CreateToken("aaaa", "sadsadas")
	for  {
		time.Sleep(time.Second*5)
		if checkToken, b := CheckToken(token, "sadsadas");b{
			fmt.Println(checkToken)
		}else{
			t.Log("false")
			break
		}

	}


}
