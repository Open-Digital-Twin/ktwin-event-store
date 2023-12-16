//go:build wireinject
// +build wireinject

package twininterface

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininterface/controller"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininterface/domain/repository"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininterface/usecase"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/db"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/validator"

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

func InitializeTwinInterfaceContainer(dbConnection db.DBConnection) TwinInterfaceContainer {
	wire.Build(
		NewTwinInterfaceContainer,
		controller.NewTwinInterfaceController,
		usecase.NewTwinInterfaceUseCase,
		repository.NewTwinInterfaceRepository,
		repository.NewTwinInterfaceMapper,

		validator.NewValidator,
	)

	return TwinInterfaceContainer{}
}
