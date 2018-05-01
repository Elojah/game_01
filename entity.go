package game

// Entity represents a dynamic entity.
type Entity struct {
	ID       ID
	HP       uint8
	MP       uint8
	Position Vec3

	// Static belongs for no HP/MP (no damage/heal) and no move unless special skills.
	Static bool
}

// EntitySubset is a subset for Entity.
type EntitySubset struct {
	IDs    []ID
	Nearby *Circle
}

// EntityPatch is a patch for Entity.
type EntityPatch struct {
	HP       *uint8
	MP       *uint8
	Position *Vec3
}

// EntityService is a REST interface for Entity object.
type EntityService interface {
	CreateEntity([]Entity) error
	UpdateEntity(EntitySubset, EntityPatch) error
	DeleteEntity(EntitySubset) error
	ListEntity(EntitySubset) ([]Entity, error)
}
