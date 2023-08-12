package controller

import (
	"net/http"
	"reflect"

	"github.com/ktwins/event-store/internal/app/context/twininstance/domain"
	"github.com/ktwins/event-store/internal/app/context/twininstance/usecase"
	"github.com/ktwins/event-store/internal/app/infra/validator"

	"github.com/gin-gonic/gin"
)

func NewTwinInstanceController(
	twinInstanceUseCase usecase.TwinInstanceUseCase,
	validator validator.Validator,
) TwinInstanceController {
	return &twinInstanceController{
		twinInstanceUseCase: twinInstanceUseCase,
		validator:           validator,
	}
}

type TwinInstanceController interface {
	GetAllTwinInstances(g *gin.Context)
	GetOneTwinInstance(g *gin.Context)
	DeleteTwinInstance(g *gin.Context)
	CreateTwinInstance(g *gin.Context)
}

type twinInstanceController struct {
	twinInstanceUseCase usecase.TwinInstanceUseCase
	validator           validator.Validator
}

// Get All Twin Instances godoc
// @Summary Get All Twin Instances
// @Schemes
// @Description This endpoint returns all Twin Instances registered in the Platform.
// @Tags TwinInstances
// @Accept json
// @Produce json
// @Success 200 {object} []domain.TwinInstance
// @Router /twin-instances [get]
func (tc *twinInstanceController) GetAllTwinInstances(g *gin.Context) {
	twinInstances, err := tc.twinInstanceUseCase.GetAllTwinInstances()

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusOK, twinInstances)
	}
}

// Get Twin Instance by Id godoc
// @Summary Get Twin Instance by Id
// @Schemes
// @Description This endpoint returns the Twin Instance by id.
// @Tags TwinInstances
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinInstance
// @Router /twin-instances/{instanceId}/{id} [get]
func (tc *twinInstanceController) GetOneTwinInstance(g *gin.Context) {
	interfaceId, hasInterfaceId := g.Params.Get("interfaceId")
	id, hasId := g.Params.Get("id")

	if !hasInterfaceId && !hasId {
		g.JSON(http.StatusBadRequest, "Missing interfaceId or id")
		return
	}

	twinInstance, err := tc.twinInstanceUseCase.GetOneTwinInstance(interfaceId, id)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
		return
	} else if reflect.DeepEqual(twinInstance, domain.TwinInstance{}) {
		g.JSON(http.StatusNotFound, "Not Found")
	} else {
		g.JSON(http.StatusOK, twinInstance)
	}
}

// Create Twin Instance
// @Summary Create Twin Instance
// @Schemes
// @Description This endpoint creates the Twin Instance.
// @Tags TwinInstances
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinInstance
// @Router /twin-Instances [post]
func (tc *twinInstanceController) CreateTwinInstance(g *gin.Context) {
	var twinInstance domain.TwinInstance
	err := g.BindJSON(&twinInstance)

	if err != nil {
		g.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	if err := tc.validator.ValidateStruct(&twinInstance); err != nil {
		g.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	err = tc.twinInstanceUseCase.CreateTwinInstance(twinInstance)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusAccepted, "")
	}
}

// Delete Twin Instance
// @Summary Delete Twin Instance
// @Schemes
// @Description This endpoint deletes the Twin Instance.
// @Tags TwinInstances
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinInstance
// @Router /twin-Instances/{instanceId} [delete]
func (tc *twinInstanceController) DeleteTwinInstance(g *gin.Context) {
	interfaceId, hasInterfaceId := g.Params.Get("interfaceId")
	id, hasId := g.Params.Get("id")

	if !hasInterfaceId && !hasId {
		g.JSON(http.StatusBadRequest, "Missing interfaceId or id")
		return
	}

	err := tc.twinInstanceUseCase.DeleteTwinInstance(interfaceId, id)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusOK, "")
	}
}
