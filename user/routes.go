package user

import (
	"bitbucket.org/devmach/gomeetups/models"
	"github.com/gin-gonic/gin"
)

func handleSearch(services *models.Services) gin.HandlerFunc {
	return func(c *gin.Context) {

		var params = models.ValidUserSearchParams{
			DisplayName: c.DefaultQuery("displayName", ""),
		}

		if users, err := services.UserService.Search(&params); err == nil {
			scrubbedData := make(map[string]models.User)

			for idx, record := range users {
				user := *record
				user.Email = "***"

				scrubbedData[idx] = user
			}

			c.JSON(200, gin.H{
				"entries": scrubbedData,
			})

		} else {
			c.JSON(404, gin.H{"error": err.Error()})
		}

	}
}

func handleUserDetails(services *models.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Params.ByName("userID")
		if record, err := services.UserService.GetByID(userID); err == nil {
			user := *record
			user.Email = "***"

			c.JSON(200, gin.H{"user": user})
		} else {
			statusCode := 500

			if err == models.ErrUserNotFound {
				statusCode = 404
			}

			c.JSON(statusCode, gin.H{"error": err.Error()})
		}
	}
}

// Router Contains routes for Group endpoints
func Router(router *gin.RouterGroup, services *models.Services) {
	router.GET("/", handleSearch(services))
	router.GET("/:userID", handleUserDetails(services))
}
