package controller

import (
	"agwermann/dt-service/internal/app/context/twininstance/domain"
	"agwermann/dt-service/internal/app/context/twininstance/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewTwinInstanceController(
	twinInstanceUseCase usecase.TwinInstanceUseCase,
) TwinInstanceController {
	return &twinInstanceController{
		twinInstanceUseCase: twinInstanceUseCase,
	}
}

type TwinInstanceController interface {
	GetAllTwinInterfaces(g *gin.Context)
	GetOneTwinInterfaces(g *gin.Context)
	DeleteTwinInterface(g *gin.Context)
	CreateTwinInterface(g *gin.Context)
}

type twinInstanceController struct {
	twinInstanceUseCase usecase.TwinInstanceUseCase
}

// Get All Twin Interfaces godoc
// @Summary Get All Twin Interfaces
// @Schemes
// @Description This endpoint returns all Twin Interfaces registered in the Platform.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-interfaces [get]
func (tc *twinInstanceController) GetAllTwinInterfaces(g *gin.Context) {
	twinInstance, err := tc.twinInstanceUseCase.GetAllTwinInterfaces()

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusOK, twinInstance)
	}
}

// Get Twin Interface by Id godoc
// @Summary Get Twin Interface by Id
// @Schemes
// @Description This endpoint returns the Twin Interface by id.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-interfaces/{interfaceId} [get]
func (tc *twinInstanceController) GetOneTwinInterfaces(g *gin.Context) {
	tc.twinInstanceUseCase.GetOneTwinInterfaces("")
	g.JSON(http.StatusNotImplemented, "Not Implemented")
}

// Create Twin Interface
// @Summary Create Twin Interface
// @Schemes
// @Description This endpoint creates the Twin Interface.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-interfaces [post]
func (tc *twinInstanceController) CreateTwinInterface(g *gin.Context) {
	tc.twinInstanceUseCase.CreateTwinInterface(domain.TwinInstance{})
	g.JSON(http.StatusNotImplemented, "Not Implemented")
}

// Delete Twin Interface
// @Summary Delete Twin Interface
// @Schemes
// @Description This endpoint deletes the Twin Interface.
// @Tags TwinInterface
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-interfaces/{interfaceId} [delete]
func (tc *twinInstanceController) DeleteTwinInterface(g *gin.Context) {
	tc.twinInstanceUseCase.DeleteTwinInterface("")
	g.JSON(http.StatusNotImplemented, "Not Implemented")
}
