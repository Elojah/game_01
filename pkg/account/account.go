package account

// Subset is the subset to retrieve an account.
type Subset struct {
	Username string
	Password string
}

// Mapper wraps account interactions.
type Mapper interface {
	SetAccount(A) error
	GetAccount(Subset) (A, error)
	DelAccount(Subset) error
}
