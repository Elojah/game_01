package game

// Actor represents a dynamic entity.
type Actor struct {
	ID       ID
	HP       uint8
	MP       uint8
	Position Vec3

	// Static belongs for no HP/MP (no damage/heal) and no move unless special skills.
	Static bool
}

// ActorSubset is a subset for Actor.
type ActorSubset struct {
	IDs    []ID
	Nearby *Circle
}

// ActorPatch is a patch for Actor.
type ActorPatch struct {
	HP       *uint8
	MP       *uint8
	Position *Vec3
}

// ActorService is a REST interface for Actor object.
type ActorService interface {
	CreateActor([]Actor) error
	UpdateActor(ActorSubset, ActorPatch) error
	DeleteActor(ActorSubset) error
	ListActor(ActorSubset) ([]Actor, error)
}
