package game

// Right is a enum for different rights between entities.
type Right uint8

const (
	// Read right gives access to retrieve the value.
	Read Right = 0
	// Update right gives access to update an already existing value.
	Update Right = 1
	// Set right gives access to create a new entity.
	Set Right = 2
	// Delete right gives access to delete an already existing entity.
	Delete Right = 3
	// Owner right gives access to all rights.
	Owner Right = 4
)
