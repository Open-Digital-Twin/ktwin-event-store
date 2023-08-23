package controller

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/ktwins/event-store/internal/app/context/twinevent/domain"
	"github.com/ktwins/event-store/internal/app/context/twinevent/usecase"
	"github.com/ktwins/event-store/internal/app/infra/validator"

	cloudEvents "github.com/cloudevents/sdk-go/v2"

	"github.com/gin-gonic/gin"
)

type TwinEventController interface {
	GetAllTwinEvents(g *gin.Context)
	GetTwinEvents(g *gin.Context)
	GetLatestTwinEvent(g *gin.Context)
	CreateTwinEvent(g *gin.Context)
	UpdateTwinEvent(g *gin.Context)
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
// @Success 200 {object} []domain.TwinEvent
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
// @Success 200 {object} []domain.TwinEvent
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

// Get Twin Event Events godoc
// @Summary Get Twin Events
// @Schemes
// @Description do ping
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200 {object} []domain.TwinEvent
// @Router /twin-events/{interfaceId}/{instanceId} [get]
func (t *twinEventController) GetLatestTwinEvent(g *gin.Context) {
	interfaceId, hasInterfaceId := g.Params.Get("interfaceId")
	instanceId, hasInstanceId := g.Params.Get("instanceId")

	if !hasInterfaceId && !hasInstanceId {
		g.JSON(http.StatusBadRequest, "Missing interfaceId or id")
		return
	}

	twinEvent, err := t.twinEventUseCase.GetLatestTwinEvent(interfaceId, instanceId)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
		return
	} else if reflect.DeepEqual(twinEvent, domain.TwinEvent{}) {
		g.JSON(http.StatusNotFound, "Not Found")
	} else {
		var response map[string]interface{}
		err = json.Unmarshal(twinEvent.EventData, &response)
		if err != nil {
			g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		g.Header("ce-id", twinEvent.Id)
		g.Header("ce-type", twinEvent.Type)
		g.Header("ce-source", twinEvent.Source)
		g.Header("ce-timestamp", twinEvent.Time.String())
		g.Header("ce-specversion", "1.0")
		g.JSON(http.StatusOK, response)
	}
}

// Create Twin Event godoc
// @Summary Create Twin Event
// @Schemes
// @Description This endpoint populates a Twin Event.
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinEvent
// @Router /twin-events [post]
func (t *twinEventController) CreateTwinEvent(g *gin.Context) {
	cloudEvent, err := cloudEvents.NewEventFromHTTPRequest(g.Request)

	if err != nil {
		g.JSON(http.StatusBadRequest, "Invalid Cloud event: "+err.Error())
		return
	}

	twinEvent := domain.TwinEvent{
		Id:          cloudEvent.ID(),
		Time:        cloudEvent.Time(),
		Source:      cloudEvent.Source(),
		Type:        cloudEvent.Type(),
		InterfaceId: extractInstanceId(cloudEvent.Type()),
		InstanceId:  cloudEvent.Source(),
		EventData:   cloudEvent.Data(),
		CreatedAt:   cloudEvent.Time(),
	}

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

// Update Twin Event godoc
// @Summary Update Twin Event
// @Schemes
// @Description This endpoint updates the latest TwinEvent.
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinEvent
// @Router /twin-events [post]
func (t *twinEventController) UpdateTwinEvent(g *gin.Context) {
	cloudEvent, err := cloudEvents.NewEventFromHTTPRequest(g.Request)

	if err != nil {
		g.JSON(http.StatusBadRequest, "Invalid Cloud event: "+err.Error())
		return
	}

	twinEvent := domain.TwinEvent{
		Id:          cloudEvent.ID(),
		Time:        cloudEvent.Time(),
		Source:      cloudEvent.Source(),
		Type:        cloudEvent.Type(),
		InterfaceId: extractInstanceId(cloudEvent.Type()),
		InstanceId:  cloudEvent.Source(),
		EventData:   cloudEvent.Data(),
		CreatedAt:   cloudEvent.Time(),
	}

	if err != nil {
		g.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	if err := t.validator.ValidateStruct(&twinEvent); err != nil {
		g.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	err = t.twinEventUseCase.UpdateTwinEvent(twinEvent)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		g.JSON(http.StatusAccepted, "")
	}
}

func extractInstanceId(eventType string) string {
	eType := strings.Split(eventType, ".")

	if len(eType) > 2 {
		return eType[2]
	}

	return eventType
}

// Delete Twin Event godoc
// @Summary Delete Twin Event Event
// @Schemes
// @Description This endpoint deletes the Twin Event Event.
// @Tags TwinEvents
// @Accept json
// @Produce json
// @Success 200
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
