package api

import (
    "github.com/gin-gonic/gin/binding"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    db "github.com/JohannSuarez/GoBackend/db/sqlc"
)

// SQLStore provides all functions to execute SQL queries and transactions
type Server struct {
    store db.Store
    router *gin.Engine
}


// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {

    server := &Server{store: store}
    router := gin.Default()

    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("currency", validCurrency)
    }

    // Add routes to router
    router.POST("/users", server.createUser)
    router.POST("/accounts", server.createAccount)
    router.GET("/accounts/:id", server.getAccount)
    router.GET("/accounts", server.listAccount)

    router.POST("/transfers", server.createTransfer)

    server.router = router
    return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
    return server.router.Run(address)
}

func errorResponse(err error) gin.H {
    return gin.H{"error": err.Error()}
}
