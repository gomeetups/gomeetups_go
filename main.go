package main

import (
	"net/http"
	"time"

	"bitbucket.org/devmach/gomeetups/address"
	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/group"
	"bitbucket.org/devmach/gomeetups/models"
	"github.com/gin-gonic/gin"
)

var newYork, _ = time.LoadLocation("America/New_York")

func main() {
	router := gin.Default()

	for i := 0; i < 1000000; i++ {
		fixtures.Groups[string(i)] = models.Group{
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
			CreatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: asydin
			UpdatedAt:       time.Now(),
			UpdatedBy:       "d468bd91-39a2-46a1-99c3-4c4b0f20e78a", // user: aydin
		}

	}

	group.Router(router.Group(
		"/api/v1/groups"),
		&group.ServiceMemory{},
		&address.ServiceMemory{})

	http.ListenAndServe(":3000", router)
}
