package models

import (
	"time"

	pq "github.com/lib/pq"
)

// Group Core model definition
type Group struct {
	GroupID         string         `json:"groupId" db:"group_id"`
	Name            string         `json:"name" db:"name"`
	Slug            string         `json:"slug" db:"slug"`
	Description     string         `json:"description,omitempty" db:"description"`
	Owner           string         `json:"owner" db:"owner"`
	Admins          pq.StringArray `json:"admins" db:"admins"`
	JoinRestriction string         `json:"joinRestriction" db:"join_restriction"` // needs_approval, invite, open
	IsPublic        bool           `json:"isPublic" db:"is_public"`
	IsArchived      bool           `json:"isArchived" db:"is_archived"`
	CreatedAt       time.Time      `json:"createdAt" db:"created_at"`
	CreatedBy       string         `json:"createdBy" db:"created_by"`
	UpdatedAt       time.Time      `json:"updatedAt" db:"updated_at"`
	UpdatedBy       string         `json:"updatedBy" db:"updated_by"`

	// Without a `pointer`, go will always create a new Group.Address
	// Whiloi initializing object, `omitempty` will be ignored. This
	// will cause junk data in the json output.
	Address *Address `json:"address,omitempty" db:"-"`
	Photos  []*Photo `json:"photos,omitempty" db:"-"`
}

// ValidGroupSearchParams Valid search parameters for the groups endpoint
type ValidGroupSearchParams struct {
	Name        string
	Description string
}

// GroupService - Contains necessary methods to be implemented
// by address group service
type GroupService interface {
	SearchGroups(filters *ValidGroupSearchParams) (groups map[string]*Group, err error)
	Get(groupID string) (group Group, err error)
}
