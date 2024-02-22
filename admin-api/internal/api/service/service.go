package service

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	Id                 uuid.UUID `json:"id"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	Name               string    `json:"name"`
	Host               string    `json:"host"`
	Port               uint      `json:"port"`
	Protocol           string    `json:"protocol"`
	Tags               []string  `json:"tags"`
	Enabled            bool      `json:"enabled"`
	HealthCheckEnabled bool      `json:"healthCheckEnabled"`
	Health             string    `json:"health"`
}
