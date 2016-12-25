package space

import (
	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/models"
)

// ServiceMemory Address store uses an in memory store
type ServiceMemory struct{}

// Get Returns space details for given space id
func (*ServiceMemory) Get(spaceID string) (space *models.Space, err error) {

	for idx, record := range fixtures.Spaces {
		if record.SpaceID != spaceID {
			continue
		}

		space = &fixtures.Spaces[idx]
	}

	return space, nil
}

func (*ServiceMemory) GetMultiple(spaceIds []string) (spaces map[string]*models.Space, err error) {
	spaceMap := make(map[string]bool, len(spaceIds))
	spaces = make(map[string]*models.Space)

	for _, id := range spaceIds {
		spaceMap[id] = true
	}

	for idx, space := range fixtures.Spaces {
		if _, ok := spaceMap[space.SpaceID]; !ok {
			continue
		}

		spaces[space.SpaceID] = &fixtures.Spaces[idx]
	}

	return spaces, nil

}
