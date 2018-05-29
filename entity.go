package game

// EntityType represents the type of an entity.
type EntityType = ID

// Entity represents a dynamic entity.
type Entity struct {
	ID       ID         `json:"id"`
	Type     EntityType `json:"type"`
	Name     string     `json:"name"`
	HP       uint64     `json:"hp"`
	MP       uint64     `json:"mp"`
	Position Vec3       `json:"position"`
}

// MoveTo moves entity to position p.
func (e *Entity) MoveTo(p Vec3) {
	e.Position = p
}

// EntityMapper is an interface for Entity object.
type EntityMapper interface {
	SetEntity(Entity, int64) error
	GetEntity(EntitySubset) (Entity, error)
}

// EntitySubset is a subset to retrieve one entity.
type EntitySubset struct {
	Key   string
	MaxTS int64
}
