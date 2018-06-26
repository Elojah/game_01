package game

// EntityType represents the type of an entity.
type EntityType = ID

// Position represents an entity position in world.
type Position struct {
	Coord    Vec3
	SectorID ID
}

// Entity represents a dynamic entity.
type Entity struct {
	ID       ID         `json:"id"`
	Type     EntityType `json:"type"`
	Name     string     `json:"name"`
	HP       uint64     `json:"hp"`
	MP       uint64     `json:"mp"`
	Position Position   `json:"position"`
}

// Move moves entity to position p.
func (e *Entity) Move(p Vec3) {
	e.Position.Coord.Add(p)
}

// EntityMapper is an interface for Entity object.
type EntityMapper interface {
	SetEntity(Entity, int64) error
	GetEntity(EntitySubset) (Entity, error)
}

// EntitySubset is a subset to retrieve one entity.
type EntitySubset struct {
	Key    string
	MaxTS  int64
	Cursor uint64
	Count  int64
}

// Equal returns if both entities are equal.
func (e Entity) Equal(entity Entity) bool {
	if e.ID.Compare(entity.ID) != 0 {
		return false
	}
	if e.Type.Compare(entity.Type) != 0 {
		return false
	}
	if e.Name != entity.Name {
		return false
	}
	if e.HP != entity.HP {
		return false
	}
	if e.MP != entity.MP {
		return false
	}
	if e.Position.SectorID.Compare(entity.Position.SectorID) != 0 {
		return false
	}
	if e.Position.Coord.X != entity.Position.Coord.X {
		return false
	}
	if e.Position.Coord.Y != entity.Position.Coord.Y {
		return false
	}
	if e.Position.Coord.Z != entity.Position.Coord.Z {
		return false
	}
	return true
}
