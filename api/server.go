package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vinicius-gregorio/receita_na_mao/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	//category
	router.POST("/category", server.createCategory)
	router.GET("/category", server.getAllCategories)
	router.PUT("/category/:id", server.updateCategory)

	//recipe
	router.POST("/recipe", server.createRecipe)
	router.GET("/recipe", server.ListRecipes)
	router.PUT("/recipe/:id", server.updateCategory)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
