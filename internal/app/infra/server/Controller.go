package server

import (
	"agwermann/event-store-service/internal/app/context/twinevent"
	"agwermann/event-store-service/internal/app/context/twininstance"
	"agwermann/event-store-service/internal/app/context/twininterface"

	"github.com/gin-gonic/gin"
)

// Twin Interfaces

func GetAllTwinInterfaces(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer()
	container.Controller.GetAllTwinInterfaces(g)
}

func GetOneTwinInterface(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer()
	container.Controller.GetOneTwinInterface(g)
}

func CreateTwinInterface(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer()
	container.Controller.CreateTwinInterface(g)
}

func DeleteTwinInterface(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer()
	container.Controller.DeleteTwinInterface(g)
}

// Twin Instances

func GetAllTwinInstances(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer()
	container.Controller.GetAllTwinInstances(g)
}

func GetOneTwinInstance(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer()
	container.Controller.GetOneTwinInstance(g)
}

func CreateTwinInstance(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer()
	container.Controller.CreateTwinInstance(g)
}

func DeleteTwinInstance(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer()
	container.Controller.DeleteTwinInstance(g)
}

// Twin Event

func GetAllTwinEvents(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer()
	container.Controller.GetAllTwinEvents(g)
}

func GetTwinEvents(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer()
	container.Controller.GetTwinEvents(g)
}

func CreateTwinEvent(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer()
	container.Controller.CreateTwinEvent(g)
}

func DeleteTwinEvent(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer()
	container.Controller.DeleteTwinEvent(g)
}
