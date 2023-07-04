//go:build wireinject
// +build wireinject

package twininterface

import (
	"agwermann/dt-service/internal/app/context/twininterface/controller"
	"agwermann/dt-service/internal/app/context/twininterface/domain/repository"
	"agwermann/dt-service/internal/app/context/twininterface/usecase"
	"agwermann/dt-service/internal/app/infra/validator"

	"github.com/google/wire"
)

func NewTwinInterfaceContainer(
	controller controller.TwinInterfaceController,
	repository repository.TwinInterfaceRepository,
	useCase usecase.TwinInterfaceUseCase,
) TwinInterfaceContainer {
	return TwinInterfaceContainer{
		Controller: controller,
		Repository: repository,
		UseCase:    useCase,
	}
}

type TwinInterfaceContainer struct {
	Controller controller.TwinInterfaceController
	Repository repository.TwinInterfaceRepository
	UseCase    usecase.TwinInterfaceUseCase
}

func InitializeTwinInterfaceContainer() TwinInterfaceContainer {
	wire.Build(
		NewTwinInterfaceContainer,
		controller.NewTwinInterfaceController,
		usecase.NewTwinInterfaceUseCase,
		repository.NewTwinInterfaceRepository,
		// repository.NewTwinInterfaceMapper,

		validator.NewValidator,
		// db.NewDBConnection,
	)

	return TwinInterfaceContainer{}
}
