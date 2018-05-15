package game

// Entity represents a dynamic entity.
type Entity struct {
	ID       ID
	HP       uint8
	MP       uint8
	Position Vec3
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
	CreateEntity(Entity, int64) error
	GetEntity(EntityBuilder) (Entity, error)
}

// EntityBuilder is a builder to retrieve one entity.
type EntityBuilder struct {
	Key string
	Max int
}
