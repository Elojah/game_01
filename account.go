package game

// Account represents an user account.
type Account struct {
	ID       ID
	Username string
	Password string `json:"-"`
}

// AccountSubset is the subset to retrieve an account.
type AccountSubset struct {
	Username string
	Password string
}

// AccountMapper wraps account interactions.
type AccountMapper interface {
	SetAccount(Account) error
	GetAccount(AccountSubset) (Account, error)
}
