package server

import (
	"fmt"
	"go-hex-http-api/internal/infra/server/handler/health"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("#{host}:#{port}"),
		engine: gin.New(),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
}
