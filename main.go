package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	router.Run(":8080")
}
