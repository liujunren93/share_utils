package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/common/auth"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
)

func Session(au auth.Auther, setsession func(ctx *gin.Context, authData interface{}) errors.Error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		authData, tp, err := au.Inspect(token)
		if err != nil || tp != 1 {
			ctx.Next()
			return
		}
		errs := setsession(ctx, authData)
		if errs != nil {
			netHelper.Response(ctx, nil, errs, nil)
			return
		}

	}
}
