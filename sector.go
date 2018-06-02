package game

// Direction represents a direction relative to a sector. Must have 26 values.
type Direction uint8

const (
	// Back is the back direction relative to sector.
	Back Direction = 0
	// Down is the down direction relative to sector.
	Down Direction = 1
	// DownBack is the down back direction relative to sector.
	DownBack Direction = 2
	// DownFront is the down front direction relative to sector.
	DownFront Direction = 3
	// DownLeft is the down left direction relative to sector.
	DownLeft Direction = 4
	// DownLeftBack is the down left back direction relative to sector.
	DownLeftBack Direction = 5
	// DownLeftFront is the down left front direction relative to sector.
	DownLeftFront Direction = 6
	// DownRight is the down right direction relative to sector.
	DownRight Direction = 7
	// DownRightBack is the down right back direction relative to sector.
	DownRightBack Direction = 8
	// DownRightFront is the down right front direction relative to sector.
	DownRightFront Direction = 9
	// Front is the front direction relative to sector.
	Front Direction = 10
	// Left is the left direction relative to sector.
	Left Direction = 11
	// LeftBack is the left back direction relative to sector.
	LeftBack Direction = 12
	// LeftFront is the left front direction relative to sector.
	LeftFront Direction = 13
	// Right is the right direction relative to sector.
	Right Direction = 14
	// RightBack is the right back direction relative to sector.
	RightBack Direction = 15
	// RightFront is the right front direction relative to sector.
	RightFront Direction = 16
	// Up is the up direction relative to sector.
	Up Direction = 17
	// UpBack is the up back direction relative to sector.
	UpBack Direction = 18
	// UpFront is the up front direction relative to sector.
	UpFront Direction = 19
	// UpLeft is the up left direction relative to sector.
	UpLeft Direction = 20
	// UpLeftBack is the up left back direction relative to sector.
	UpLeftBack Direction = 21
	// UpLeftFront is the up left front direction relative to sector.
	UpLeftFront Direction = 22
	// UpRight is the up right direction relative to sector.
	UpRight Direction = 23
	// UpRightBack is the up right back direction relative to sector.
	UpRightBack Direction = 24
	// UpRightFront is the up right front direction relative to sector.
	UpRightFront Direction = 25
)

// ExitPoint represents a central point to another sector.
type ExitPoint struct {
	SectorID ID
	Position Vec3
}

// Sector represents a cuboid in the world.
type Sector struct {
	ID         ID
	Size       Vec3
	ExitPoints [26][]ExitPoint
}

// SectorMapper is the service for Sector.
type SectorMapper interface {
	SetSector(Sector) error
	GetSector(SectorSubset) (Sector, error)
}

// SectorSubset allows to retrieve on sector by ID.
type SectorSubset struct {
	ID ID
}
