package domain

import (
	"demo/global"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context){
	id := ctx.Param("id")
	likeKey := "article:"+id+":likes"
	if err := global.RedisDb.Incr(likeKey).Err(); err!= nil {
		ctx.JSON(400,gin.H{
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(200,gin.H{
		"message":"like success",
	})
}

func GetLikes(ctx *gin.Context){
	id := ctx.Param("id")
	likeKey := "article:"+id+":likes"
	LikeCount,err:=global.RedisDb.Get(likeKey).Result()
	if err == redis.Nil {
		LikeCount="0"
	}else if err!= nil {
		ctx.JSON(500,gin.H{
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(200,gin.H{
		"likes":LikeCount,
	})
}