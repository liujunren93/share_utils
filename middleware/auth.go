package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

//在客户端验证
func ClientAuth(client client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.GetHeader("Authorization")
		//client.
	}
}
