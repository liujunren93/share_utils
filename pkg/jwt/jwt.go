package jwt

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

func (c CustomClaims) CreateJwtToken(data interface{}, claims *jwt.StandardClaims, secret string) (string, error) {
	c.Data = data
	c.StandardClaims = *claims
	withClaims := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return withClaims.SignedString([]byte(secret))
}

func (c CustomClaims) CheckJwtToken(token, secret string) (*CustomClaims, bool) {
	parse, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false
	}
	if claims, ok := parse.Claims.(*CustomClaims); ok && parse.Valid {
		return claims, true
	}
	return nil, false
}
