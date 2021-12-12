package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GreetMessage struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	router.GET("/:name", func(context *gin.Context) {
		name := context.Params.ByName("name")
		context.JSON(http.StatusOK, gin.H{
			"hello": name,
		})
	})

	router.POST("/greeting", func(context *gin.Context) {
		var greetMessage GreetMessage
		name := context.BindJSON(&greetMessage)
		context.JSON(http.StatusOK, gin.H{"hi": name})
	})

	router.Run(":8080")
}
