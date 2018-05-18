package game

const (
	// MaxPC is the maximum number of characters an account can have.
	MaxPC = 4
	// PCLeftKey is the redis key which substitute Source in Permission.
	PCLeftKey = "pc_left"
)

// PC alias an entity.
type PC Entity

// PCService is an interface to create a new PC.
type PCService interface {
	SetPC(PC, ID) error
	ListPC(PCSubset) ([]PC, error)
}

// PCSubset represents a subset of PC by account ID.
type PCSubset struct {
	AccountID ID
}
