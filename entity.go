package game

// EntityType represents the type of an entity.
type EntityType uint8

const (
	// Trickster represents a PC with Trickster class.
	Trickster EntityType = 0
	// Mesmerist represents a PC with Mesmerist class.
	Mesmerist EntityType = 1
	// Inquisitor represents a PC with Inquisitor class.
	Inquisitor EntityType = 2
	// Totemist represents a PC with Totemist class.
	Totemist EntityType = 3
	// Scavenger represents a PC with Scavenger class.
	Scavenger EntityType = 4
)

func (e EntityType) String() string {
	switch e {
	case Trickster:
		return "trickster"
	case Mesmerist:
		return "mesmerist"
	case Inquisitor:
		return "inquisitor"
	case Totemist:
		return "totemist"
	case Scavenger:
		return "scavenger"
	default:
		return "unknown"
	}
}

// Entity represents a dynamic entity.
type Entity struct {
	ID       ID
	HP       uint8
	MP       uint8
	Position Vec3
	Type     EntityType
}

// MoveTo moves entity to position p.
func (e *Entity) MoveTo(p Vec3) {
	e.Position = p
}

// EntityService is an interface for Entity object.
type EntityService interface {
	SetEntity(Entity, int64) error
	GetEntity(EntitySubset) (Entity, error)
}

// EntitySubset is a subset to retrieve one entity.
type EntitySubset struct {
	Key string
	Max int64
}
