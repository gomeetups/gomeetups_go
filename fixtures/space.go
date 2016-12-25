package fixtures

import (
	"time"

	"bitbucket.org/devmach/gomeetups/models"
)

// Spaces Contains space fixtures for memory service
var Spaces = []models.Space{
	{
		SpaceID:     "863137fc-3cee-4e8c-ae49-0375fcbe2707",
		GroupID:     "6db72c07-1fdd-480e-b9af-7dd96efa4986", // group: GoLang NYC - Manhattan
		Name:        "The Museum of Modern Art",
		Description: "The Museum of Modern Art (MoMA) is a place that fuels creativity, ignites minds, and provides inspiration",
		CreatedAt:   time.Date(2016, time.December, 25, 12, 25, 0, 0, newYork),
		CreatedBy:   "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:   time.Date(2016, time.December, 25, 12, 25, 0, 0, newYork),
		UpdatedBy:   "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin

		Address: &Addresses[2],
	},
}
