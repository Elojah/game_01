package game

// Services handles all services and rules interaction between them.
type Services struct {
	EntityService
	TokenService
	EventService
	AccountService
	ActionService
}

// NewServices returns a new empty datagate..
func NewServices() Services {
	return Services{}
}
