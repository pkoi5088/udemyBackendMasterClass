package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

//Server serves HTTP requests for out banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

//NewServer creates a new HTTP server and setup routing
func NerServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add route to router

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

//Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"errer": err.Error()}
}