package game

const (
	// MaxPC is the maximum number of characters an account can have.
	MaxPC = 4
)

// PC alias an entity.
type PC Entity

// PCService is an interface to create a new PC.
type PCService interface {
	SetPC(PC, ID) error
	GetPC(PCSubset) (PC, error)
	ListPC(PCSubset) ([]PC, error)
}

// PCSubset represents a subset of PC by account ID.
type PCSubset struct {
	ID        ID
	AccountID ID
}

// PCLeft represents the number of character an account can still create.
type PCLeft int

// PCLeftService interfaces creation/retrieval of PCLeft.
type PCLeftService interface {
	SetPCLeft(PCLeft, ID) error
	GetPCLeft(PCLeftSubset) (PCLeft, error)
}

// PCLeftSubset represents a subset of PCLeft per account.
type PCLeftSubset struct {
	AccountID ID
}
