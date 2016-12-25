package fixtures

import (
	"time"

	"bitbucket.org/devmach/gomeetups/models"
)

// "1d7bffd6-80ab-48f1-b35f-96378f0e78a8", // group: GoLang NYC - Manhattan
// "1d7bffd6-80ab-48f1-b35f-96378f0e78a8", // group: GoLang NYC - Queens
// "75b7410a-f42d-4305-8372-439afb24f83f", // group: GoLang NYC - Brooklyn

// "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
// "0c8e35a5-4c1c-4294-8927-d23f0b1b4969", // user: kaylyn
// "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc

//Photos Contains fixtures for phtos
var Photos = []models.Photo{
	{
		PhotoID:    "911c22ab-398c-4f60-b8e8-1df630741503",
		PhotoType:  "group",
		BelongsTo:  "6db72c07-1fdd-480e-b9af-7dd96efa4986", // group: GoLang NYC - Manhattan
		Resolution: "thumbnail",
		URL:        "https://unsplash.it/120/120",
		CreatedAt:  time.Now(),
		CreatedBy:  "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:  time.Now(),
		UpdatedBy:  "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
	},
	{
		PhotoID:    "55224161-1160-491b-8c21-94f20edf9a87",
		PhotoType:  "group",
		BelongsTo:  "1d7bffd6-80ab-48f1-b35f-96378f0e78a8", // group: GoLang NYC - Queens
		Resolution: "thumbnail",
		URL:        "https://unsplash.it/g/120/120",
		CreatedAt:  time.Now(),
		CreatedBy:  "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
		UpdatedAt:  time.Now(),
		UpdatedBy:  "478f4b8e-b231-4efe-828b-f63b877fbbe3", // user: arc
	},
	{
		PhotoID:    "848bcae1-1812-4c42-af82-0e132015ab8e",
		PhotoType:  "group",
		BelongsTo:  "75b7410a-f42d-4305-8372-439afb24f83f", // group: GoLang NYC - Brooklyn
		Resolution: "thumbnail",
		URL:        "https://unsplash.it/g/120/120?blur",
		CreatedAt:  time.Now(),
		CreatedBy:  "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		UpdatedAt:  time.Now(),
		UpdatedBy:  "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
	},
}
