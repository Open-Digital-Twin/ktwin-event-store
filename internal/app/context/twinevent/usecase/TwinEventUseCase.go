package usecase

import (
	"agwermann/dt-service/internal/app/context/twinevent/domain"
	"agwermann/dt-service/internal/app/context/twinevent/domain/repository"
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
	CreateTwinEvent(twinInterface domain.TwinEvent) error
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

func (t *twinEventUseCase) CreateTwinEvent(twinEvent domain.TwinEvent) error {
	return t.repository.CreateTwinEvent(twinEvent)
}

func (t *twinEventUseCase) DeleteTwinEvent(interfaceId string, id string) error {
	return t.repository.DeleteTwinEvent(interfaceId, interfaceId)
}
