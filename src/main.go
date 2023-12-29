package main

import (
	"flowcraft/auth-api/v2/src/modules/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware()) // CORS middlewar
	auth.NewHandler(r)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("Hello World!")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	},
	)
	err := r.Run(":4050") // listen and serve on
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
