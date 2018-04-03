package game

// Services handles all services and rules interaction between them.
type Services struct {
	ActorService
}

// NewServices returns a new empty datagate..
func NewServices() Services {
	return Services{}
}
