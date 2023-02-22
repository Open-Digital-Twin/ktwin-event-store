package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All Twin Components godoc
// @Summary Get All Twin Components
// @Schemes
// @Description This endpoint returns all Twin Components registered in the Platform.
// @Tags TwinComponent
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-components [get]
func GetAllTwinComponents(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Get Twin Component by Id godoc
// @Summary Get Twin Component by Id
// @Schemes
// @Description This endpoint returns the Twin Component by id.
// @Tags TwinComponent
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-components/{componentId} [get]
func GetOneTwinComponents(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Create Twin Component
// @Summary Create Twin Component
// @Schemes
// @Description This endpoint creates the Twin Component.
// @Tags TwinComponent
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-components [post]
func CreateTwinComponent(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Update Twin Component
// @Summary Update Twin Component
// @Schemes
// @Description This endpoint updates the Twin Component.
// @Tags TwinComponent
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-components [put]
func UpdateTwinComponent(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Delete Twin Component
// @Summary Delete Twin Component
// @Schemes
// @Description This endpoint deletes the Twin Component.
// @Tags TwinComponent
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-components/{componentId} [delete]
func DeleteTwinComponent(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Get Twin Component Classes
// @Summary Get Twin Component Classes
// @Schemes
// @Description This endpoint get the Twin Component classes.
// @Tags TwinComponent
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-components/{componentId}/classes [get]
func GetTwinComponentClasses(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Get Twin Component Class by Id
// @Summary Get Twin Component Class by Id
// @Schemes
// @Description This endpoint get the Twin Component class by Id.
// @Tags TwinComponent
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /twin-components/{componentId}/classes/{classId} [get]
func GetOneTwinComponentClass(g *gin.Context) {
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
