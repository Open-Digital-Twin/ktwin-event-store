package domain

import "time"

type TwinInterface struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Active         bool      `json:"active"`
	Parent         string    `json:"parent"`
	DataDefinition string    `json:"dataDefinition"`
	CreatedAt      time.Time `json:"createdAt"`
}
