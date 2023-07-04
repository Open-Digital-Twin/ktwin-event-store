package repository

import "agwermann/dt-service/internal/app/context/twinevent/domain"

type TwinEventMapper interface {
	ToDomainList(twinEvents []TwinEvent) []domain.TwinEvent
	ToDomain(twinEvent TwinEvent) domain.TwinEvent
}

type twinEventMapper struct {
}

func NewTwinEventMapper() TwinEventMapper {
	return &twinEventMapper{}
}

func (*twinEventMapper) ToDomainList(twinEvents []TwinEvent) []domain.TwinEvent {
	var twinEventsDomain []domain.TwinEvent

	for _, twinEvent := range twinEvents {
		twinEventsDomain = append(twinEventsDomain, domain.TwinEvent(twinEvent))
	}

	return twinEventsDomain
}

func (*twinEventMapper) ToDomain(twinEvent TwinEvent) domain.TwinEvent {
	return domain.TwinEvent(twinEvent)
}
