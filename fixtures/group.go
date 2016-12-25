package fixtures

import (
	"time"

	"bitbucket.org/devmach/gomeetups/models"
)

// Groups Contains fixtures for groups
var Groups = []models.Group{
	{
		GroupID:     "6db72c07-1fdd-480e-b9af-7dd96efa4986",
		Name:        "GoLang NYC - Manhattan",
		Slug:        "golang-nyc-manhattan",
		Description: "Manhattan go users group",
		Owner:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		Admins: []string{
			"d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
			"0c8e35a5-4c1c-4294-8927-d23f0b1b4969", // user: kaylyn
			"478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		},
		JoinRestriction: "open",
		IsPublic:        true,
		IsArchived:      false,
		CreatedAt:       time.Now(),
		CreatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:       time.Now(),
		UpdatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
	},
	{
		GroupID:     "1d7bffd6-80ab-48f1-b35f-96378f0e78a8",
		Name:        "GoLang NYC - Queens",
		Slug:        "golang-nyc-queens",
		Description: "Queens go users group",
		Owner:       "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		Admins: []string{
			"478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		},
		JoinRestriction: "open",
		IsPublic:        true,
		IsArchived:      false,
		CreatedAt:       time.Now(),
		CreatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:       time.Now(),
		UpdatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
	},
	{
		GroupID:     "75b7410a-f42d-4305-8372-439afb24f83f",
		Name:        "GoLang NYC - Brooklyn",
		Slug:        "golang-nyc-brooklyn",
		Description: "Brooklyn go users group",
		Owner:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		Admins: []string{
			"d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		},
		JoinRestriction: "open",
		IsPublic:        true,
		IsArchived:      false,
		CreatedAt:       time.Now(),
		CreatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:       time.Now(),
		UpdatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
	},
}
