package route

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/internal/controller"
)

func PrivateRoutes(router *gin.Engine) {
	v1 := router.Group("/admin/v1")
	{
		v1.POST("/categories", controller.CreateCategoryController)
	}
}
