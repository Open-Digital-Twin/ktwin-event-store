package controller

import (
	"agwermann/dt-service/internal/app/context/twininterface/domain"
	"agwermann/dt-service/internal/app/context/twininterface/usecase"
	"agwermann/dt-service/internal/app/infra/validator"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type TwinInterfaceController interface {
	GetAllTwinInterfaces(g *gin.Context)
	GetOneTwinInterface(g *gin.Context)
	CreateTwinInterface(g *gin.Context)
	DeleteTwinInterface(g *gin.Context)
}

func NewTwinInterfaceController(
	twinInterfaceUseCase usecase.TwinInterfaceUseCase,
	validator validator.Validator,
) TwinInterfaceController {
	return &twinInterfaceController{
		twinInterfaceUseCase: twinInterfaceUseCase,
		validator:            validator,
	}
}

type twinInterfaceController struct {
	twinInterfaceUseCase usecase.TwinInterfaceUseCase
	validator            validator.Validator
}

// Get All Twin Interfaces godoc
// @Summary Get All Twin Interfaces
// @Schemes
// @Description This endpoint returns all Twin Interfaces registered in the Platform.
// @Tags TwinInterfaces
// @Accept json
// @Produce json
// @Success 200 {object} []domain.TwinInterface
// @Router /twin-interfaces [get]
func (t *twinInterfaceController) GetAllTwinInterfaces(g *gin.Context) {
	twinInterfaces, err := t.twinInterfaceUseCase.GetAllTwinInterfaces()

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusOK, twinInterfaces)
	}
}

// Get Twin Interface by Id godoc
// @Summary Get Twin Interface by Id
// @Schemes
// @Description This endpoint returns the Twin Interface by id.
// @Tags TwinInterfaces
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinInterface
// @Router /twin-interfaces/{id} [get]
func (t *twinInterfaceController) GetOneTwinInterface(g *gin.Context) {
	id, hasId := g.Params.Get("id")

	if !hasId {
		g.JSON(http.StatusBadRequest, "Missing interfaceId or id")
		return
	}

	twinInterface, err := t.twinInterfaceUseCase.GetOneTwinInterface(id)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
		return
	} else if reflect.DeepEqual(twinInterface, domain.TwinInterface{}) {
		g.JSON(http.StatusNotFound, "Not Found")
	} else {
		g.JSON(http.StatusOK, twinInterface)
	}
}

// Create Twin Interface
// @Summary Create Twin Interface
// @Schemes
// @Description This endpoint creates the Twin Interface.
// @Tags TwinInterfaces
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinInterface
// @Router /twin-interfaces [post]
func (t *twinInterfaceController) CreateTwinInterface(g *gin.Context) {
	var twinInterface domain.TwinInterface
	err := g.BindJSON(&twinInterface)

	if err != nil {
		g.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	if err := t.validator.ValidateStruct(&twinInterface); err != nil {
		g.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	err = t.twinInterfaceUseCase.CreateTwinInterface(twinInterface)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusAccepted, "")
	}
}

// Delete Twin Interface
// @Summary Delete Twin Interface
// @Schemes
// @Description This endpoint deletes the Twin Interface.
// @Tags TwinInterfaces
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinInterface
// @Router /twin-interfaces/{instanceId} [delete]
func (t *twinInterfaceController) DeleteTwinInterface(g *gin.Context) {
	id, hasId := g.Params.Get("id")

	if !hasId {
		g.JSON(http.StatusBadRequest, "Missing interfaceId or id")
		return
	}

	err := t.twinInterfaceUseCase.DeleteTwinInterface(id)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusOK, "")
	}
}
