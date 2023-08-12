package repository

import (
	"github.com/ktwins/event-store/internal/app/context/twininterface/domain"
)

type TwinInterfaceMapper interface {
	ToDomainList(twinInterfaces []TwinInterface) []domain.TwinInterface
	ToDomain(twinInterfaces TwinInterface) domain.TwinInterface
}

type twinInstanceMapper struct {
}

func NewTwinInterfaceMapper() TwinInterfaceMapper {
	return &twinInstanceMapper{}
}

func (*twinInstanceMapper) ToDomainList(twinInterfaces []TwinInterface) []domain.TwinInterface {
	var twinInterfacesDomain []domain.TwinInterface

	for _, twinInterface := range twinInterfaces {
		twinInterfacesDomain = append(twinInterfacesDomain, domain.TwinInterface(twinInterface))
	}

	return twinInterfacesDomain
}

func (*twinInstanceMapper) ToDomain(twinInstance TwinInterface) domain.TwinInterface {
	return domain.TwinInterface(twinInstance)
}
