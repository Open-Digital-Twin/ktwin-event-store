package repository

import (
	"agwermann/dt-service/internal/app/context/twininstance/domain"
)

func NewTwinInstanceRepository() TwinInstanceRepository {
	return &twinInstanceRepository{}
}

type TwinInstanceRepository interface {
	GetAllTwinInstances()
	GetTwinInstance(id string) domain.TwinInstance
	InsertTwinInstance(twinInstance domain.TwinInstance) error
	DeleteTwinInstance(id string) error
}

type twinInstanceRepository struct {
}

func (*twinInstanceRepository) GetAllTwinInstances() {
	return
}

func (*twinInstanceRepository) GetTwinInstance(id string) domain.TwinInstance {
	return domain.TwinInstance{}
}

func (*twinInstanceRepository) InsertTwinInstance(twinInstance domain.TwinInstance) error {
	return nil
}

func (*twinInstanceRepository) DeleteTwinInstance(id string) error {
	return nil
}
