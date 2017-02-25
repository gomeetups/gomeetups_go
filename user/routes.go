package user

import (
	"net/http"

	"github.com/gomeetups/gomeetups/models"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func handleSearch(services *models.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params = models.ValidUserSearchParams{
			DisplayName: r.URL.Query().Get("displayName"),
		}

		if users, err := services.UserService.Search(&params); err == nil {
			scrubbedData := make(map[string]models.User)

			for idx, record := range users {
				user := *record
				user.Email = "***"

				scrubbedData[idx] = user
			}

			render.JSON(w, r, map[string]interface{}{
				"entries": scrubbedData,
			})

		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]interface{}{
				"error": err.Error(),
			})
		}

	}
}

func handleUserDetails(services *models.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		if record, err := services.UserService.GetByID(userID); err == nil {
			user := *record
			user.Email = "***"

			render.JSON(w, r, map[string]interface{}{
				"user": user,
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

// Router Contains routes for user endpoints
func Router(services *models.Services) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handleSearch(services))
	router.Get("/:userID", handleUserDetails(services))

	return router
}
