package domain

import (
	"demo/global"
	"demo/repository"

	"github.com/gin-gonic/gin"
)

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
	ctx.JSON(201, gin.H{
		"message": "success",
	})
}

func GetArticle(ctx *gin.Context) {
	var articles []repository.Article
	if err := global.Db.Find(&articles).Error; err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, articles)
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