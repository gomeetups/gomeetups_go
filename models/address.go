package models

import (
	"database/sql"
	"time"
)

type Address struct {
	AddressId        string         `json:"addressId" db:"address_id"`
	AddressType      string         `json:"addressType" db:"address_type"` // group, event, user
	BelongsTo        string         `json:"belongsTo" db:"belongs_to"`
	Lat              float64        `json:"lat" db:"lat"`
	Long             float64        `json:"long" db:"long"`
	AddressLine1     sql.NullString `json:"addressLine1,omitempty" db:"address_line_1"`
	AddressLine2     sql.NullString `json:"addressLine2,omitempty" db:"address_line_2"`
	FormattedAddress sql.NullString `json:"formattedAddress,omitempty" db:"formatted_address"`
	Country          sql.NullString `json:"country,omitempty" db:"country"`
	State            sql.NullString `json:"state,omitempty" db:"state"`
	City             sql.NullString `json:"city,omitempty" db:"city"`
	ZipCode          sql.NullString `json:"zipCode,omitempty" db:"zip_code"`
	TimeZone         sql.NullString `json:"timeZone,omitempty" db:"timezone"` // See: Select * from pg_timezone_names()
	MapsLink         sql.NullString `json:"mapsLink,omitempty" db:"maps_link"`
	CreatedAt        time.Time      `json:"createdAt" db:"created_at"`
	CreatedBy        string         `json:"createdBy" db:"created_by"`
	UpdatedAt        time.Time      `json:"updatedAt" db:"updated_at"`
	UpdatedBy        string         `json:"updatedBy" db:"updated_by"`
}

type AddressService interface {
	Get(struct {
		addressId string
		eventId   string
		groupId   string
	}) (address Address, err error)
}
