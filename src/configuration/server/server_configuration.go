package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-runner/src/configuration/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (server *Server) Run() {
	router := routes.RouterConfiguration(server.server)

	log.Print("Server is running at port: ", server.port)
	log.Fatal(router.Run(":" + server.port))
}