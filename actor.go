package game

// ActorService is a REST interface for Actor object.
type ActorService interface {
	CreateActor(...Actor) error
	UpdateActor(ActorSubset, ActorPatch) error
	DeleteActor(ActorSubset) error
	ListActor(ActorSubset) ([]byte, error)
}
