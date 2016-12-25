package event

import (
	"errors"
	"strings"

	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/models"
)

//ServiceMemory In memory groups service
type ServiceMemory struct{}

//SearchEvents Finds all events for given filters
func (*ServiceMemory) SearchEvents(filters *models.ValidEventSearchParams) (events map[string]*models.Event, err error) {
	events = make(map[string]*models.Event)

	// Walk over the events fixtures array
	for idx, event := range fixtures.Events {

		// Check if name or description filter defined
		if filters.Name != "" || filters.Description != "" || filters.GroupID != "" {
			var doesMatch = false
			// Check if group name contains name filter

			if filters.GroupID != "" && strings.ToLower(event.GroupID) == strings.ToLower(filters.GroupID) {
				doesMatch = true
			}

			if !doesMatch && filters.Name != "" && strings.Contains(strings.ToLower(event.Name), strings.ToLower(filters.Name)) {
				doesMatch = true
			}

			if !doesMatch && filters.Description != "" && strings.Contains(strings.ToLower(event.Description), strings.ToLower(filters.Description)) {
				doesMatch = true
			}

			if doesMatch {
				events[event.EventID] = &fixtures.Events[idx]
			}

		} else {
			events[event.EventID] = &fixtures.Events[idx]
		}
	}

	return events, nil
}

//Get Fetches event details for given event Id
func (*ServiceMemory) Get(eventID string) (event models.Event, err error) {

	for _, record := range fixtures.Events {
		if record.EventID == eventID {
			return record, nil
		}
	}

	return models.Event{}, errors.New("Not found!")
}
