package domain

import "time"

type TwinEvent struct {
	Id          string    `json:"id"`
	Time        time.Time `json:"time"`
	Source      string    `json:"source"`
	Type        string    `json:"type"`
	InstanceId  string    `json:"instanceId"`
	InterfaceId string    `json:"interfaceId"`
	EventData   []byte    `json:"eventData"`
	CreatedAt   time.Time `json:"createdAt"`
}
