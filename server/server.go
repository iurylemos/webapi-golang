package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iurylemos/webapi-golang/server/routes"
)

// Struct to gin create server
type Server struct {
	port string
	server *gin.Engine
}

// Method to create instance of server
func NewServer() Server {
	return Server{
		port: "5000",
		server: gin.Default(),
	}
}

// Method to run server
func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}