package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type GreetMessage struct {
	Name string `json:"name" binding:"required"`
}

func init() {
	config.Init()
}

func main() {
	_, err := gorm.Open(postgres.Open(config.AppConfig.Database.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

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

	router.Run(fmt.Sprintf(":%v", config.AppConfig.Port))
}
