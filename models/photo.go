package models

import "time"

// PhotoTypes ENUM for photo types...
var PhotoTypes = map[string]string{
	"GROUP": "group",
	"EVENT": "event",
	"SPACE": "space",
	"USER":  "user",
}

// Photo Photo model
type Photo struct {
	PhotoID    string    `json:"photoid" db:"photo_id"`
	PhotoType  string    `json:"photoType" db:"photo_type"` // group, event, user
	BelongsTo  string    `json:"belongsTo" db:"belongs_to"`
	Resolution string    `json:"resolution" db:"resolution"` // thumbnail, web, hi-res, 800x600
	URL        string    `json:"url" db:"url"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	CreatedBy  string    `json:"createdBy" db:"created_by"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
	UpdatedBy  string    `json:"updatedBy" db:"updated_by"`
}

// PhotoService Handles photo related operations
type PhotoService interface {
	FindPhotos(photoType string, belongsTo string, resolution string) (photos map[string]*Photo, err error)
	GetByGroupID(groupIds []string) (photos map[string][]*Photo, err error)
}
