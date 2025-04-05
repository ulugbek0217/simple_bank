package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/ulugbek0217/simple_bank/db/sqlc"
)

// Server serves HTTP requests for banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new Gin server with routes
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// adding routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts/balance", server.updateAccountBalance)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start starts the gin server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse returns the error body in JSON format
func errorResponse(err error) gin.H {
	return gin.H{"error:": err.Error()}
}
