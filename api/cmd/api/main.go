package main

import (
	"github.com/gin-gonic/gin"
	"github.com/s9swata/billSplit/internal/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to billSplit API",
		})
	})
	router.Run(":8080")
}
