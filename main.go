package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type GreetMessage struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	dsn := "host=db user=gorm password=gormdeptrai3 dbname=gorm port=5432 sslmode=disable"
	gorm.Open(postgres.Open(dsn), &gorm.Config{})
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
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
