package errors

import (
	"fmt"

	perrors "github.com/pkg/errors"
)

// #Dev errors

// ErrNotImplementedYet is raised when a resource is not implemented yet.
type ErrNotImplementedYet struct{}

func (ErrNotImplementedYet) Error() string {
	return "not implemented yet"
}

// #Token/login errors

// ErrWrongIP is raised when a source IP doesn't match with token-associated IP.
type ErrWrongIP struct {
	TokenID  string
	Expected string
	Actual   string
}

func (err ErrWrongIP) Error() string {
	return fmt.Sprintf("token %s ip is %s, received %s", err.TokenID, err.Expected, err.Actual)
}

// ErrInvalidTS is raised when a packet has a TS out of valid range.
type ErrInvalidTS struct {
	MsgID string
	TS    uint64
	Now   uint64
}

func (err ErrInvalidTS) Error() string {
	return fmt.Sprintf("msg %s with TS %d received too late at %d", err.MsgID, err.TS, err.Now)
}

// ErrWrongCredentials is raised when user logs with invalid username/password.
type ErrWrongCredentials struct {
	Username string
}

func (err ErrWrongCredentials) Error() string {
	return fmt.Sprintf("invalid credentials for user %s", err.Username)
}

// ErrInsufficientACLs is raised when a user apply an action without valid rights.
type ErrInsufficientACLs struct{}

func (ErrInsufficientACLs) Error() string {
	return "insufficient rights"
}

// ErrMultipleLogin is raised zhen an account is already logged.
type ErrMultipleLogin struct {
	AccountID string
}

func (err ErrMultipleLogin) Error() string {
	return fmt.Sprintf("account %s already logged", err.AccountID)
}

// ErrInvalidEntityType is raised when an entity doesn't respect the correct type.
type ErrInvalidEntityType struct{}

func (ErrInvalidEntityType) Error() string {
	return "invalid entity type"
}

// #Common api/core errors

// ErrInvalidAction is raised when an action is not possible following game rules.
type ErrInvalidAction struct{}

func (ErrInvalidAction) Error() string {
	return "action is not possible"
}

// ErrNotFound is raised when a mandatory resource is not found in storage.
type ErrNotFound struct{}

func (ErrNotFound) Error() string {
	return "no results found"
}

// ErrMissingTarget is raised when a ability is performed with a missing component targets.
type ErrMissingTarget struct{}

func (ErrMissingTarget) Error() string {
	return "target missing"
}

// ErrTooManyTargets is raised when a ability is casted/performed with too many targets for a component.
type ErrTooManyTargets struct{}

func (ErrTooManyTargets) Error() string {
	return "too many targets"
}

// ErrOutOfRange is raised when a component ability is performed out of range on a target or when an item is looted in a further position.
type ErrOutOfRange struct{}

func (ErrOutOfRange) Error() string {
	return "out of range"
}

// ErrMissingItem is raised when a loot action is performed on an non present item in target inventory.
type ErrMissingItem struct{}

func (ErrMissingItem) Error() string {
	return "item missing"
}

// ErrFullInventory is raised when a loot action is performed and the source inventory is full.
type ErrFullInventory struct{}

func (ErrFullInventory) Error() string {
	return "inventory full"
}

// ErrIneffectiveCancel is raised when a cancel event is applied but no event was found to cancel.
type ErrIneffectiveCancel struct{}

func (ErrIneffectiveCancel) Error() string {
	return "ineffective cancel"
}

// ErrNotCancellable is raised when a cancel event is applied but the event is a non cancellable.
type ErrNotCancellable struct{}

func (ErrNotCancellable) Error() string {
	return "event can't be cancel"
}

func IsGameLogicError(err error) bool {
	switch err := perrors.Cause(err).(type) {
	case ErrNotImplementedYet:
	case ErrWrongIP:
	case ErrInvalidTS:
	case ErrWrongCredentials:
	case ErrInsufficientACLs:
	case ErrMultipleLogin:
	case ErrInvalidEntityType:
	case ErrInvalidAction:
	case ErrNotFound:
	case ErrMissingTarget:
	case ErrTooManyTargets:
	case ErrOutOfRange:
	case ErrMissingItem:
	case ErrFullInventory:
	case ErrIneffectiveCancel:
	case ErrNotCancellable:
	}
}
