package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/online-store/internal/controller"
)

func PrivateRoutes(router *gin.Engine) {
	v1 := router.Group("/admin/v1")
	{
		controller, err := controller.InitializeCategoryController()
		if err != nil {
			log.Fatalf("error initializing category controller: %v", err)
		}

		v1.POST("/categories", controller.Create)
	}
}
