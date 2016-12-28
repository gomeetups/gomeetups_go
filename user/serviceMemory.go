package user

import (
	"strings"

	"golang.org/x/crypto/bcrypt"

	"bitbucket.org/devmach/gomeetups/fixtures"
	"bitbucket.org/devmach/gomeetups/models"
)

// ServiceMemory User store uses an in memory store
type ServiceMemory struct{}

// GetByEmail Returns user data for given user email
func (*ServiceMemory) GetByEmail(email string) (user *models.User, err error) {
	for idx, record := range fixtures.Users {
		if record.IsActive == true && record.Email == email {
			return &fixtures.Users[idx], nil
		}
	}

	return nil, models.ErrUserNotFound
}

// GetByID Returns user data for given user id
func (*ServiceMemory) GetByID(userID string) (user *models.User, err error) {
	for idx, record := range fixtures.Users {
		if record.IsActive == true && record.UserID == userID {
			return &fixtures.Users[idx], nil
		}
	}

	return nil, models.ErrUserNotFound
}

// VerifyPassword Returns user data for given user email
func (*ServiceMemory) VerifyPassword(identifier string, plainTextPassword string) (err error) {
	var user *models.User

	for idx, record := range fixtures.Users {
		if strings.ToLower(record.UserID) == strings.ToLower(identifier) {
			user = &fixtures.Users[idx]
			break
		}

		if strings.ToLower(record.Email) == strings.ToLower(identifier) {
			user = &fixtures.Users[idx]
			break
		}
	}

	if user == nil {
		return models.ErrUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte((*user).Password), []byte(plainTextPassword)) != nil {
		return models.ErrUserInvalidPassword
	}

	return nil
}

// Search Searches users based on the given parameters
func (*ServiceMemory) Search(filters *models.ValidUserSearchParams) (users map[string]*models.User, err error) {
	users = make(map[string]*models.User)

	for idx, user := range fixtures.Users {
		if !user.IsActive {
			continue
		}

		if filters.DisplayName != "" {

			if strings.Contains(strings.ToLower(user.DisplayName), strings.ToLower(filters.DisplayName)) {
				users[user.UserID] = &fixtures.Users[idx]

			}

		} else {
			users[user.UserID] = &fixtures.Users[idx]
		}
	}

	return users, nil
}
