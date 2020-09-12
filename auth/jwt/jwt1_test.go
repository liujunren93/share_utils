package jwt

import (
	"fmt"
	auth2 "github.com/liujunren93/share_utils/auth"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	data := struct {
		ID string
	}{"sads"}
	auth := NewAuth()
	token, _ := auth.Token(auth2.WithExpiry(time.Hour), auth2.WithData(data), auth2.WithSecret("adsadsadsa"))
	//inspect, err := auth.Inspect(token.RefreshToken)

	//fmt.Println(inspect, err )
	refresh(token.AccessToken)

}

func refresh(token string) {
	auth := NewAuth()
	t, err := auth.Token(auth2.WithToken(token),auth2.WithSecret("adsadsadsa"))
	fmt.Println(t,err)
}
