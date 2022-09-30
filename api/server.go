package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/rest-api-go/db/gorm/models"
	db "github.com/go-study/rest-api-go/db/sqlc"
)

type Server struct {
	store2 *models.Gorm
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store, store2 *models.Gorm) *Server {
	server := &Server{store: store, store2: store2}

	g := models.Gorm{}
	g.DbInit()

	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/tokens", server.FindTokens)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
