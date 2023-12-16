package server

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twinevent"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininstance"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininterface"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/db"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetAllTwinInterfaces(g *gin.Context)
	GetOneTwinInterface(g *gin.Context)
	CreateTwinInterface(g *gin.Context)
	DeleteTwinInterface(g *gin.Context)
	GetAllTwinInstances(g *gin.Context)
	GetOneTwinInstance(g *gin.Context)
	CreateTwinInstance(g *gin.Context)
	DeleteTwinInstance(g *gin.Context)
	GetAllTwinEvents(g *gin.Context)
	GetTwinEvents(g *gin.Context)
	GetLatestTwinEvent(g *gin.Context)
	CreateTwinEvent(g *gin.Context)
	UpdateTwinEvent(g *gin.Context)
	DeleteTwinEvent(g *gin.Context)
}

func NewAppController(dbConnection db.DBConnection) Controller {
	return &controller{dbConnection: dbConnection}
}

type controller struct {
	dbConnection db.DBConnection
}

// Twin Interfaces

func (c *controller) GetAllTwinInterfaces(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer(c.dbConnection)
	container.Controller.GetAllTwinInterfaces(g)
}

func (c *controller) GetOneTwinInterface(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer(c.dbConnection)
	container.Controller.GetOneTwinInterface(g)
}

func (c *controller) CreateTwinInterface(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer(c.dbConnection)
	container.Controller.CreateTwinInterface(g)
}

func (c *controller) DeleteTwinInterface(g *gin.Context) {
	container := twininterface.InitializeTwinInterfaceContainer(c.dbConnection)
	container.Controller.DeleteTwinInterface(g)
}

// Twin Instances

func (c *controller) GetAllTwinInstances(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer(c.dbConnection)
	container.Controller.GetAllTwinInstances(g)
}

func (c *controller) GetOneTwinInstance(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer(c.dbConnection)
	container.Controller.GetOneTwinInstance(g)
}

func (c *controller) CreateTwinInstance(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer(c.dbConnection)
	container.Controller.CreateTwinInstance(g)
}

func (c *controller) DeleteTwinInstance(g *gin.Context) {
	container := twininstance.InitializeTwinInstanceContainer(c.dbConnection)
	container.Controller.DeleteTwinInstance(g)
}

// Twin Event

func (c *controller) GetAllTwinEvents(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer(c.dbConnection)
	container.Controller.GetAllTwinEvents(g)
}

func (c *controller) GetTwinEvents(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer(c.dbConnection)
	container.Controller.GetTwinEvents(g)
}

func (c *controller) GetLatestTwinEvent(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer(c.dbConnection)
	container.Controller.GetLatestTwinEvent(g)
}

func (c *controller) CreateTwinEvent(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer(c.dbConnection)
	container.Controller.CreateTwinEvent(g)
}

func (c *controller) UpdateTwinEvent(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer(c.dbConnection)
	container.Controller.UpdateTwinEvent(g)
}

func (c *controller) DeleteTwinEvent(g *gin.Context) {
	container := twinevent.InitializeTwinEventContainer(c.dbConnection)
	container.Controller.DeleteTwinEvent(g)
}
