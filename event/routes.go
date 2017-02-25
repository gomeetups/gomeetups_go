package event

import (
	"net/http"

	"github.com/gomeetups/gomeetups/models"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func handleSearch(services *models.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params = models.ValidEventSearchParams{
			Name:        r.URL.Query().Get("name"),
			Description: r.URL.Query().Get("description"),
			GroupID:     r.URL.Query().Get("group"),
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

			render.JSON(w, r, map[string]interface{}{
				"entries": events,
			})

		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]interface{}{
				"error": err,
			})
		}

	}
}

func handleEventDetails(services *models.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventID := chi.URLParam(r, "eventID")
		if event, err := services.EventService.Get(eventID); err == nil {

			if space, _ := services.SpaceService.Get(event.SpaceID); space != nil {
				event.Space = space
			}

			event.Rsvps, _ = services.RsvpService.GetByEventID(event.EventID)

			render.JSON(w, r, map[string]interface{}{
				"event": event,
			})
		} else {
			status := http.StatusInternalServerError

			if err == models.ErrRecordNotFound {
				status = http.StatusNotFound
			}
			render.Status(r, status)

			render.JSON(w, r, map[string]interface{}{
				"error": err.Error(),
			})
		}
	}
}

// Router Contains routes for event endpoints
func Router(services *models.Services) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handleSearch(services))
	router.Get("/:eventID", handleEventDetails(services))

	return router
}
