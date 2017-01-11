package event

import (
	"github.com/gomeetups/gomeetups/models"
	"github.com/gin-gonic/gin"
)

func handleSearch(services *models.Services) gin.HandlerFunc {
	return func(c *gin.Context) {

		var params = models.ValidEventSearchParams{
			Name:        c.DefaultQuery("name", ""),
			Description: c.DefaultQuery("description", ""),
			GroupID:     c.DefaultQuery("group", ""),
		}

		if events, err := services.EventService.SearchEvents(&params); err == nil {
			var idx = 0
			var eventIds = make([]string, len(events))
			var spaceIds = make([]string, len(events))

			for id, event := range events {
				eventIds[idx] = id
				spaceIds[idx] = event.SpaceID
				idx++
			}

			spaces, _ := services.SpaceService.GetMultiple(spaceIds)

			for idx, event := range events {
				if _, ok := spaces[event.SpaceID]; ok {
					events[idx].Space = spaces[event.SpaceID]
				}
			}

			c.JSON(200, gin.H{
				"entries": events,
			})

		} else {
			c.JSON(200, gin.H{"error": err})
		}

	}
}

func handleEventDetails(services *models.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID := c.Params.ByName("eventID")
		if event, err := services.EventService.Get(eventID); err == nil {

			if space, _ := services.SpaceService.Get(event.SpaceID); space != nil {
				event.Space = space
			}

			c.JSON(200, gin.H{"event": event})
		} else {
			c.JSON(404, gin.H{"error": err.Error()})
		}
	}
}

// Router Contains routes for Group endpoints
func Router(router *gin.RouterGroup, services *models.Services) {
	router.GET("/", handleSearch(services))
	router.GET("/:eventID", handleEventDetails(services))
}
