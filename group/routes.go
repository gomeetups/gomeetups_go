package group

import "github.com/gin-gonic/gin"
import "bitbucket.org/devmach/gomeetups/models"

func handleSearch(service models.GroupService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var params = models.GroupSearchValidParams{
			Name:        c.DefaultQuery("name", ""),
			Description: c.DefaultQuery("description", ""),
		}

		if services, err := service.SearchGroups(&params); err == nil {
			c.JSON(200, gin.H{"entries": services})
		} else {
			c.JSON(200, gin.H{"error": err})
		}

	}
}

func handleGroupDetails(service models.GroupService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Params.ByName("name")
		if group, err := service.GroupDetails(name); err == nil {
			c.JSON(200, gin.H{"group": group})
		} else {
			c.JSON(404, gin.H{"error": err.Error()})
		}
	}
}

func Router(router *gin.RouterGroup, service models.GroupService) {
	router.GET("/", handleSearch(service))
	router.GET("/:name", handleGroupDetails(service))
}
