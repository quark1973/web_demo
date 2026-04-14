package api

import (
	"demo/domain"
	"demo/middlewares"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetUpRouter() *gin.Engine {
	e := gin.Default()
	e.Use(CORS())
	v1 := e.Group("/api/auth")
	{
		// 路由
		v1.POST("/login", domain.Login)
		v1.POST("/register", domain.Register)
	}
	// v2 := e.Group("/api")
	// {
	// 	// 路由

	// }
	V3 := e.Group("/api")
	V3.GET("/exchangeRates", domain.GetExchangeRate)
	V3.Use(middlewares.AuthMiddleware())
	{
		// 路由
		V3.POST("/exchangeRates", domain.CreateExchangeRate)
		V3.POST("/articles", domain.CreateArticle)
		V3.GET("/articles", domain.GetArticle)
		V3.GET("/articles/:id", domain.GetArticalId)

		V3.POST("/articles/:id/like", domain.LikeArticle)
		V3.GET("/articles/:id/like", domain.GetLikes)

	}
	return e
}

func Login(c *gin.Context) {
	c.AbortWithStatusJSON(200, gin.H{
		"message": "Login success",
	})
}

func Register(c *gin.Context) {
	c.AbortWithStatusJSON(200, gin.H{
		"message": "Register success",
	})
}
