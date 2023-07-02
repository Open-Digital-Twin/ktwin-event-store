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
	GetAllTwinInterfaces() ([]domain.TwinInstance, error)
	GetOneTwinInterface(id string) (domain.TwinInstance, error)
	DeleteTwinInterface(id string)
	CreateTwinInterface(twinInterface domain.TwinInstance)
}

type twinInstanceUseCase struct {
	repository repository.TwinInstanceRepository
}

func (t *twinInstanceUseCase) GetAllTwinInterfaces() ([]domain.TwinInstance, error) {
	return t.repository.GetAllTwinInstances()
}

func (t *twinInstanceUseCase) GetOneTwinInterface(id string) (domain.TwinInstance, error) {
	return t.repository.GetOneTwinInstance(id)
}

func (*twinInstanceUseCase) DeleteTwinInterface(id string) {}

func (*twinInstanceUseCase) CreateTwinInterface(twinInterface domain.TwinInstance) {}
