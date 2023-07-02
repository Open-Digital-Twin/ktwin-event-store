package usecase

import (
	"agwermann/dt-service/internal/app/context/twininstance/domain"
	"agwermann/dt-service/internal/app/context/twininstance/domain/repository"
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
	GetOneTwinInstance(id string) (domain.TwinInstance, error)
	DeleteTwinInstance(id string)
	CreateTwinInstance(twinInterface domain.TwinInstance) error
}

type twinInstanceUseCase struct {
	repository repository.TwinInstanceRepository
}

func (t *twinInstanceUseCase) GetAllTwinInstances() ([]domain.TwinInstance, error) {
	return t.repository.GetAllTwinInstances()
}

func (t *twinInstanceUseCase) GetOneTwinInstance(id string) (domain.TwinInstance, error) {
	return t.repository.GetOneTwinInstance(id)
}

func (*twinInstanceUseCase) DeleteTwinInstance(id string) {}

func (t *twinInstanceUseCase) CreateTwinInstance(twinInterface domain.TwinInstance) error {
	return t.repository.InsertTwinInstance(twinInterface)
}
