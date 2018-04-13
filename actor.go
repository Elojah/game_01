package game

type Actor struct{}
type ActorSubset struct{}
type ActorPatch struct{}

// ActorService is a REST interface for Actor object.
type ActorService interface {
	CreateActor([]Actor) error
	UpdateActor(ActorSubset, ActorPatch) error
	DeleteActor(ActorSubset) error
	ListActor(ActorSubset) ([]Actor, error)
}
