package repository

import (
	"agwermann/dt-service/internal/app/context/twininterface/domain"
	"time"
)

type TwinInterfaceMapper interface {
	ToDomainList(twinInterfaces []TwinInterface) []domain.TwinInterface
	ToDomain(twinInterfaces TwinInterface) domain.TwinInterface
}

type TwinInterface struct {
	Id             string    `db:"id"`
	Name           string    `db:"name"`
	Active         bool      `db:"active"`
	Parent         string    `db:"parent"`
	DataDefinition string    `db:"data_definition"`
	CreatedAt      time.Time `db:"created_at"`
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
