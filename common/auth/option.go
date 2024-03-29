package auth

type TokenOptions struct {
	// Secret for the company_cli
	Secret string
	// Expiry is the time the token should live for
	Expiry int64
	//data
	Data map[string]interface{}
	// //RefreshTokenKeepData RefreshToken是否保存data，如果保存 生成token 使用RefreshToken的data
	// RefreshTokenKeepData bool
}

type TokenOption func(o *TokenOptions)

func WithData(data map[string]interface{}) TokenOption {
	return func(o *TokenOptions) {
		o.Data = data
	}
}

// func WithRefreshTokenKeepData(keep bool) TokenOption {
// 	return func(o *TokenOptions) {
// 		o.RefreshTokenKeepData = keep
// 	}
// }

func WithSecret(secret string) TokenOption {
	return func(o *TokenOptions) {
		o.Secret = secret
	}
}

func WithExpiry(expiry int64) TokenOption {
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
		to.Expiry = 7200
	}
	return to
}
