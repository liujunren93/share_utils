package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shareChina/utils/log"
	"time"
)

type Claims struct {
	data interface{}
	jwt.StandardClaims
}

func CreateToken(data interface{}, jwtSecret string, ExpiresTime int64) string {
	claims := Claims{
		data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpiresTime,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, er := token.SignedString([]byte(jwtSecret))
	if er != nil {
		log.Logger.Info(er)
	}
	return signedString
}

func CheckToken(tokenString, JwtSecret string) (*Claims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(JwtSecret), nil
	})

	if err != nil {
		return nil, false
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}

//GetTokenInfo .
func GetTokenInfo(tokenString string, JwtSecret string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})
	claims, ok := token.Claims.(*Claims)
	return claims, ok
}
