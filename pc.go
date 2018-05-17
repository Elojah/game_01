package game

// PC alias an entity.
type PC Entity

// PCService is an interface to create a new PC.
type PCService interface {
	CreatePC(PC) error
}
