package models

import (
	"time"

	pq "github.com/lib/pq"
)

type Group struct {
	GroupId         string         `json:"groupId" db:"group_id"`
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
}

type GroupSearchValidParams struct {
	Name        string
	Description string
}

type GroupService interface {
	SearchGroups(filters *GroupSearchValidParams) (groups map[string]Group, err error)
	GroupDetails(groupId string) (group Group, err error)
}
