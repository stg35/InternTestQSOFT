package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())
	router.Use(ManageHeader())

	router.GET("/when/:year", server.checkDate)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}