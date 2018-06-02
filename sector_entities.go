package game

// SectorEntities represents a entity sector in a world.
type SectorEntities struct {
	SectorID  ID
	EntityIDs []ID
}

// SectorEntitiesMapper set or get sector entities. Can also add or remove individual entity to sector.
type SectorEntitiesMapper interface {
	GetSectorEntities(SectorEntitiesSubset) (SectorEntities, error)
	AddEntityToSector(ID, ID) error
	RemoveEntityToSector(ID, ID) error
}

// SectorEntitiesSubset retrieves one SectorEntities per sector ID only.
type SectorEntitiesSubset struct {
	SectorID ID
}
