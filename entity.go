package game

// Entity represents a dynamic entity.
type Entity struct {
	ID       ID
	HP       uint8
	MP       uint8
	Position Vec3
}

// MoveTo moves entity to position p.
func (e *Entity) MoveTo(p Vec3) {
	e.Position = p
}

// EntityService is an interface for Entity object.
type EntityService interface {
	CreateEntity(Entity, int64) error
	GetEntity(EntitySubset) (Entity, error)
}

// EntitySubset is a subset to retrieve one entity.
type EntitySubset struct {
	Key string
	Max int64
}
