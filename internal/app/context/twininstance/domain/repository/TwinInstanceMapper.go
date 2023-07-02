package repository

import "agwermann/dt-service/internal/app/context/twininstance/domain"

type TwinInstanceMapper interface {
	ToDomain(twinInstances []TwinInstance) []domain.TwinInstance
}

type twinInstanceMapper struct {
}

func NewTwinInstanceMapper() TwinInstanceMapper {
	return &twinInstanceMapper{}
}

func (*twinInstanceMapper) ToDomain(twinInstances []TwinInstance) []domain.TwinInstance {
	var twinInstancesDomain []domain.TwinInstance

	for _, twinInstance := range twinInstances {
		twinInstancesDomain = append(twinInstancesDomain, domain.TwinInstance(twinInstance))
	}

	return twinInstancesDomain
}
