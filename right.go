package game

// Right represents a right access to a resource.
type Right uint8

const (
	// Read is the most basic access. It allows user to retrieve this resource.
	Read Right = iota
	// Update allows user to update this resource.
	Update
	// Create allows user to create a new resource of this type.
	Create
	// Delete allows user to delete this resource.
	Delete
)
