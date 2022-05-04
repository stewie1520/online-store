package main

import (
	"fmt"
	"github.com/stewie1520/online-store/internal/route/private_route"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stewie1520/online-store/internal/config"
)

func init() {
	config.Init()
}

func main() {
	router := gin.Default()

	private_route.Use(router)

	err := router.Run(fmt.Sprintf( ":%v", config.AppConfig.Port))
	if err != nil {
		return
	}
}
