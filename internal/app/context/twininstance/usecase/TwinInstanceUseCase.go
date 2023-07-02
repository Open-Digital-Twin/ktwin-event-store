package usecase

import "agwermann/dt-service/internal/app/context/twininstance/domain"

func NewTwinInstanceUseCase() TwinInstanceUseCase {
	return &twinInstanceUseCase{}
}

type TwinInstanceUseCase interface {
	GetAllTwinInterfaces()
	GetOneTwinInterfaces(id string)
	DeleteTwinInterface(id string)
	CreateTwinInterface(twinInterface domain.TwinInstance)
}

type twinInstanceUseCase struct{}

func (*twinInstanceUseCase) GetAllTwinInterfaces() {}

func (*twinInstanceUseCase) GetOneTwinInterfaces(id string) {}

func (*twinInstanceUseCase) DeleteTwinInterface(id string) {}

func (*twinInstanceUseCase) CreateTwinInterface(twinInterface domain.TwinInstance) {}
