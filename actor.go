package game

type ActorService interface {
	CreateActor(...Actor) error
	UpdateActor(ActorSubset, ActorPatch) error
	DeleteActor(ActorSubset) error
	ListActor(ActorSubset) ([]byte, error)
}
