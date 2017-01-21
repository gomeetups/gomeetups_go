package models

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

var (
	ErrRsvpClosed              = errors.New("RSVPs closed")
	ErrRsvpMaxAdditionalGuests = errors.New("Additional guest counts is more then allowed")

	RsvpResponse    = map[string]string{"yes": "yes", "no": "no"}
	RsvpCancelledBy = map[string]string{"admin": "admin", "guest": "guest"}
	RsvpShowNoShow  = map[string]string{"show": "show", "no show": "no show"}
)

// Rsvp represents RSVP model...
type Rsvp struct {
	RsvpID       string         `json:"rsvpId" db:"rsvp_id"`
	GroupID      string         `json:"groupId" db:"group_id"`
	EventID      string         `json:"eventId" db:"event_id"`
	UserID       string         `json:"userId" db:"user_id"`
	Response     string         `json:"response" db:"response"`        // See RsvpResponse
	GuestCount   int            `json:"guestCount" db:"guest_count"`   // min:1 if user goes by herself
	GuestNames   pq.StringArray `json:"guestNames" db:"guest_names"`   // ['Display name of user']
	ShowNoShow   bool           `json:"showNoShow" db:"show_noshow"`   // So event admin will keep track of no-shows
	IsCancelled  bool           `json:"isCancelled" db:"is_cancelled"` // 'show': user attended the event, 'no show': user responded yes but didn't showed up
	CancelledBy  string         `json:"cancelledBy" db:"cancelled_by"` // admin - guest
	IsInWaitList bool           `json:"isInWaitList" db:"is_in_wait_list"`
	CreatedAt    time.Time      `json:"createdAt" db:"created_at"`
	CreatedBy    string         `json:"createdBy" db:"created_by"` // User herself
	UpdatedAt    time.Time      `json:"updatedAt" db:"updated_at"`
	UpdatedBy    string         `json:"updatedBy" db:"updated_by"`
}

// RsvpService Handles Rsvp related operations
type RsvpService interface {
	// GetAll returns all Rsvps for given event
	GetAll(EventID string) (rsvps []Rsvp, err error)
}
