package private_route

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/online-store/internal/controller"
	"log"
)

func Use(router *gin.Engine) {
	v1 := router.Group("/admin/v1")
	{
		ctrl, err := controller.NewCategoryController()
		// TODO: Handle error
		if err != nil {
			log.Fatalf("error while create new category controller %v\n", err)
		}
		useCategory(v1.Group("/categories"), ctrl)
	}
}
