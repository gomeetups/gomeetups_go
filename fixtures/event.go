package fixtures

import (
	"time"

	"github.com/gomeetups/gomeetups/models"
)

// "62036543-292e-4751-ace0-13624d09c79e"
// "309a4f28-249f-4471-8a97-60e33e980f34"
// "972d9fe9-8e41-4fda-a8e9-ec7bb8d17c25"

// "6db72c07-1fdd-480e-b9af-7dd96efa4986", // group: GoLang NYC - Manhattan
// "1d7bffd6-80ab-48f1-b35f-96378f0e78a8", // group: GoLang NYC - Queens
// "75b7410a-f42d-4305-8372-439afb24f83f", // group: GoLang NYC - Brooklyn

// "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
// "0c8e35a5-4c1c-4294-8927-d23f0b1b4969", // user: kaylyn
// "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc

// Events Contains event fixtures for in memory service
var Events = []models.Event{
	{
		EventID:   "43e3aef7-af84-468c-a65a-0fca47ca3c9d",
		GroupID:   "6db72c07-1fdd-480e-b9af-7dd96efa4986", // group: GoLang NYC - Manhattan
		SpaceID:   "863137fc-3cee-4e8c-ae49-0375fcbe2707", // space: Moma NYC
		EventTime: time.Date(2016, time.December, 31, 20, 30, 0, 0, newYork),
		Name:      "Manhattan Gophers New Year Party!",
		Description: `
      Hello Manhattan Gophers!,
      Let's meet and celebrate new year together!
    `,
		Speakers: []string{
			"d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		},
		Admins: []string{
			"d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		},
		Status:               "active",
		AttendeeLimit:        100,
		AdditionalGuestLimit: 2,
		IsPublic:             true,
		IsArchived:           false,
		CreatedAt:            time.Now(),
		CreatedBy:            "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:            time.Now(),
		UpdatedBy:            "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydins
	},
	{
		EventID:   "33402f90-d64b-44ed-8f26-dceb258e2989",
		GroupID:   "6db72c07-1fdd-480e-b9af-7dd96efa4986", // group: GoLang NYC - Manhattan
		SpaceID:   "863137fc-3cee-4e8c-ae49-0375fcbe2707", // space: Moma NYC
		EventTime: time.Date(2016, time.January, 15, 18, 30, 0, 0, newYork),
		Name:      "Monthly Manhattan Gophers Meet-up!",
		Description: `
      Hello Manhattan Gophers!,
      Let's get togeter at 2017's first monthly meetup
    `,
		Speakers: []string{
			"d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		},
		Admins: []string{
			"d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		},
		Status:               "active",
		AttendeeLimit:        20,
		AdditionalGuestLimit: 5,
		IsPublic:             true,
		IsArchived:           false,
		CreatedAt:            time.Now(),
		CreatedBy:            "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:            time.Now(),
		UpdatedBy:            "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydins
	},
	{
		EventID:   "3a5d5357-437a-4b11-b63f-180ebdb27f35",
		GroupID:   "1d7bffd6-80ab-48f1-b35f-96378f0e78a8", // group: GoLang NYC - Queens
		SpaceID:   "77ab205b-494e-4012-a1f8-96a38e1c3e52", // space: Bohemian Hall & Beer Garden
		EventTime: time.Date(2016, time.January, 20, 18, 30, 0, 0, newYork),
		Name:      "Monthly Queens Gophers Meet-up!",
		Description: `
      Hello Queens Gophers!,
      Let's get togeter at Jan 20th for 2017's first monthly meetup
    `,
		Speakers: []string{
			"478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		},
		Admins: []string{
			"478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		},
		Status:               "active",
		AttendeeLimit:        20,
		AdditionalGuestLimit: 5,
		IsPublic:             true,
		IsArchived:           false,
		CreatedAt:            time.Now(),
		CreatedBy:            "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		UpdatedAt:            time.Now(),
		UpdatedBy:            "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
	},
}
