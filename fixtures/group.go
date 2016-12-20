package fixtures

import (
	"time"

	models "bitbucket.org/devmach/gomeetups/models"
)

var Groups = map[string]models.Group{
	"6db72c07-1fdd-480e-b9af-7dd96efa4986": {
		GroupId:     "6db72c07-1fdd-480e-b9af-7dd96efa4986",
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
		CreatedBy:       "Aydin",
		UpdatedAt:       time.Now(),
		UpdatedBy:       "Aydin",
	},
	"1d7bffd6-80ab-48f1-b35f-96378f0e78a8": {
		GroupId:     "1d7bffd6-80ab-48f1-b35f-96378f0e78a8",
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
		CreatedBy:       "Aydin",
		UpdatedAt:       time.Now(),
		UpdatedBy:       "Aydin",
	},
}
