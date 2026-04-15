package domain

import (
	"demo/global"
	"demo/repository"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gin-gonic/gin"
)

var cachekey = "article"
func CreateArticle(ctx *gin.Context) {
	var article repository.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	global.RedisDb.Del(cachekey).Err()
	ctx.JSON(201, gin.H{
		"message": "success",
	})

}

func GetArticle(ctx *gin.Context) {
	cachedata,err:= global.RedisDb.Get(cachekey).Result()
	if err == redis.Nil{
	var articles []repository.Article
	if err := global.Db.Find(&articles).Error; err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, articles)
	articleJson,err := json.Marshal(articles)
	if err!= nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := global.RedisDb.Set(cachekey,articleJson,60).Err(); err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
		ctx.JSON(200, articles)
}else if err != nil{
	ctx.JSON(500, gin.H{
		"message": err.Error(),
	})
	return
}else {
		var articles []repository.Article
		if err := json.Unmarshal([]byte(cachedata), &articles); err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, articles)
}
}

func GetArticalId(ctx *gin.Context){
	id :=ctx.Param("id")
	var article repository.Article
	if err:=global.Db.Where("id = ?",id).First(&article).Error;err!=nil{
		ctx.JSON(404,gin.H{
			"message":err.Error(),
		})
		return
	}
	ctx.JSON(200,article)
}