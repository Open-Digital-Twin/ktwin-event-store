package server

import (
	"agwermann/dt-service/internal/app/context/twininstance"
	"agwermann/dt-service/internal/app/context/twininterface"

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

// Twin Instance Events

// Get All Twin Instance Events godoc
// @Summary Get All Twin Instances Events
// @Schemes
// @Description do ping
// @Tags TwinInstanceEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-instances/{instanceId}/events [get]
// func GetAllTwinInstanceEvents(g *gin.Context) {
// 	g.JSON(http.StatusNotImplemented, "Not Implemented")
// }

// Get Twin Instance godoc
// @Summary Get One Twin Instance Event
// @Schemes
// @Description do ping
// @Tags TwinInstanceEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-instances/{instanceId}/events/{eventId} [get]
// func GetOneTwinInstancesEvent(g *gin.Context) {
// 	g.JSON(http.StatusNotImplemented, "Not Implemented")
// }

// Create Twin Instance godoc
// @Summary Create Twin Instance Event
// @Schemes
// @Description This endpoint populates a Twin Instance Event.
// @Tags TwinInstanceEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-instances/{instanceId}/events [post]
// func CreateTwinInstanceEvent(g *gin.Context) {
// 	g.JSON(http.StatusNotImplemented, "Not Implemented")
// }

// Delete Twin Instance godoc
// @Summary Delete Twin Instance Event
// @Schemes
// @Description This endpoint deletes the Twin Instance Event.
// @Tags TwinInstanceEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-instances/{instanceId}/events/{eventId} [delete]
// func DeleteTwinInstanceEvent(g *gin.Context) {
// 	g.JSON(http.StatusNotImplemented, "Not Implemented")
// }
