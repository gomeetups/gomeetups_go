package main

import (
	"net/http"
	"time"

	"github.com/gomeetups/gomeetups/address"
	"github.com/gomeetups/gomeetups/event"
	"github.com/gomeetups/gomeetups/group"
	"github.com/gomeetups/gomeetups/models"
	"github.com/gomeetups/gomeetups/photo"
	"github.com/gomeetups/gomeetups/space"
	"github.com/gomeetups/gomeetups/user"
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
		UserService:    &user.ServiceMemory{},
	}

	group.Router(router.Group("/api/v1/groups"), &services)
	event.Router(router.Group("/api/v1/events"), &services)
	user.Router(router.Group("/api/v1/users"), &services)

	http.ListenAndServe(":3000", router)
}
