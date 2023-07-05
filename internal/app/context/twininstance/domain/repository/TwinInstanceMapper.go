package repository

import "agwermann/event-store-service/internal/app/context/twininstance/domain"

type TwinInstanceMapper interface {
	ToDomainList(twinInstances []TwinInstance) []domain.TwinInstance
	ToDomain(twinInstances TwinInstance) domain.TwinInstance
}

type twinInstanceMapper struct {
}

func NewTwinInstanceMapper() TwinInstanceMapper {
	return &twinInstanceMapper{}
}

func (*twinInstanceMapper) ToDomainList(twinInstances []TwinInstance) []domain.TwinInstance {
	var twinInstancesDomain []domain.TwinInstance

	for _, twinInstance := range twinInstances {
		twinInstancesDomain = append(twinInstancesDomain, domain.TwinInstance(twinInstance))
	}

	return twinInstancesDomain
}

func (*twinInstanceMapper) ToDomain(twinInstance TwinInstance) domain.TwinInstance {
	return domain.TwinInstance(twinInstance)
}
