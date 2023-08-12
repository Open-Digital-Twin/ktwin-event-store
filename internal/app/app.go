package app

import (
	"github.com/ktwins/event-store/internal/app/infra/server"
)

func StartApp() {
	httpServer := server.NewHttpServer()
	httpServer.Configure()
	httpServer.Start()
}
