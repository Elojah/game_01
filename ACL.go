package game

// ACL is a enum for different rights between entities.
type ACL uint8

const (
	// Read right gives access to retrieve the value.
	Read ACL = 0
	// Update right gives access to update an already existing value.
	Update ACL = 1
	// Set right gives access to create a new entity.
	Set ACL = 2
	// Delete right gives access to delete an already existing entity.
	Delete ACL = 3
	// Owner right gives access to all rights.
	Owner ACL = 4
)
