package game

// Entity represents a dynamic entity.
type Entity struct {
	ID       ID
	HP       uint8
	MP       uint8
	Position Vec3
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
