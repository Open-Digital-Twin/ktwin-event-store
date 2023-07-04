// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package twininterface

import (
	"agwermann/dt-service/internal/app/context/twininterface/controller"
	"agwermann/dt-service/internal/app/context/twininterface/domain/repository"
	"agwermann/dt-service/internal/app/context/twininterface/usecase"
	"agwermann/dt-service/internal/app/infra/validator"
)

// Injectors from wire.go:

func InitializeTwinInterfaceContainer() TwinInterfaceContainer {
	twinInterfaceRepository := repository.NewTwinInterfaceRepository()
	twinInterfaceUseCase := usecase.NewTwinInterfaceUseCase(twinInterfaceRepository)
	validatorValidator := validator.NewValidator()
	twinInterfaceController := controller.NewTwinInterfaceController(twinInterfaceUseCase, validatorValidator)
	twinInterfaceContainer := NewTwinInterfaceContainer(twinInterfaceController, twinInterfaceRepository, twinInterfaceUseCase)
	return twinInterfaceContainer
}

// wire.go:

func NewTwinInterfaceContainer(controller2 controller.TwinInterfaceController, repository2 repository.TwinInterfaceRepository,

	useCase usecase.TwinInterfaceUseCase,
) TwinInterfaceContainer {
	return TwinInterfaceContainer{
		Controller: controller2,
		Repository: repository2,
		UseCase:    useCase,
	}
}

type TwinInterfaceContainer struct {
	Controller controller.TwinInterfaceController
	Repository repository.TwinInterfaceRepository
	UseCase    usecase.TwinInterfaceUseCase
}
