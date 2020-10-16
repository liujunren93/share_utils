package jwt

import (
	"github.com/liujunren93/share_utils/auth/jwt"
	"testing"
)

type data struct {
	ID   int
	Name string
}

func TestCustomClaims_CheckJwtToken(t *testing.T) {
	auth := jwt.NewAuth()
	auth.Token()
}
