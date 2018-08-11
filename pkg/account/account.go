package account

// Subset is the subset to retrieve an account.
type Subset struct {
	Username string
	Password string
}

// Store wraps account interactions.
type Store interface {
	SetAccount(A) error
	GetAccount(Subset) (A, error)
	DelAccount(Subset) error
}
