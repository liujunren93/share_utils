package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/errors"
)

func ShouldBindJSON(ctx *gin.Context, dest interface{}) errors.Error {
	err := ctx.ShouldBindJSON(dest)
	if err != nil {
		return errors.NewBadRequest(err.Error())
	}
	return nil
}

func ShouldBindQuery(ctx *gin.Context, dest interface{}) errors.Error {
	err := ctx.ShouldBindQuery(dest)
	if err != nil {
		return errors.NewBadRequest(err.Error())
	}
	return nil
}
