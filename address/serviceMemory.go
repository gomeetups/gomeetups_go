package address

import (
	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/models"
)

// ServiceMemory Address store uses an in memory store
type ServiceMemory struct{}

// GetByGroupID Find all addresses belongs to given group ids..
func (*ServiceMemory) GetByGroupID(groupIds []string) (addresses map[string]*models.Address, err error) {

	// Create a map of group id's so we can check if group id of an address is
	// relevant later. Since we interested only in map's keys, we set map type to
	// a primative type like bool.
	groupMap := make(map[string]bool, len(groupIds))

	// Walk over groupIds and add an element into the lookup map
	for _, id := range groupIds {
		groupMap[id] = true
	}

	// Initialize our return map, contains pointers to the addresses in the store.
	addresses = make(map[string]*models.Address)

	// Walk over the addresses array
	for idx, address := range fixtures.Addresses {
		// Check if address belongs to a group
		if address.AddressType != models.AddressTypes["GROUP"] {
			continue
		}

		// Check if address belongs to any of to given groups
		if _, ok := groupMap[address.BelongsTo]; !ok {
			continue
		}

		// Add address to addresses map, by using group id as key.
		addresses[address.BelongsTo] = &fixtures.Addresses[idx]
	}

	return addresses, nil
}

// Get Fetch address details for a given address id
func (*ServiceMemory) Get(addressID string) (address models.Address, err error) {
	// Walk over the addresses array
	for _, record := range fixtures.Addresses {

		// If that's not the address that we are lookin, skip it...
		if record.AddressID != addressID {
			continue
		}

		// Copy record value to address
		address = record

		// Quit the for loop
		break
	}

	// Return the adress. Error set to nil in memory store
	return address, nil
}
