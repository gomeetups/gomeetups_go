package models

import (
	"time"

	pq "github.com/lib/pq"
)

// Event Defines a single event
type Event struct {
	EventID              string         `json:"eventId" db:"event_id"` // Unique event id
	GroupID              string         `json:"groupId" db:"group_id"` // To which group this event belongs
	SpaceID              string         `json:"spaceID" db:"space_id"` // Where is this event happening
	EventTime            time.Time      `json:"eventTime" db:"event_time"`
	Name                 string         `json:"name" db:"name"`
	Description          string         `json:"description,omitempty" db:"description"`
	Speakers             pq.StringArray `json:"owner,omitempty" db:"speakers"`
	Admins               pq.StringArray `json:"admins,omitempty" db:"admins"`
	Status               string         `json:"status" db:"status"`                                //	draft, active, done, cancelled, deleted
	AttendeeLimit        int64          `json:"attendeeLimit" db:"attendee_limit"`                 // How many people can attend to this event
	AdditionalGuestLimit int64          `json:"additionalGuestLimmit" db:"addiitonal_guest_limit"` // how many people a guest can bring with her
	IsPublic             bool           `json:"isPublic" db:"is_public"`
	IsArchived           bool           `json:"isArchived" db:"is_archived"`
	CreatedAt            time.Time      `json:"createdAt" db:"created_at"`
	CreatedBy            string         `json:"createdBy" db:"created_by"`
	UpdatedAt            time.Time      `json:"updatedAt" db:"updated_at"`
	UpdatedBy            string         `json:"updatedBy" db:"updated_by"`

	Space  *Space   `json:"space,omitempty" db:"-"`
	Photos []*Photo `json:"photos,omitempty" db:"-"`
}

// ValidEventSearchParams Valid search parameters for the events endpoint
type ValidEventSearchParams struct {
	Name        string
	Description string
	GroupID     string
}

// EventService Defines what methods are available for events
type EventService interface {
	SearchEvents(filters *ValidEventSearchParams) (events map[string]*Event, err error)
	Get(eventID string) (event Event, err error)
}
