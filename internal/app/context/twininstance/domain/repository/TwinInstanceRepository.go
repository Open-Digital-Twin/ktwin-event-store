package repository

import (
	"agwermann/dt-service/internal/app/context/twininstance/usecase"
)

func NewTwinInstanceRepository() TwinInstanceRepository {
	return &twinInstanceRepository{}
}

type TwinInstanceRepository interface {
	GetAllTwinInstances()
	GetTwinInstance(id string) usecase.TwinInstance
	InsertTwinInstance(twinInstance usecase.TwinInstance) error
	DeleteTwinInstance(id string) error
}

type twinInstanceRepository struct {
}

func (*twinInstanceRepository) GetAllTwinInstances() {
	return
}

func (*twinInstanceRepository) GetTwinInstance(id string) usecase.TwinInstance {
	return usecase.TwinInstance{}
}

func (*twinInstanceRepository) InsertTwinInstance(twinInstance usecase.TwinInstance) error {
	return nil
}

func (*twinInstanceRepository) DeleteTwinInstance(id string) error {
	return nil
}
