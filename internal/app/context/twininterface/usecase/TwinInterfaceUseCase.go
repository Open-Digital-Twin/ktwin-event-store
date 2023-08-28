package usecase

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininterface/domain"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininterface/domain/repository"
)

func NewTwinInterfaceUseCase(
	repository repository.TwinInterfaceRepository,
) TwinInterfaceUseCase {
	return &twinInstanceUseCase{
		repository: repository,
	}
}

type TwinInterfaceUseCase interface {
	GetAllTwinInterfaces() ([]domain.TwinInterface, error)
	GetOneTwinInterface(id string) (domain.TwinInterface, error)
	DeleteTwinInterface(id string) error
	CreateTwinInterface(twinInterface domain.TwinInterface) error
}

type twinInstanceUseCase struct {
	repository repository.TwinInterfaceRepository
}

func (t *twinInstanceUseCase) GetAllTwinInterfaces() ([]domain.TwinInterface, error) {
	return t.repository.GetAllTwinInterfaces()
}

func (t *twinInstanceUseCase) GetOneTwinInterface(id string) (domain.TwinInterface, error) {
	return t.repository.GetOneTwinInterface(id)
}

func (t *twinInstanceUseCase) DeleteTwinInterface(id string) error {
	return t.repository.DeleteTwinInterface(id)
}

func (t *twinInstanceUseCase) CreateTwinInterface(twinInterface domain.TwinInterface) error {
	return t.repository.CreateTwinInterface(twinInterface)
}
