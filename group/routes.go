package group

import (
	"net/http"

	"github.com/gomeetups/gomeetups/models"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func handleSearch(services *models.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params = models.ValidGroupSearchParams{
			Name:        r.URL.Query().Get("name"),
			Description: r.URL.Query().Get("description"),
		}

		if groups, err := services.GroupService.SearchGroups(&params); err == nil {
			var idx = 0
			var groupIds = make([]string, len(groups))

			for id := range groups {
				groupIds[idx] = id
				idx++
			}

			addresses, _ := services.AddressService.GetByGroupID(groupIds)
			photos, _ := services.PhotoService.GetByGroupID(groupIds)

			for id := range groups {
				if _, ok := addresses[id]; ok {
					groups[id].Address = addresses[id]
				}

				if _, ok := photos[id]; ok {
					groups[id].Photos = photos[id]
				}

			}

			render.JSON(w, r, map[string]interface{}{
				"entries": groups,
			})

		} else {
			render.JSON(w, r, map[string]interface{}{
				"error": err,
			})
		}

	}
}

func handleGroupDetails(services *models.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupID := chi.URLParam(r, "groupID")
		if group, err := services.GroupService.Get(groupID); err == nil {

			addresses, _ := services.AddressService.GetByGroupID([]string{group.GroupID})
			photos, _ := services.PhotoService.GetByGroupID([]string{group.GroupID})

			if _, ok := addresses[group.GroupID]; ok {
				group.Address = addresses[group.GroupID]
			}

			if _, ok := photos[group.GroupID]; ok {
				group.Photos = photos[group.GroupID]
			}

			render.JSON(w, r, map[string]interface{}{
				"group": group,
			})

		} else {
			render.JSON(w, r, map[string]interface{}{
				"error": err.Error(),
			})
		}
	}
}

// Router Contains routes for Group endpoints
func Router(services *models.Services) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handleSearch(services))
	router.Get("/:groupID", handleGroupDetails(services))

	return router
}
