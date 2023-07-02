package domain

import "time"

type TwinInstance struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	InterfaceId   string    `db:"interfaceId"`
	Active        bool      `json:"active"`
	Parent        string    `json:"parent"`
	LastEventData string    `json:"lastEventData"`
	CreatedAt     time.Time `json:"createdAt"`
}
