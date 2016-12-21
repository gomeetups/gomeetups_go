package group

import (
	"errors"
	"strings"

	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/models"
)

//ServiceMemory In memory groups service
type ServiceMemory struct{}

func (*ServiceMemory) SearchGroups(filter *models.ValidGroupSearchParams) (groups map[string]*models.Group, err error) {
	groups = make(map[string]*models.Group)

	for uuid, group := range fixtures.Groups {

		if filter.Name != "" || filter.Description != "" {
			var doesMatch = false

			if filter.Name != "" && strings.Contains(strings.ToLower(group.Name), strings.ToLower(filter.Name)) {
				doesMatch = true
			}

			if filter.Description != "" && strings.Contains(strings.ToLower(group.Name), strings.ToLower(filter.Description)) {
				doesMatch = true
			}

			if doesMatch {
				grp := group
				groups[uuid] = &grp
			}

		} else {
			grp := group
			groups[uuid] = &grp
		}
	}

	return groups, nil
}

func (*ServiceMemory) GroupDetails(groupId string) (group models.Group, err error) {
	if val, ok := fixtures.Groups[groupId]; ok {
		return val, nil
	}

	return models.Group{}, errors.New("Not found!")
}
