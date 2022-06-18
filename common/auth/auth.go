package auth

type Auther interface {
	// Token generated using refresh token or credentials
	Token(refreshToken string) (*Token, error)
	Inspect(token string) (data interface{}, tokenType int8, err error)
	SetData(k string, v interface{})
}

// Token can be short or long lived
type Token struct {
	// The token to be used for accessing resources
	AccessToken string `json:"access_token"`
	// RefreshToken to be used to generate a new token
	RefreshToken string `json:"refresh_token"`
	// Time of token creation
	Created int64 `json:"created"`
	// Time of token expiry
	Expiry int64 `json:"expiry"`
}
