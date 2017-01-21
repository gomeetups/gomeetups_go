package models

// Services Stores available services
type Services struct {
	GroupService   GroupService
	AddressService AddressService
	PhotoService   PhotoService
	SpaceService   SpaceService
	EventService   EventService
	UserService    UserService
	RsvpService    RsvpService
}
