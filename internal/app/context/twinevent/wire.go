//go:build wireinject
// +build wireinject

package twinevent

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twinevent/controller"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twinevent/domain/repository"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twinevent/usecase"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/db"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/validator"

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

func InitializeTwinEventContainer(dbConnection db.DBConnection) TwinEventContainer {
	wire.Build(
		NewTwinEventContainer,
		controller.NewTwinEventController,
		usecase.NewTwinEventUseCase,
		repository.NewTwinEventRepository,
		repository.NewTwinEventMapper,

		validator.NewValidator,
	)

	return TwinEventContainer{}
}
