package main

import (
	"net/http"
	"time"

	"bitbucket.org/devmach/gomeetups/address"
	"bitbucket.org/devmach/gomeetups/group"
	"bitbucket.org/devmach/gomeetups/photo"
	"github.com/gin-gonic/gin"
)

var newYork, _ = time.LoadLocation("America/New_York")

func main() {
	router := gin.Default()

	group.Router(router.Group(
		"/api/v1/groups"),
		&group.ServiceMemory{},
		&address.ServiceMemory{},
		&photo.ServiceMemory{})

	http.ListenAndServe(":3000", router)
}
