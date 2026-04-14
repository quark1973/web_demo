package domain

import (
	"demo/repository"
	"demo/utils"
	"demo/global"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context)  {
	var user repository.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedpwd,err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return		
	}
	user.Password = hashedpwd

	token,err :=utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return								
	}

	if err:=global.Db.AutoMigrate(&user);err!=nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return								
	}

	if err:=global.Db.Create(&user).Error;err!=nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})		
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func Login(ctx *gin.Context)  {
	var input struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	var user repository.User
	global.Db.Where("username = ?",input.Username).First(&user)

	if !utils.CheckPassword(input.Password,user.Password){
		ctx.JSON(400, gin.H{
			"error": "password password is incorrect",
		})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}
