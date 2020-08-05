package auth

import "time"

type Auth interface {
	// Token generated using refresh token or credentials
	Token(...TokenOption) (*Token, error)
	Inspect(token string) (interface{}, error)
}

// Token can be short or long lived
type Token struct {
	// The token to be used for accessing resources
	AccessToken string `json:"access_token"`
	// RefreshToken to be used to generate a new token
	RefreshToken string `json:"refresh_token"`
	// Time of token creation
	Created time.Time `json:"created"`
	// Time of token expiry
	Expiry time.Time `json:"expiry"`
}
