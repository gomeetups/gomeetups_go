package main

import (
	"net/http"
	"time"

	"bitbucket.org/devmach/gomeetups/address"
	"bitbucket.org/devmach/gomeetups/event"
	"bitbucket.org/devmach/gomeetups/group"
	"bitbucket.org/devmach/gomeetups/models"
	"bitbucket.org/devmach/gomeetups/photo"
	"bitbucket.org/devmach/gomeetups/space"
	"github.com/gin-gonic/gin"
)

var newYork, _ = time.LoadLocation("America/New_York")

func main() {
	router := gin.Default()

	services := models.Services{
		GroupService:   &group.ServiceMemory{},
		AddressService: &address.ServiceMemory{},
		PhotoService:   &photo.ServiceMemory{},
		SpaceService:   &space.ServiceMemory{},
		EventService:   &event.ServiceMemory{},
	}

	group.Router(router.Group("/api/v1/groups"), &services)
	event.Router(router.Group("/api/v1/events"), &services)

	http.ListenAndServe(":3000", router)
}
