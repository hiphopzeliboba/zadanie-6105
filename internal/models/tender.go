package models

import (
	"github.com/google/uuid"
	"time"
)

type Tender struct {
	Id        uuid.UUID  `json:"id"`
	Info      TenderInfo `json:"info"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Status    string     `json:"status"`
	Version   int        `json:"version"`
}

type TenderInfo struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ServiceType     string    `json:"serviceType"`
	OrganizationId  uuid.UUID `json:"organizationId"`
	CreatorUsername string    `json:"creatorUsername"`
}
