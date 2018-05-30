package game

// Sector represents a portion of map.
type Sector struct {
	Cuboid

	ID ID

	EntitiesID []ID

	// Adjacent faces
	Left  SideTransfer
	Right SideTransfer
	Up    SideTransfer
	Down  SideTransfer
	Front SideTransfer
	Back  SideTransfer

	// Adjacent edge up
	UpLeft  SideTransfer
	UpRight SideTransfer
	UpFront SideTransfer
	UpBack  SideTransfer

	// Adjacent edge down
	DownLeft  SideTransfer
	DownRight SideTransfer
	DownFront SideTransfer
	DownBack  SideTransfer

	// Adjacent edge side
	FrontLeft  SideTransfer
	FrontRight SideTransfer
	BackLeft   SideTransfer
	BackRight  SideTransfer
}

// SectorMapper interfaces Sector interactions.
type SectorMapper interface {
	CreateSector(Sector) error
	GetSector(SectorSubset) (Sector, error)
}

// SectorSubset is a subset to retrieve one or adjacent sectors.
type SectorSubset struct {
	ID ID
}

// SideTransfer represents data needed to move from one sector to another.
type SideTransfer struct {
	SectorIDs []ID
}
