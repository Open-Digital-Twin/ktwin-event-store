//go:build wireinject
// +build wireinject

package twinevent

import (
	"agwermann/event-store-service/internal/app/context/twinevent/controller"
	"agwermann/event-store-service/internal/app/context/twinevent/domain/repository"
	"agwermann/event-store-service/internal/app/context/twinevent/usecase"
	"agwermann/event-store-service/internal/app/infra/db"
	"agwermann/event-store-service/internal/app/infra/validator"

	"github.com/google/wire"
)

func NewTwinEventContainer(
	controller controller.TwinEventController,
	repository repository.TwinEventRepository,
	useCase usecase.TwinEventUseCase,
) TwinEventContainer {
	return TwinEventContainer{
		Controller: controller,
		Repository: repository,
		UseCase:    useCase,
	}
}

type TwinEventContainer struct {
	Controller controller.TwinEventController
	Repository repository.TwinEventRepository
	UseCase    usecase.TwinEventUseCase
}

func InitializeTwinEventContainer() TwinEventContainer {
	wire.Build(
		NewTwinEventContainer,
		controller.NewTwinEventController,
		usecase.NewTwinEventUseCase,
		repository.NewTwinEventRepository,
		repository.NewTwinEventMapper,

		validator.NewValidator,
		db.NewDBConnection,
	)

	return TwinEventContainer{}
}
