package errors

import (
	"fmt"

	perrors "github.com/pkg/errors"
)

// #Dev errors

// ErrNotImplementedYet is raised when a resource is not implemented yet.
type ErrNotImplementedYet struct {
	Version string
}

func (err ErrNotImplementedYet) Error() string {
	return fmt.Sprintf("not implemented in version %s", err.Version)
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
type ErrInsufficientACLs struct {
	Value  int
	Source string
	Target string
}

func (err ErrInsufficientACLs) Error() string {
	return fmt.Sprintf("insufficient rights %d for source %s on target %s", err.Value, err.Source, err.Target)
}

// ErrMultipleLogin is raised zhen an account is already logged.
type ErrMultipleLogin struct {
	AccountID string
}

func (err ErrMultipleLogin) Error() string {
	return fmt.Sprintf("account %s already logged", err.AccountID)
}

// ErrInvalidEntityType is raised when an entity doesn't respect the correct type.
type ErrInvalidEntityType struct {
	EntityType string
}

func (err ErrInvalidEntityType) Error() string {
	return fmt.Sprintf("invalid entity type %s", err.EntityType)
}

// #Common api/core errors

// ErrInvalidAction is raised when an action is not possible following game rules.
type ErrInvalidAction struct {
	Action string
}

func (err ErrInvalidAction) Error() string {
	return fmt.Sprintf("action %s is not possible", err.Action)
}

// ErrNotFound is raised when a mandatory resource is not found in storage.
type ErrNotFound struct {
	Store string
	Index string
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("no results found in store %s for index %s", err.Store, err.Index)
}

// ErrMissingTarget is raised when a ability is performed with a missing component targets.
type ErrMissingTarget struct {
	AbilityID   string
	ComponentID string
}

func (err ErrMissingTarget) Error() string {
	return fmt.Sprintf("missing target for ability %s component %s", err.AbilityID, err.ComponentID)
}

// ErrTooManyTargets is raised when a ability is casted/performed with too many targets for a component.
type ErrTooManyTargets struct {
	NTargets    int
	Max         uint64
	AbilityID   string
	ComponentID string
}

func (err ErrTooManyTargets) Error() string {
	return fmt.Sprintf("too many targets %d for ability %s component %s", err.NTargets, err.AbilityID, err.ComponentID)
}

// ErrOutOfRange is raised when a component ability is performed out of range on a target or when an item is looted in a further position.
type ErrOutOfRange struct {
	Dist  float64
	Range float64
}

func (err ErrOutOfRange) Error() string {
	return fmt.Sprintf("out of range %f for %f", err.Dist, err.Range)
}

// ErrMissingItem is raised when a loot action is performed on an non present item in target inventory.
type ErrMissingItem struct {
	ItemID      string
	InventoryID string
}

func (err ErrMissingItem) Error() string {
	return fmt.Sprintf("item %s missing in inventory %s", err.ItemID, err.InventoryID)
}

// ErrFullInventory is raised when a loot action is performed and the source inventory is full.
type ErrFullInventory struct {
	InventoryID string
}

func (err ErrFullInventory) Error() string {
	return fmt.Sprintf("inventory %s full", err.InventoryID)
}

// ErrIneffectiveCancel is raised when a cancel event is applied but no event was found to cancel.
type ErrIneffectiveCancel struct {
	TriggerID string
}

func (err ErrIneffectiveCancel) Error() string {
	return fmt.Sprintf("ineffective cancel trigger %s", err.TriggerID)
}

// IsGameLogicError returns if error type is a game logic error and needs a cancel propagation.
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
		_ = err
	}

	return false
}
