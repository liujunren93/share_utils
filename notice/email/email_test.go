package email

import (
	"log"
	"testing"
)

func TestSend(t *testing.T) {
	//
	config := Config{
		Host:     "smtp.qq.com:25",
		UserName: "768712894@qq.com",
		Password: "harouaiqldpcbdfe",
	}

	email := Email{
		Subject:  "aaaaa",
		To:       []string{"768712894@qq.com"},
		MailType: "html",
		Body:     `<a href="http://www.baidu.com">asadsad</a>`,
	}
	err := config.SendMsg(email)
	log.Fatal(err)
	return

}
