package game

// Services handles all services and rules interaction between them.
type Services struct {
	ActorService
	TokenService
}

// NewServices returns a new empty datagate..
func NewServices() Services {
	return Services{}
}
