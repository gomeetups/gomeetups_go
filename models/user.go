package models

import (
	"errors"
	"time"
)

var (
	ErrUserInvalidParam    = errors.New("Invalid parameter!")
	ErrUserInvalidPassword = errors.New("Invalid Password!")
)

// ValidUserSearchParams Valid search parameters for the events endpoint
type ValidUserSearchParams struct {
	DisplayName string
}

//User Defines users
type User struct {
	UserID      string    `json:"userId" db:"user_id"`
	Email       string    `json:"email" db:"email"`
	DisplayName string    `json:"displayName" db:"display_name"`
	Password    string    `json:"-" db:"password"`
	IsVerified  bool      `json:"isVerified" db:"is_verified"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	CreatedBy   string    `json:"createdBy" db:"created_by"` // User herself
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
	UpdatedBy   string    `json:"updatedBy" db:"updated_by"`

	Photos  []*Photo `json:"photos,omitempty" db:"-"`
	Address *Address `json:"address,omitempty" db:"-"`
}

//UserService contains necessary methods that have to be implemented by
// a user service...
type UserService interface {
	GetByEmail(email string) (user *User, err error)
	GetByID(userID string) (user *User, err error)
	VerifyPassword(identifier string, plainTextPassword string) (err error)
	Search(filters *ValidUserSearchParams) (users map[string]*User, err error)
}
