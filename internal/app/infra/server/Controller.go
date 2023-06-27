package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All Twin Interfaces godoc
// @Summary Get All Twin Interfaces
// @Schemes
// @Description This endpoint returns all Twin Interfaces registered in the Platform.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-interfaces [get]
func GetAllTwinInterfaces(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Get Twin Interface by Id godoc
// @Summary Get Twin Interface by Id
// @Schemes
// @Description This endpoint returns the Twin Interface by id.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-interfaces/{interfaceId} [get]
func GetOneTwinInterfaces(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Create Twin Interface
// @Summary Create Twin Interface
// @Schemes
// @Description This endpoint creates the Twin Interface.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-interfaces [post]
func CreateTwinInterface(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Update Twin Interface
// @Summary Update Twin Interface
// @Schemes
// @Description This endpoint updates the Twin Interface.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-interfaces [put]
func UpdateTwinInterface(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Delete Twin Interface
// @Summary Delete Twin Interface
// @Schemes
// @Description This endpoint deletes the Twin Interface.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-interfaces/{interfaceId} [delete]
func DeleteTwinInterface(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Get Twin Interface Classes
// @Summary Get Twin Interface Classes
// @Schemes
// @Description This endpoint get the Twin Interface classes.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-interfaces/{interfaceId}/classes [get]
func GetTwinInterfaceClasses(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Get Twin Interface Class by Id
// @Summary Get Twin Interface Class by Id
// @Schemes
// @Description This endpoint get the Twin Interface class by Id.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-interfaces/{interfaceId}/classes/{classId} [get]
func GetOneTwinInterfaceClass(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Twin Instances

// Get All Twin Instances godoc
// @Summary Get All Twin Instances
// @Schemes
// @Description do ping
// @Tags TwinInstance
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-instances [get]
func GetAllTwinInstances(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Get Twin Instance godoc
// @Summary Get One Twin Instance
// @Schemes
// @Description do ping
// @Tags TwinInstance
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-instances/{instanceId} [get]
func GetOneTwinInstances(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Create Twin Instance godoc
// @Summary Create Twin Instance
// @Schemes
// @Description This endpoint creates a Twin Instance.
// @Tags TwinInstance
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-instances [post]
func CreateTwinInstance(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Update Twin Instance godoc
// @Summary Update Twin Instance
// @Schemes
// @Description This endpoint updates the Twin Instance.
// @Tags TwinInstance
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-instances [put]
func UpdateTwinInstance(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Delete Twin Instance godoc
// @Summary Delete Twin Instance
// @Schemes
// @Description This endpoint deletes the Twin Instance.
// @Tags TwinInstance
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-instances/{instanceId} [delete]
func DeleteTwinInstance(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
