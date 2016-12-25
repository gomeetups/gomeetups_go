package fixtures

// Useful SQL queries:
// ---
// Generate 10 UUID v4
// 	SELECT uuid_generate_v4() FROM generate_series(1,10);
//
// Get timezones
// 	Select * from pg_timezone_names()

import (
	"database/sql"
	"time"

	models "bitbucket.org/devmach/gomeetups/models"
)

// Addresses - Contains address fixtures for groups, events and users.... va
var Addresses = []models.Address{
	{
		AddressID:        "e7b33956-c64c-4643-830d-e681663528e5",
		AddressType:      "group",
		BelongsTo:        "6db72c07-1fdd-480e-b9af-7dd96efa4986", // group: GoLang NYC - Manhattan
		Lat:              40.754336,
		Long:             -73.968502,
		AddressLine1:     models.NullString{sql.NullString{String: "928 2nd Ave", Valid: true}},
		FormattedAddress: models.NullString{sql.NullString{String: "928 2nd Ave\nNew York, NY 10022\n", Valid: true}},
		Country:          models.NullString{sql.NullString{String: "US", Valid: true}},
		State:            models.NullString{sql.NullString{String: "NY", Valid: true}},
		City:             models.NullString{sql.NullString{String: "New York", Valid: true}},
		ZipCode:          models.NullString{sql.NullString{String: "10022", Valid: true}},
		TimeZone:         models.NullString{sql.NullString{String: "America/New_York", Valid: true}},
		MapsLink:         models.NullString{sql.NullString{String: "https://www.google.com/maps/place/Sip+Sak/@40.754364,-73.9707237,17z/data=!3m1!4b1!4m5!3m4!1s0x89c258e2f309a3e1:0x349646006d6ae1a2!8m2!3d40.754364!4d-73.968535", Valid: true}},
		CreatedAt:        time.Date(2016, time.December, 20, 7, 10, 0, 0, newYork),
		CreatedBy:        "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:        time.Date(2016, time.December, 20, 7, 10, 0, 0, newYork),
		UpdatedBy:        "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
	},
	{
		AddressID:        "15baeb85-ecee-4363-86c8-0133cea23809",
		AddressType:      "group",
		BelongsTo:        "1d7bffd6-80ab-48f1-b35f-96378f0e78a8", // group: GoLang NYC - Manhattan
		Lat:              40.74387,
		Long:             -73.9221448,
		AddressLine1:     models.NullString{sql.NullString{String: "42-03 Queens Blvd", Valid: true}},
		FormattedAddress: models.NullString{sql.NullString{String: "42-03 Queens Blvd\nSunnyside, NY 11104\n", Valid: true}},
		Country:          models.NullString{sql.NullString{String: "US", Valid: true}},
		State:            models.NullString{sql.NullString{String: "NY", Valid: true}},
		City:             models.NullString{sql.NullString{String: "New York", Valid: true}},
		ZipCode:          models.NullString{sql.NullString{String: "11104", Valid: true}},
		TimeZone:         models.NullString{sql.NullString{String: "America/New_York", Valid: true}},
		MapsLink:         models.NullString{sql.NullString{String: "https://www.google.com/maps/place/42-03+Queens+Blvd,+Sunnyside,+NY+11104/@40.743874,-73.9243335,17z/data=!3m1!4b1!4m5!3m4!1s0x89c25ed9b2e3016b:0xd8eac585eecdeac5!8m2!3d40.74387!4d-73.9221448", Valid: true}},
		CreatedAt:        time.Date(2016, time.December, 20, 8, 30, 0, 0, newYork),
		CreatedBy:        "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		UpdatedAt:        time.Date(2016, time.December, 20, 8, 30, 0, 0, newYork),
		UpdatedBy:        "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
	},
}
