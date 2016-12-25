package photo

import (
	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/models"
)

// ServiceMemory Address store uses an in memory store
type ServiceMemory struct{}

// FindPhotos Find all photos for given parameters
func (*ServiceMemory) FindPhotos(photoType string, belongsTo string, resolution string) (photos map[string]*models.Photo, err error) {
	photos = make(map[string]*models.Photo)

	for idx, photo := range fixtures.Photos {
		if photo.PhotoType != photoType || photo.BelongsTo != belongsTo {
			continue
		}

		if resolution != "" && photo.Resolution != resolution {
			continue
		}

		photos[photo.PhotoID] = &fixtures.Photos[idx]
	}

	return photos, nil
}

// GetByGroupID Find all photos for given groups
func (*ServiceMemory) GetByGroupID(groupIds []string) (photos map[string]*models.Photo, err error) {
	photos = make(map[string]*models.Photo)
	groupMap := make(map[string]bool, len(groupIds))

	// Walk over groupIds and add an element into the lookup map
	for _, id := range groupIds {
		groupMap[id] = true
	}

	for idx, photo := range fixtures.Photos {
		if photo.PhotoType != models.PhotoTypes["GROUP"] {
			continue
		}

		if _, ok := groupMap[photo.BelongsTo]; !ok {
			continue
		}

		photos[photo.BelongsTo] = &fixtures.Photos[idx]
	}

	return photos, nil
}
