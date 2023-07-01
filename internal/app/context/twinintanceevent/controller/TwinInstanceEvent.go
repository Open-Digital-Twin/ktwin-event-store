package controller

type TwinInterfaceController interface {
	GetAllTwinInterfaces()
	GetOneTwinInterfaces()
}

func NewTwinInterfaceController() TwinInterfaceController {
	return &twinInterfaceController{}
}

type twinInterfaceController struct {
}

func (t *twinInterfaceController) GetAllTwinInterfaces() {

}

func (t *twinInterfaceController) GetOneTwinInterfaces() {

}
