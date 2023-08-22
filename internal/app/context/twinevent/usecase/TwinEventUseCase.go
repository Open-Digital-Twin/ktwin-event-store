package usecase

import (
	"encoding/json"
	"reflect"

	"dario.cat/mergo"
	"github.com/ktwins/event-store/internal/app/context/twinevent/domain"
	"github.com/ktwins/event-store/internal/app/context/twinevent/domain/repository"
)

func NewTwinEventUseCase(
	repository repository.TwinEventRepository,
) TwinEventUseCase {
	return &twinEventUseCase{
		repository: repository,
	}
}

type TwinEventUseCase interface {
	GetAllTwinEvents() ([]domain.TwinEvent, error)
	GetTwinEvents(interfaceId string, instanceId string) ([]domain.TwinEvent, error)
	GetLatestTwinEvent(interfaceId string, instanceId string) (domain.TwinEvent, error)
	CreateTwinEvent(twinInterface domain.TwinEvent) error
	UpdateTwinEvent(twinInterface domain.TwinEvent) error
	DeleteTwinEvent(interfaceId string, instanceId string) error
}

type twinEventUseCase struct {
	repository repository.TwinEventRepository
}

func (t *twinEventUseCase) GetAllTwinEvents() ([]domain.TwinEvent, error) {
	return t.repository.GetAllTwinEvents()
}

func (t *twinEventUseCase) GetTwinEvents(interfaceId string, instanceId string) ([]domain.TwinEvent, error) {
	return t.repository.GetTwinEvents(interfaceId, instanceId)
}

func (t *twinEventUseCase) GetLatestTwinEvent(interfaceId string, instanceId string) (domain.TwinEvent, error) {
	return t.repository.GetLatestTwinEvent(interfaceId, instanceId)
}

func (t *twinEventUseCase) CreateTwinEvent(twinEvent domain.TwinEvent) error {
	return t.repository.CreateTwinEvent(twinEvent)
}

func (t *twinEventUseCase) UpdateTwinEvent(twinEvent domain.TwinEvent) error {
	latestTwinEvent, err := t.repository.GetLatestTwinEvent(twinEvent.InterfaceId, twinEvent.InstanceId)

	if err != nil {
		return err
	}

	if !reflect.DeepEqual(latestTwinEvent, domain.TwinEvent{}) {
		var latestTwinEventData map[string]interface{}
		var twinEventData map[string]interface{}

		err = json.Unmarshal(latestTwinEvent.EventData, &latestTwinEventData)

		if err != nil {
			return err
		}

		err = json.Unmarshal(twinEvent.EventData, &twinEventData)

		if err != nil {
			return err
		}

		err = mergo.MergeWithOverwrite(&latestTwinEventData, &twinEventData)

		if err != nil {
			return err
		}

		mergedEventData, err := json.Marshal(latestTwinEventData)

		if err != nil {
			return err
		}

		twinEvent.EventData = mergedEventData
	}

	return t.repository.CreateTwinEvent(twinEvent)
}

func (t *twinEventUseCase) DeleteTwinEvent(interfaceId string, id string) error {
	return t.repository.DeleteTwinEvent(interfaceId, interfaceId)
}
