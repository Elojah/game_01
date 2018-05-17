package game

// EntityType represents the type of an entity.
type EntityType uint8

const (
	// Trickster represents a PJ with Trickster class.
	Trickster EntityType = 0
	// Mesmerist represents a PJ with Mesmerist class.
	Mesmerist EntityType = 1
	// Inquisitor represents a PJ with Inquisitor class.
	Inquisitor EntityType = 2
	// Totemist represents a PJ with Totemist class.
	Totemist EntityType = 3
	// Scavenger represents a PJ with Scavenger class.
	Scavenger EntityType = 4
)

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
