package usecase

func NewTwinInstanceUseCase() TwinInstanceUseCase {
	return &twinInstanceUseCase{}
}

type TwinInstanceUseCase interface {
	GetAllTwinInterfaces()
	GetOneTwinInterfaces(id string)
	DeleteTwinInterface(id string)
	CreateTwinInterface(twinInterface TwinInstance)
}

type twinInstanceUseCase struct{}

func (*twinInstanceUseCase) GetAllTwinInterfaces() {}

func (*twinInstanceUseCase) GetOneTwinInterfaces(id string) {}

func (*twinInstanceUseCase) DeleteTwinInterface(id string) {}

func (*twinInstanceUseCase) CreateTwinInterface(twinInterface TwinInstance) {}
