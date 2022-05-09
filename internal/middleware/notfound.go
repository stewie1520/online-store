package middleware

import "github.com/gin-gonic/gin"

func notFound() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
}
