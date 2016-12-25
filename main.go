package main

import (
	"net/http"
	"time"

	"bitbucket.org/devmach/gomeetups/address"
	"bitbucket.org/devmach/gomeetups/group"
	"bitbucket.org/devmach/gomeetups/models"
	"bitbucket.org/devmach/gomeetups/photo"
	"github.com/gin-gonic/gin"
)

var newYork, _ = time.LoadLocation("America/New_York")

func main() {
	router := gin.Default()

	services := models.Services{
		GroupService:   &group.ServiceMemory{},
		AddressService: &address.ServiceMemory{},
		PhotoService:   &photo.ServiceMemory{},
	}
	group.Router(router.Group(
		"/api/v1/groups"),
		&services)

	http.ListenAndServe(":3000", router)
}
