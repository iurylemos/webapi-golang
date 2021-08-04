package server

import "github.com/gin-gonic/gin"

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

}