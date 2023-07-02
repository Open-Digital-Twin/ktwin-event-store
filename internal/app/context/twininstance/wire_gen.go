// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package twininstance

import (
	"agwermann/dt-service/internal/app/context/twininstance/controller"
	"agwermann/dt-service/internal/app/context/twininstance/domain/repository"
	"agwermann/dt-service/internal/app/context/twininstance/usecase"
)

// Injectors from wire.go:

func InitializeTwinInstanceContainer() TwinInstanceContainer {
	twinInstanceUseCase := usecase.NewTwinInstanceUseCase()
	twinInstanceController := controller.NewTwinInstanceController(twinInstanceUseCase)
	twinInstanceRepository := repository.NewTwinInstanceRepository()
	twinInstanceContainer := NewTwinInstanceContainer(twinInstanceController, twinInstanceRepository, twinInstanceUseCase)
	return twinInstanceContainer
}

// wire.go:

func NewTwinInstanceContainer(controller2 controller.TwinInstanceController, repository2 repository.TwinInstanceRepository,

	useCase usecase.TwinInstanceUseCase,
) TwinInstanceContainer {
	return TwinInstanceContainer{
		Controller: controller2,
		Repository: repository2,
		UseCase:    useCase,
	}
}

type TwinInstanceContainer struct {
	Controller controller.TwinInstanceController
	Repository repository.TwinInstanceRepository
	UseCase    usecase.TwinInstanceUseCase
}
