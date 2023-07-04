package domain

import "time"

type TwinEvent struct {
	InstanceId  string    `json:"instanceId"`
	InterfaceId string    `json:"interfaceId"`
	EventData   string    `json:"eventData"`
	CreatedAt   time.Time `json:"createdAt"`
}
