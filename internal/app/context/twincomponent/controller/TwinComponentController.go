package controller

type TwinComponentController interface {
	GetAllTwinComponents()
	GetOneTwinComponents()
}

func NewTwinComponentController() TwinComponentController {
	return &twinComponentController{}
}

type twinComponentController struct {
}

func (t *twinComponentController) GetAllTwinComponents() {

}

func (t *twinComponentController) GetOneTwinComponents() {

}
