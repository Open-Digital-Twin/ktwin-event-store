package domain

import "time"

type TwinInstance struct {
	Id            string    `json:"id" validate:"required"`
	Name          string    `json:"name" validate:"required"`
	InterfaceId   string    `db:"interfaceId" validate:"required"`
	Active        bool      `json:"active"`
	Parent        string    `json:"parent"`
	LastEventData string    `json:"lastEventData"`
	CreatedAt     time.Time `json:"createdAt"`
}
