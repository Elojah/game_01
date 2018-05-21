package game

// Mappers handles all services and rules interaction between them.
type Mappers struct {
	AccountMapper
	EntityMapper
	EventMapper
	QEventMapper
	QListenerMapper
	SubscriptionMapper
	TokenMapper
}

// NewMappers returns a new empty datagate..
func NewMappers() Mappers {
	return Mappers{}
}
