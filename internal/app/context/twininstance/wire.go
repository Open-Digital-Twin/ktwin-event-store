//go:build wireinject
// +build wireinject

package twininstance

import (
	"github.com/ktwins/event-store/internal/app/context/twininstance/controller"
	"github.com/ktwins/event-store/internal/app/context/twininstance/domain/repository"
	"github.com/ktwins/event-store/internal/app/context/twininstance/usecase"
	"github.com/ktwins/event-store/internal/app/infra/db"
	"github.com/ktwins/event-store/internal/app/infra/validator"

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

		validator.NewValidator,
		db.NewDBConnection,
	)

	return TwinInstanceContainer{}
}
