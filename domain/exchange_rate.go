package domain

import (
	"demo/global"
	"demo/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context){
	var exchangeRate repository.ExchangeRate
	if err:=ctx.ShouldBindJSON(&exchangeRate);err!=nil{
		ctx.JSON(400,gin.H{
			"message":err.Error(),
		})
		return
	}
	exchangeRate.DateTime = time.Now()
	if err:=global.Db.AutoMigrate(&exchangeRate);err!=nil{
		ctx.JSON(500,gin.H{
			"message":err.Error(),
		})
		return
	}
	if err:=global.Db.Create(&exchangeRate).Error;err!=nil{
		ctx.JSON(500,gin.H{
			"message":err.Error(),
		})
		return
	}
	ctx.JSON(201,gin.H{
		"message":"success",
	})
}

func GetExchangeRate(ctx *gin.Context){
	var exchangeRate []repository.ExchangeRate
	// err:=global.Db.Where("date_time = ?",time.Now()).Find(&exchangeRate).Error;
	if err:=global.Db.Find(&exchangeRate).Error;err!=nil{
		ctx.JSON(404,gin.H{
			"message":err.Error(),
		})
		return
	}
	ctx.JSON(200,gin.H{
		"data":exchangeRate,
	})
}
