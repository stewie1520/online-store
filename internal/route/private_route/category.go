package private_route

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/online-store/internal/controller"
)

func useCategory(routeGroup *gin.RouterGroup, categoryController controller.CategoryController) {
	routeGroup.POST("", categoryController.Create)
	routeGroup.GET("", categoryController.FetchMany)
}
