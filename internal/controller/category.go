package controller

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/online-store/internal/domain"
	"github.com/stewie1520/online-store/internal/dto"
	"github.com/stewie1520/online-store/internal/service"
)

type categoryController struct {
	service service.CategoryService
}

type CategoryController interface {
	Create(*gin.Context)
}

func NewCategoryController(cs service.CategoryService) CategoryController {
	return &categoryController{
		service: cs,
	}
}

func (c *categoryController) Create(ctx *gin.Context) {
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

	if category, err := c.service.CreateCategory(context, &createCategoryParams); err != nil {
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
}
