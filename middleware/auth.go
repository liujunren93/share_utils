package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
)

const ISLOGIN = "is_login"

func Auth(ctx *gin.Context) {
	if ctx.GetBool(ISLOGIN) {
		ctx.Next()
	} else {
		// reflushToken := ctx.Request.Header.Get("ReflushToken")

		netHelper.Response(ctx, errors.NewPublic(errors.StatusTokenTimeout, "登录信息已过期"), nil, nil)
		ctx.Abort()
		return
	}

}
