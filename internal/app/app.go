package app

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/server"
)

func StartApp() {
	httpServer := server.NewHttpServer()
	httpServer.Configure()
	httpServer.Start()
}
