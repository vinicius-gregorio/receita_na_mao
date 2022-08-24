package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vinicius-gregorio/receita_na_mao/db/sqlc"
)

type createCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.CreateCategory(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) getAllCategories(ctx *gin.Context) {

	categories, err := server.store.ListCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

type updateCategoryRequestUri struct {
	ID int64 `uri:"id" binding:"required"`
}
type updateCategoryRequestBody struct {
	Name string `json:"name" binding:"required,min=3"`
}

func (server *Server) updateCategory(ctx *gin.Context) {
	var requestUri updateCategoryRequestUri
	var requestBody updateCategoryRequestBody

	if err := ctx.ShouldBindUri(&requestUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCategoryParams{
		ID:   requestUri.ID,
		Name: requestBody.Name,
	}
	category, err := server.store.UpdateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}
