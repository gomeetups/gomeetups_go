package group

import (
	"errors"
	"strings"

	"github.com/gomeetups/gomeetups/fixtures"
	"github.com/gomeetups/gomeetups/models"
)

//ServiceMemory In memory groups service
type ServiceMemory struct{}

//SearchGroups Finds all groups for fiven filters
func (*ServiceMemory) SearchGroups(filter *models.ValidGroupSearchParams) (groups map[string]*models.Group, err error) {
	// Initialize return map, contains pointers to the groups in the store
	groups = make(map[string]*models.Group)

	// Walk over the groups fixtures array
	for idx, group := range fixtures.Groups {

		// Check if name or description filter defined
		if filter.Name != "" || filter.Description != "" {
			var doesMatch = false

			// Check if group name contains name filter
			if filter.Name != "" && strings.Contains(strings.ToLower(group.Name), strings.ToLower(filter.Name)) {
				doesMatch = true
			}

			// Check if group description contains group filter
			if filter.Description != "" && strings.Contains(strings.ToLower(group.Name), strings.ToLower(filter.Description)) {
				doesMatch = true
			}

			// If name or description filters matches with group's , add it to results
			if doesMatch {
				groups[group.GroupID] = &fixtures.Groups[idx]
			}

		} else {
			// There is no filter given, just return each group.
			groups[group.GroupID] = &fixtures.Groups[idx]
		}
	}

	return groups, nil
}

//Get Fetches group details for given group Id
func (*ServiceMemory) Get(groupID string) (group models.Group, err error) {

	// Walk over groups fixtures and return group with it's address as soon as
	// found
	for _, grp := range fixtures.Groups {
		if grp.GroupID == groupID {
			return grp, nil
		}
	}

	return models.Group{}, errors.New("Not found!")
}
