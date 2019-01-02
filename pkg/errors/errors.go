package errors

import (
	"errors"
)

var (
	// #Dev errors

	// ErrNotImplementedYet is raised when a resource is not implemented yet.
	ErrNotImplementedYet = errors.New("not implemented yet")

	// #Token/login errors

	// ErrWrongIP is raised when a source IP doesn't match with token-associated IP.
	ErrWrongIP = errors.New("ip don't match")
	// ErrInvalidTS is raised when a packet has a TS out of valid range.
	ErrInvalidTS = errors.New("packet TS is out of valid range")
	// ErrWrongCredentials is raised when user logs with invalid username/password.
	ErrWrongCredentials = errors.New("invalid credentials")
	// ErrInsufficientACLs is raised when a user apply an action without valid rights.
	ErrInsufficientACLs = errors.New("insufficient rights")
	// ErrMultipleLogin is raised zhen an account is already logged.
	ErrMultipleLogin = errors.New("account already logged")
	// ErrInvalidEntityType is raised when an entity doesn't respect the correct type.
	ErrInvalidEntityType = errors.New("invalid entity type")

	// #Common api/core errors

	// ErrInvalidAction is raised when an action is not possible following game rules.
	ErrInvalidAction = errors.New("action is not possible")
	// ErrNotFound is raised when a mandatory resource is not found in storage.
	ErrNotFound = errors.New("no results found")
	// ErrMissingTarget is raised when a ability is performed with a missing component targets.
	ErrMissingTarget = errors.New("target missing")
	// ErrTooManyTargets is raised when a ability is casted/performed with too many targets for a component.
	ErrTooManyTargets = errors.New("too many targets")
	// ErrOutOfRange is raised when a component ability is performed out of range on a target or when an item is looted in a further position.
	ErrOutOfRange = errors.New("out of range")
	// ErrMissingItem is raised when a loot action is performed on an non present item in target inventory.
	ErrMissingItem = errors.New("item missing")
	// ErrFullInventory is raised when a loot action is performed and the source inventory is full.
	ErrFullInventory = errors.New("inventory full")
	// ErrIneffectiveCancel is raised when a cancel event is applied but no event was found to cancel.
	ErrIneffectiveCancel = errors.New("ineffective cancel")
	// ErrNotCancellable is raised when a cancel event is applied but the event is a non cancellable.
	ErrNotCancellable = errors.New("event can't be cancel")
)
