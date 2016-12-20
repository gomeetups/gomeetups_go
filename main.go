package main

import (
	"net/http"

	"bitbucket.org/devmach/gomeetups/group"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	group.Router(router.Group("/api/v1/groups"), &group.GroupServiceMem{Limit: 100})

	http.ListenAndServe(":3000", router)
}
