package app

import (
	"agwermann/dt-service/internal/app/infra/server"
)

func StartApp() {
	httpServer := server.NewHttpServer()
	httpServer.Configure()
	httpServer.Start()
}
