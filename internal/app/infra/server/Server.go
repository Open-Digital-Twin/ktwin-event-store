package server

import (
	"agwermann/event-store-service/internal/app/infra/middleware"

	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Configure()
	Start()
}

type httpServer struct {
	engine *gin.Engine
}

func NewHttpServer() HttpServer {
	return &httpServer{
		engine: gin.Default(),
	}
}

func (s *httpServer) Configure() {
	middleware.UseCors(s.engine)
	ConfigureRoutes(s.engine)
	ConfigureSwagger(s.engine)
}

func (s *httpServer) Start() {
	s.engine.Run(":8080")
}
