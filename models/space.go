package models

import "time"

// Space Defines a space where an event can occur....
type Space struct {
	SpaceID     string    `json:"spaceId" db:"space_id"`
	GroupID     string    `json:"groupId" db:"group_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt,omitempty" db:"created_at"`
	CreatedBy   string    `json:"createdBy,omitempty" db:"created_by"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	UpdatedBy   string    `json:"updatedBy,omitempty" db:"updated_by"`

	Address *Address `json:"address,omitempty" db:"-"`
}

// SpaceService Handles space related operations
type SpaceService interface {
	Get(spaceID string) (space *Space, err error)
	GetMultiple(spaceIds []string) (spaces map[string]*Space, err error)
}
