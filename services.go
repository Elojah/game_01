package game

// Services handles all services and rules interaction between them.
type Services struct {
	EntityService
	TokenService
	QEventService
	EventService
	AccountService
	ListenerService
}

// NewServices returns a new empty datagate..
func NewServices() Services {
	return Services{}
}
