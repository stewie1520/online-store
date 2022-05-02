package controller

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/internal/domain"
	"github.com/stewie1520/internal/dto"
	"github.com/stewie1520/internal/service"
)

func CreateCategoryController(ctx *gin.Context) {
	var categoryDto dto.CreateCategoryDto
	if err := ctx.ShouldBindJSON(&categoryDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	createCategoryParams := domain.CreateCategoryParams{
		Name:     categoryDto.Name,
		Slug:     categoryDto.Slug,
		Author:   sql.NullInt32{Int32: categoryDto.Author, Valid: true},
		Isactive: sql.NullBool{Bool: categoryDto.IsActive, Valid: true},
	}

	context := context.Background()

	if category, err := service.CreateCategory(context, &createCategoryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"payload": category,
			"status":  "ok",
		})
	}
	return
}
