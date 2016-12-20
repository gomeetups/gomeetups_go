package models

import (
	"time"

	pq "github.com/lib/pq"
)

type Event struct {
	EventId     string         `json:"eventId" db:"event_id"`
	GroupId     string         `json:"groupId" db:"group_id"`
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description,omitempty" db:"description"`
	Speakers    pq.StringArray `json:"owner,omitempty" db:"speakers"`
	Admins      pq.StringArray `json:"admins,omitempty" db:"admins"`
	Status      string         `json:"status" db:"status"` //	draft, active, done, cancelled, deleted
	IsPublic    bool           `json:"isPublic" db:"is_public"`
	IsArchived  bool           `json:"isArchived" db:"is_archived"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	CreatedBy   string         `json:"createdBy" db:"created_by"`
	UpdatedAt   time.Time      `json:"updatedAt" db:"updated_at"`
	UpdatedBy   string         `json:"updatedBy" db:"updated_by"`
}

type EventService interface {
	SearchGroups() (groups *map[string]Group, err error)
	GroupDetails(groupId string) (group Group, err error)
}
