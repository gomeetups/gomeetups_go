package fixtures

import (
	"time"

	models "github.com/gomeetups/gomeetups/models"
)

// Rsvps Contains fictures for in memory service
var Rsvps = []models.Rsvp{
	{
		RsvpID:       "62036543-292e-4751-ace0-13624d09c79e",
		GroupID:      "6db72c07-1fdd-480e-b9af-7dd96efa4986",
		EventID:      "43e3aef7-af84-468c-a65a-0fca47ca3c9d",
		UserID:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a",
		Response:     models.RsvpResponse["yes"],
		GuestCount:   1,
		GuestNames:   []string{"Aydin U."},
		IsCancelled:  false,
		IsInWaitList: false,
		CreatedAt:    time.Now(),
		CreatedBy:    "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:    time.Now(),
		UpdatedBy:    "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydins
	},
	{
		RsvpID:       "309a4f28-249f-4471-8a97-60e33e980f34",
		GroupID:      "6db72c07-1fdd-480e-b9af-7dd96efa4986",
		EventID:      "43e3aef7-af84-468c-a65a-0fca47ca3c9d",
		UserID:       "478f4b8e-b231-4efe-828b-f63b877fbbe3",
		Response:     models.RsvpResponse["yes"],
		GuestCount:   2,
		GuestNames:   []string{"Arc K.", "Ken F."},
		IsCancelled:  false,
		IsInWaitList: false,
		CreatedAt:    time.Now(),
		CreatedBy:    "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		UpdatedAt:    time.Now(),
		UpdatedBy:    "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
	},
}
