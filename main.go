package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stewie1520/online-store/internal/config"
	"github.com/stewie1520/online-store/internal/route"
)

func init() {
	config.Init()
}

func main() {
	router := gin.Default()

	route.PrivateRoutes(router)

	router.Run(fmt.Sprintf(":%v", config.AppConfig.Port))
}
