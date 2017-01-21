package rsvp

import (
	"github.com/gomeetups/gomeetups/fixtures"
	"github.com/gomeetups/gomeetups/models"
)

type ServiceMemory struct{}

func (*ServiceMemory) GetByEventID(eventID string) (rsvps []*models.Rsvp, err error) {

	for _, rsvp := range fixtures.Rsvps {
		if rsvp.EventID != eventID {
			continue
		}
		rsvps = append(rsvps, &rsvp)
	}

	return rsvps, nil
}
