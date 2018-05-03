package game

// Account represents an user account.
type Account struct {
	ID       ID
	Username string
	Password string `json:"-"`
}

// AccountBuilder is the builder to retrieve an account.
type AccountBuilder struct {
	Username string
	Password string
}

// AccountService wraps account interactions.
type AccountService interface {
	CreateAccount(Account) error
	GetAccount(AccountBuilder) (Account, error)
}
