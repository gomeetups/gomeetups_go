package models

import (
	"time"

	pq "github.com/lib/pq"
)

type Photo struct {
	PhotoId   string           `json:"photoid" db:"photo_id"`
	PhotoType string           `json:"photoType" db:"photo_type"` // group, event, user
	BelongsTo string           `json:"belongsTo" db:"belongs_to"`
	Links     []pq.StringArray `json:"links" db:"links"`
	CreatedAt time.Time        `json:"createdAt" db:"created_at"`
	CreatedBy string           `json:"createdBy" db:"created_by"`
	UpdatedAt time.Time        `json:"updatedAt" db:"updated_at"`
	UpdatedBy string           `json:"updatedBy" db:"updated_by"`
}

type PhotoService interface {
	FindPhotos(photoType string, belongsTo string) (photos []Photo, err error)
}
