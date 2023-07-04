package repository

import "agwermann/dt-service/internal/app/context/twininterface/domain"

type TwinInterfaceRepository interface {
	GetAllTwinInterfaces() ([]domain.TwinInterface, error)
	GetOneTwinInterface(id string) (domain.TwinInterface, error)
	DeleteTwinInterface(id string) error
	CreateTwinInterface(twinInterface domain.TwinInterface) error
}

func NewTwinInterfaceRepository() TwinInterfaceRepository {
	return &twinInterfaceRepository{}
}

type twinInterfaceRepository struct {
}

func (*twinInterfaceRepository) GetAllTwinInterfaces() ([]domain.TwinInterface, error) {
	return []domain.TwinInterface{}, nil
}

func (*twinInterfaceRepository) GetOneTwinInterface(id string) (domain.TwinInterface, error) {
	return domain.TwinInterface{}, nil
}

func (t *twinInterfaceRepository) DeleteTwinInterface(id string) error {
	return nil
}

func (t *twinInterfaceRepository) CreateTwinInterface(twinInterface domain.TwinInterface) error {
	return nil
}
