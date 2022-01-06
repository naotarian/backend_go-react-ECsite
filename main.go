package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowMethods = []string{}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	}
	config.AllowCredentials = false
	config.MaxAge = 24 * time.Hour
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "top root.",
		})
	})
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api root.",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test root.",
		})
	})
	r.Run(":8080")
}
