package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gomeetups/gomeetups/address"
	"github.com/gomeetups/gomeetups/event"
	"github.com/gomeetups/gomeetups/group"
	"github.com/gomeetups/gomeetups/models"
	"github.com/gomeetups/gomeetups/photo"
	"github.com/gomeetups/gomeetups/rsvp"
	"github.com/gomeetups/gomeetups/space"
	"github.com/gomeetups/gomeetups/user"
	"github.com/pressly/chi"
	"github.com/pressly/lg"
)

var newYork, _ = time.LoadLocation("America/New_York")

func getListenAddr() string {
	PORT := os.Getenv("PORT")
	HOST := os.Getenv("HOST")

	if PORT == "" {
		PORT = "5000"
	}

	return fmt.Sprintf("%s:%s", HOST, PORT)
}

func main() {
	services := models.Services{
		GroupService:   &group.ServiceMemory{},
		AddressService: &address.ServiceMemory{},
		PhotoService:   &photo.ServiceMemory{},
		SpaceService:   &space.ServiceMemory{},
		EventService:   &event.ServiceMemory{},
		UserService:    &user.ServiceMemory{},
		RsvpService:    &rsvp.ServiceMemory{},
	}

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	lg.RedirectStdlogOutput(logger)
	lg.DefaultLogger = logger

	serverCtx := context.Background()
	serverCtx = lg.WithLoggerContext(serverCtx, logger)

	lg.Log(serverCtx).Infof("Listening on %s", getListenAddr())

	router := chi.NewRouter()
	router.Use(lg.RequestLogger(logger))

	router.Route("/api/v1", func(router chi.Router) {
		router.Mount("/groups", group.Router(&services))
		router.Mount("/events", event.Router(&services))
		router.Mount("/users", user.Router(&services))
	})

	service := chi.ServerBaseContext(router, serverCtx)
	http.ListenAndServe(getListenAddr(), service)
}
