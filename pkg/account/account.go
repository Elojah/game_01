package account

// Store wraps account interactions.
type Store interface {
	SetAccount(A) error
	GetAccount(string) (A, error)
	DelAccount(string) error
}
