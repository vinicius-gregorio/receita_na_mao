package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/vinicius-gregorio/receita_na_mao/db/sqlc"
)

type createRecipeRequest struct {
	Name            string    `json:"name" binding:"required,min=3"`
	Categories      []string  `json:"categories" binding:"required"`
	Description     string    `json:"description" binding:"required,min=3"`
	PrepareMethod   string    `json:"prepare_method" binding:"required,min=3"`
	Ingredients     []string  `json:"ingredients" binding:"required,min=3"`
	PreparationTime time.Time `json:"preparation_time" binding:"required"`
}

func (server *Server) createRecipe(ctx *gin.Context) {
	var request createRecipeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRecipeParams{
		Name:            request.Name,
		Categories:      request.Categories,
		Description:     request.Description,
		PrepareMethod:   request.PrepareMethod,
		Ingredients:     request.Ingredients,
		Rating:          sql.NullInt32{5, true},
		PreparationTime: request.PreparationTime,
	}

	recipe, err := server.store.CreateRecipe(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

type listRecipesRequest struct {
	//Query PARAMS
	PAGEID   int32 `form:"page_id" binding:"required,min=1"`
	PAGESize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListRecipes(ctx *gin.Context) {

	var req listRecipesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListRecipesParams{
		Limit:  req.PAGESize,
		Offset: (req.PAGEID - 1) * req.PAGESize,
	}
	recipes, err := server.store.ListRecipes(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, recipes)
}

type listRecipesByCategoryRequest struct {
	//Query PARAMS
	PAGEID     int32    `form:"page_id" binding:"required,min=1"`
	PAGESize   int32    `form:"page_size" binding:"required,min=5,max=10"`
	Categories []string `form:"categories" binding:"required"`
}

func (server *Server) ListRecipesByCategory(ctx *gin.Context) {

	var req listRecipesByCategoryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListRecipesByCategoryParams{
		Limit:      req.PAGESize,
		Offset:     (req.PAGEID - 1) * req.PAGESize,
		Categories: req.Categories,
	}
	recipes, err := server.store.ListRecipesByCategory(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, recipes)
}
