package middleware

import (
	"github.com/gin-gonic/gin"
	client2 "github.com/liujunren93/share_utils/client"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
	"github.com/liujunren93/share_utils/pkg/storage/userStore"
)

var client *client2.Client

func SetClient(c *client2.Client) {
	client = c
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newUserStore := userStore.NewUserStore(client.UserStore.KeepLoginTime, client.UserStore.Secret, client.Redis)
		if _, ok := newUserStore.Load(ctx); !ok {
			netHelper.Response(ctx.Writer, errors.StatusUnauthorized, errors.New(401, "登录信息失效"),nil)
			return
		}
		ctx.Set("client", client)
		ctx.Next()
	}
}
