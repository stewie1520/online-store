package middleware

import "github.com/gin-gonic/gin"

func UseAppMiddlewares(router *gin.Engine) {
	router.Use(appCORS())
	router.Use(appLogger())
	router.NoRoute(notFound())
}
