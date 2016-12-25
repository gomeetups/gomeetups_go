package group

import (
	"bitbucket.org/devmach/gomeetups/models"
	"github.com/gin-gonic/gin"
)

func handleSearch(
	groupService models.GroupService,
	addressService models.AddressService,
	photoService models.PhotoService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var params = models.ValidGroupSearchParams{
			Name:        c.DefaultQuery("name", ""),
			Description: c.DefaultQuery("description", ""),
		}

		if groups, err := groupService.SearchGroups(&params); err == nil {
			var idx = 0
			var groupIds = make([]string, len(groups))

			for id := range groups {
				groupIds[idx] = id
				idx++
			}

			addresses, _ := addressService.GetByGroupID(groupIds)
			photos, _ := photoService.GetByGroupID(groupIds)

			for id := range groups {
				if _, ok := addresses[id]; ok {
					groups[id].Address = addresses[id]
				}

				if _, ok := photos[id]; ok {
					groups[id].Photos = photos[id]
				}

			}

			c.JSON(200, gin.H{
				"entries": groups,
			})

		} else {
			c.JSON(200, gin.H{"error": err})
		}

	}
}

func handleGroupDetails(
	groupService models.GroupService,
	addressService models.AddressService) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Params.ByName("groupID")
		if group, err := groupService.Get(groupID); err == nil {

			addresses, _ := addressService.GetByGroupID([]string{group.GroupID})

			if _, ok := addresses[group.GroupID]; ok {
				group.Address = addresses[group.GroupID]
			}

			c.JSON(200, gin.H{"group": group})
		} else {
			c.JSON(404, gin.H{"error": err.Error()})
		}
	}
}

// Router Contains routes for Group endpoints
func Router(router *gin.RouterGroup, groupService models.GroupService, addressService models.AddressService, photoService models.PhotoService) {
	router.GET("/", handleSearch(groupService, addressService, photoService))
	router.GET("/:groupID", handleGroupDetails(groupService, addressService))
}
