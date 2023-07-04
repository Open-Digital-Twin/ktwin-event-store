package controller

import (
	"agwermann/dt-service/internal/app/context/twinevent/domain"
	"agwermann/dt-service/internal/app/context/twinevent/usecase"
	"agwermann/dt-service/internal/app/infra/validator"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type TwinEventController interface {
	GetAllTwinEvents(g *gin.Context)
	GetTwinEvents(g *gin.Context)
	CreateTwinEvent(g *gin.Context)
	DeleteTwinEvent(g *gin.Context)
}

func NewTwinEventController(
	twinEventUseCase usecase.TwinEventUseCase,
	validator validator.Validator,
) TwinEventController {
	return &twinEventController{
		twinEventUseCase: twinEventUseCase,
		validator:        validator,
	}
}

type twinEventController struct {
	twinEventUseCase usecase.TwinEventUseCase
	validator        validator.Validator
}

// Get Twin Event Events godoc
// @Summary Get Twin Events
// @Schemes
// @Description do ping
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-events [get]
func (t *twinEventController) GetAllTwinEvents(g *gin.Context) {
	twinEvents, err := t.twinEventUseCase.GetAllTwinEvents()

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
		return
	} else if reflect.DeepEqual(twinEvents, domain.TwinEvent{}) {
		g.JSON(http.StatusNotFound, "Not Found")
	} else {
		g.JSON(http.StatusOK, twinEvents)
	}
}

// Get Twin Event Events godoc
// @Summary Get Twin Events
// @Schemes
// @Description do ping
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-events/{interfaceId}/{instanceId} [get]
func (t *twinEventController) GetTwinEvents(g *gin.Context) {
	interfaceId, hasInterfaceId := g.Params.Get("interfaceId")
	instanceId, hasInstanceId := g.Params.Get("instanceId")

	if !hasInterfaceId && !hasInstanceId {
		g.JSON(http.StatusBadRequest, "Missing interfaceId or id")
		return
	}

	twinInstance, err := t.twinEventUseCase.GetTwinEvents(interfaceId, instanceId)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
		return
	} else if reflect.DeepEqual(twinInstance, domain.TwinEvent{}) {
		g.JSON(http.StatusNotFound, "Not Found")
	} else {
		g.JSON(http.StatusOK, twinInstance)
	}
}

// Create Twin Event godoc
// @Summary Create Twin Event
// @Schemes
// @Description This endpoint populates a Twin Event Event.
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-events [post]
func (t *twinEventController) CreateTwinEvent(g *gin.Context) {
	var twinEvent domain.TwinEvent
	err := g.BindJSON(&twinEvent)

	if err != nil {
		g.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	if err := t.validator.ValidateStruct(&twinEvent); err != nil {
		g.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	err = t.twinEventUseCase.CreateTwinEvent(twinEvent)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusAccepted, "")
	}
}

// Delete Twin Event godoc
// @Summary Delete Twin Event Event
// @Schemes
// @Description This endpoint deletes the Twin Event Event.
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200 {string} Not Implemented
// @Router /twin-events/{interfaceId}/{instanceId} [delete]
func (t *twinEventController) DeleteTwinEvent(g *gin.Context) {
	interfaceId, hasInterfaceId := g.Params.Get("interfaceId")
	instanceId, hasInstanceId := g.Params.Get("instanceId")

	if !hasInterfaceId && !hasInstanceId {
		g.JSON(http.StatusBadRequest, "Missing interfaceId or instanceId")
		return
	}

	err := t.twinEventUseCase.DeleteTwinEvent(interfaceId, instanceId)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusOK, "")
	}
}
