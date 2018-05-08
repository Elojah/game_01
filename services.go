package game

// Services handles all services and rules interaction between them.
type Services struct {
	AccountService
	EntityService
	EventService
	QEventService
	QListenerService
	SubscriptionService
	TokenService
}

// NewServices returns a new empty datagate..
func NewServices() Services {
	return Services{}
}
