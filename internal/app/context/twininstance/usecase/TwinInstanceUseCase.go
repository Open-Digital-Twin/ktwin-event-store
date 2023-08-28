package usecase

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininstance/domain"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininstance/domain/repository"
)

func NewTwinInstanceUseCase(
	repository repository.TwinInstanceRepository,
) TwinInstanceUseCase {
	return &twinInstanceUseCase{
		repository: repository,
	}
}

type TwinInstanceUseCase interface {
	GetAllTwinInstances() ([]domain.TwinInstance, error)
	GetOneTwinInstance(interfaceId string, id string) (domain.TwinInstance, error)
	DeleteTwinInstance(interfaceId string, id string) error
	CreateTwinInstance(twinInterface domain.TwinInstance) error
}

type twinInstanceUseCase struct {
	repository repository.TwinInstanceRepository
}

func (t *twinInstanceUseCase) GetAllTwinInstances() ([]domain.TwinInstance, error) {
	return t.repository.GetAllTwinInstances()
}

func (t *twinInstanceUseCase) GetOneTwinInstance(interfaceId string, id string) (domain.TwinInstance, error) {
	return t.repository.GetOneTwinInstance(interfaceId, id)
}

func (t *twinInstanceUseCase) DeleteTwinInstance(interfaceId string, id string) error {
	return t.repository.DeleteTwinInstance(interfaceId, id)
}

func (t *twinInstanceUseCase) CreateTwinInstance(twinInterface domain.TwinInstance) error {
	return t.repository.InsertTwinInstance(twinInterface)
}
