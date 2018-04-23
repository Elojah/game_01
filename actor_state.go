package game

// ActorState represents a dynamic entity.
type ActorState struct {
	ActorID  ID
	StateID  ID
	HPMax    uint8
	HPLeft   uint8
	MPMax    uint8
	MPLeft   uint8
	Position Vec3
}

// ActorStateSubset is a subset for ActorState.
type ActorStateSubset struct {
	IDs    []ID
	Nearby *Circle
}

// ActorStatePatch is a patch for ActorState.
type ActorStatePatch struct {
	HP       *uint8
	MP       *uint8
	Position *Vec3
}

// ActorStateService is a REST interface for ActorState object.
type ActorStateService interface {
	CreateActorState([]ActorState) error
	UpdateActorState(ActorStateSubset, ActorStatePatch) error
	DeleteActorState(ActorStateSubset) error
	ListActorState(ActorStateSubset) ([]ActorState, error)
}
