package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/liujunren93/share_utils/auth"
)

type jwtAuth struct {
	options auth.TokenOptions
}

func NewAuth(option ...auth.TokenOption) auth.Auther {
	jAuth := new(jwtAuth)
	op := auth.NewOption(option...)
	jAuth.options = op

	return jAuth
}

type JwtClaims struct {
	Data map[string]interface{}
	Type int8 //1:token 2:refresh token
	jwt.StandardClaims
}

func (j *jwtAuth) SetData(k string, v interface{}) {
	if j.options.Data == nil {
		j.options.Data = make(map[string]interface{})
	}

	j.options.Data[k] = v
}
func (j *jwtAuth) GetOptions() auth.TokenOptions {
	return j.options
}

//Inspect 验证token
func (j *jwtAuth) Inspect(tokenStr string) (interface{}, error) {

	tk, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.options.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tk.Claims.(*JwtClaims); ok && tk.Valid {
		if claims.Type == 1 {
			return claims.Data, nil
		}
	}

	return nil, errors.New("token error")
}

//Token get token, if RefreshToken!="" will refresh token
func (j *jwtAuth) Token(RefreshToken string) (*auth.Token, error) {
	var token auth.Token
	if j.options.Secret == "" {
		return nil, errors.New("secret is empty")
	}
	if RefreshToken != "" { //刷新token
		inspect, err := j.Inspect(RefreshToken)
		if err != nil {
			return nil, err
		}
		if jc, ok := inspect.(*JwtClaims); ok {
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
		expiry += 7200

	}

	claims := JwtClaims{
		Data: j.options.Data,
		Type: tkType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiry) * time.Second).Local().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(j.options.Secret))
	j.options.Data = nil
	return signedString, err
}
