package server

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/config"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/db"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/middleware"

	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Configure()
	Start()
}

type httpServer struct {
	engine        *gin.Engine
	appController Controller
}

func NewHttpServer() HttpServer {
	config.Load()
	dbConnection := db.NewDBConnection()
	appController := NewAppController(dbConnection)
	return &httpServer{
		engine:        gin.Default(),
		appController: appController,
	}
}

func (s *httpServer) Configure() {
	middleware.UseCors(s.engine)
	ConfigureRoutes(s.engine, s.appController)
	ConfigureSwagger(s.engine)
}

func (s *httpServer) Start() {
	s.engine.Run(":8080")
}
