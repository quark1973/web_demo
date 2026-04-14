package middlewares

import (
	"demo/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware()gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token:=ctx.GetHeader("Authorization")
		if token==""{
			ctx.JSON(401,gin.H{
				"message":"token is empty",
			})
			ctx.Abort()
			return
		}
		username,err:=utils.ParseJWT(token)
		if err!= nil {
			ctx.JSON(401,gin.H{
				"message":err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Set("username",username)
		ctx.Next()
	}
}