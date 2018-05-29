package game

// Sector represents a portion of map.
type Sector struct {
	Cuboid

	ID ID

	EntitiesID []ID
}

// SectorMapper interfaces Sector interactions.
type SectorMapper interface {
}

type SectorSubset {

}
