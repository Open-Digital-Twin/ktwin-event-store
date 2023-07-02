//go:build wireinject
// +build wireinject

package twininstance

import (
	"agwermann/dt-service/internal/app/context/twininstance/controller"
	"agwermann/dt-service/internal/app/context/twininstance/domain/repository"
	"agwermann/dt-service/internal/app/context/twininstance/usecase"
	"agwermann/dt-service/internal/app/infra/db"

	"github.com/google/wire"
)

func NewTwinInstanceContainer(
	controller controller.TwinInstanceController,
	repository repository.TwinInstanceRepository,
	useCase usecase.TwinInstanceUseCase,
) TwinInstanceContainer {
	return TwinInstanceContainer{
		Controller: controller,
		Repository: repository,
		UseCase:    useCase,
	}
}

type TwinInstanceContainer struct {
	Controller controller.TwinInstanceController
	Repository repository.TwinInstanceRepository
	UseCase    usecase.TwinInstanceUseCase
}

func InitializeTwinInstanceContainer() TwinInstanceContainer {
	wire.Build(
		NewTwinInstanceContainer,
		controller.NewTwinInstanceController,
		usecase.NewTwinInstanceUseCase,
		repository.NewTwinInstanceRepository,
		repository.NewTwinInstanceMapper,
		db.NewDBConnection,
	)

	return TwinInstanceContainer{}
}
