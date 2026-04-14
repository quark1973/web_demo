package main

import (
	"context"
	"fmt"
	"log"
	"demo/config"
	"demo/api"
	"github.com/gin-gonic/gin"
)



func main() {
	ctx:=context.Background()
	fmt.Println("start")
	//初始化配置文件
	err:=config.InitConfig(ctx)
	if err!= nil {
		log.Fatalln(err)
	}

	//打印配置文件
	fmt.Println(config.AppConfig)
	fmt.Println(config.AppConfig.App.Port)
	
	//在创建路由前设置为 Release 模式
    gin.SetMode(gin.ReleaseMode)
	e:=api.SetUpRouter()
	PORT:=config.AppConfig.App.Port
	if PORT==""{
		PORT="8080"
	}
	err=e.Run(":" + PORT)
	log.Fatal(err)

}