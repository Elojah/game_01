package game

// Account represents an user account.
type Account struct {
	Username string
	Password string
}

// AccountBuilder is the builder to retrieve an account.
type AccountBuilder struct {
	Username string
	Password string
}

// AccountService wraps account interactions.
type AccountService interface {
	GetAccount(AccountBuilder) (Account, error)
}
