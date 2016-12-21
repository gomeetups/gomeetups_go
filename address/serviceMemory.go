package address

import (
	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/models"
)

// ServiceMemory In memory address store
type ServiceMemory struct{}

func (*ServiceMemory) GetByGroupId(groupIds []string) (addresses map[string]*models.Address, err error) {

	groupMap := make(map[string]bool, len(groupIds))

	for _, id := range groupIds {
		groupMap[id] = true
	}

	addresses = make(map[string]*models.Address)

	for _, address := range fixtures.Addresses {
		// Check if address belongs to a group
		if address.AddressType != models.AddressTypes["GROUP"] {
			continue
		}

		// Check if address belongs to any of to given groups
		if _, ok := groupMap[address.BelongsTo]; !ok {
			continue
		}

		// Add address to addresses map, by using group id as key
		t := address
		addresses[address.BelongsTo] = &t
	}

	return addresses, nil
}

func (*ServiceMemory) Get(addressID string) (address models.Address, err error) {
	for id, record := range fixtures.Addresses {

		if id != addressID {
			continue
		}

		address = record
		break
	}

	return address, nil
}
