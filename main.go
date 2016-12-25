package main

import (
	"net/http"
	"time"

	"bitbucket.org/devmach/gomeetups/address"
	"bitbucket.org/devmach/gomeetups/group"
	"github.com/gin-gonic/gin"
)

var newYork, _ = time.LoadLocation("America/New_York")

func main() {
	router := gin.Default()

	group.Router(router.Group(
		"/api/v1/groups"),
		&group.ServiceMemory{},
		&address.ServiceMemory{})

	http.ListenAndServe(":3000", router)
}
