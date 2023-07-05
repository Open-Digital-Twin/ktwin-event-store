package app

import (
	"agwermann/event-store-service/internal/app/infra/server"
)

func StartApp() {
	httpServer := server.NewHttpServer()
	httpServer.Configure()
	httpServer.Start()
}
