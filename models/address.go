package models

import "time"

var AddressTypes = map[string]string{
	"GROUP": "group",
}

// Address Defines address model user by groups, events and users
type Address struct {
	AddressID        string     `json:"addressId,omitempty" db:"address_id"`
	AddressType      string     `json:"addressType,omitempty" db:"address_type"` // group, event, user
	BelongsTo        string     `json:"belongsTo,omitempty" db:"belongs_to"`
	Lat              float64    `json:"lat,omitempty" db:"lat"`
	Long             float64    `json:"long,omitempty" db:"long"`
	AddressLine1     NullString `json:"addressLine1,omitempty" db:"address_line_1"`
	AddressLine2     NullString `json:"addressLine2,omitempty" db:"address_line_2"`
	FormattedAddress NullString `json:"formattedAddress,omitempty" db:"formatted_address"`
	Country          NullString `json:"country,omitempty" db:"country"`
	State            NullString `json:"state,omitempty" db:"state"`
	City             NullString `json:"city,omitempty" db:"city"`
	ZipCode          NullString `json:"zipCode,omitempty" db:"zip_code"`
	TimeZone         NullString `json:"timeZone,omitempty" db:"timezone"` // See: Select * from pg_timezone_names()
	MapsLink         NullString `json:"mapsLink,omitempty" db:"maps_link"`
	CreatedAt        time.Time  `json:"createdAt,omitempty" db:"created_at"`
	CreatedBy        string     `json:"createdBy,omitempty" db:"created_by"`
	UpdatedAt        time.Time  `json:"updatedAt,omitempty" db:"updated_at"`
	UpdatedBy        string     `json:"updatedBy,omitempty" db:"updated_by"`
}

// AddressService - Contains necessary methods to be implemented
// by address service
type AddressService interface {
	GetByGroupId(groupIds []string) (addresses map[string]*Address, err error)
	Get(addressID string) (address Address, err error)
}
