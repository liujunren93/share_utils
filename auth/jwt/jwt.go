package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/liujunren93/share_utils/auth"
	"time"
)

type jwtAuth struct {
	options auth.TokenOptions
}

func NewAuth() auth.Auth {
	return new(jwtAuth)
}

type jwtClaims struct {
	Data interface{}
	Type int8 //1:token 2:refresh token
	jwt.StandardClaims
}

//Inspect 验证token
func (j *jwtAuth) Inspect(tokenStr string) (interface{}, error) {

	tk, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(j.options.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tk.Claims.(*jwtClaims); ok && tk.Valid {
		return claims, nil
	}

	return nil, errors.New("token error")
}

//Token get token, if option.token!="" will refresh token
func (j *jwtAuth) Token(option ...auth.TokenOption) (*auth.Token, error) {
	op := auth.NewOption(option...)
	var token auth.Token
	j.options = op
	if op.Secret == "" {
		return nil, errors.New("secret is empty")
	}
	if op.RefreshToken != "" { //刷新token
		inspect, err := j.Inspect(op.RefreshToken)
		if err != nil {
			return nil, err
		}
		if jc, ok := inspect.(*jwtClaims); ok {
			if jc.Type == 1 {
				return nil, errors.New("cannot refresh token with token")
			}
			j.options.Data = jc.Data
		}
	}
	token.Created = time.Now().Local().Unix()
	accessToken, err := j.createToken(1)
	if err != nil {
		return nil, err
	}
	token.AccessToken = accessToken

	refreshToken, err := j.createToken(2)
	if err != nil {
		return nil, err
	}
	token.RefreshToken = refreshToken

	token.Expiry = token.Created + int64(j.options.Expiry)
	return &token, nil
}

//create token
//tkType token type （token=1，refresh=2）
func (j *jwtAuth) createToken(tkType int8) (string, error) {
	expiry := j.options.Expiry

	if tkType == 2 {
		expiry += time.Hour * 2

	}

	claims := jwtClaims{
		Data: j.options.Data,
		Type: tkType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiry).Local().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(j.options.Secret))

	return signedString, err
}
