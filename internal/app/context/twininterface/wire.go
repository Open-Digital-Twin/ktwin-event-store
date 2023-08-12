//go:build wireinject
// +build wireinject

package twininterface

import (
	"github.com/ktwins/event-store/internal/app/context/twininterface/controller"
	"github.com/ktwins/event-store/internal/app/context/twininterface/domain/repository"
	"github.com/ktwins/event-store/internal/app/context/twininterface/usecase"
	"github.com/ktwins/event-store/internal/app/infra/db"
	"github.com/ktwins/event-store/internal/app/infra/validator"

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
		repository.NewTwinInterfaceMapper,

		validator.NewValidator,
		db.NewDBConnection,
	)

	return TwinInterfaceContainer{}
}
