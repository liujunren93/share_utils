package auth

import (
	"time"
)

type TokenOptions struct {
	// ID for the company_cli
	ID string
	// Secret for the company_cli
	Secret string
	// Expiry is the time the token should live for
	Expiry time.Duration
	//data
	Data interface{}
}

type TokenOption func(o *TokenOptions)



func WithData(data interface{}) TokenOption {
	return func(o *TokenOptions) {
		o.Data = data
	}
}

func WithSecret(secret string) TokenOption {
	return func(o *TokenOptions) {
		o.Secret = secret
	}
}

func WithExpiry(expiry time.Duration) TokenOption {
	return func(o *TokenOptions) {
		o.Expiry = expiry
	}
}

func NewOption(option ...TokenOption) TokenOptions {
	var to TokenOptions
	for _, a := range option {
		a(&to)
	}
	if to.Expiry == 0 {
		to.Expiry = time.Hour * 2
	}
	return to
}
