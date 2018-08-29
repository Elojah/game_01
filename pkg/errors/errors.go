package errors

import (
	"errors"
)

var (
	// ErrNotImplementedYet is raised when a resource is not implemented yet.
	ErrNotImplementedYet = errors.New("not implemented yet")

	// ErrWrongIP is raised when a source IP doesn't match with token-associated IP.
	ErrWrongIP = errors.New("ip don't match")
	// ErrInvalidTS is raised when a packet has a TS out of valid range.
	ErrInvalidTS = errors.New("packet TS is out of valid range")
	// ErrWrongCredentials is raised when user logs with invalid username/password.
	ErrWrongCredentials = errors.New("invalid credentials")
	// ErrInsufficientACLs is raised when a user apply an action without valid rights.
	ErrInsufficientACLs = errors.New("insufficient rights")
	// ErrInvalidAction is raised when an action is not possible following game rules.
	ErrInvalidAction = errors.New("action is not possible")
	// ErrMultipleLogin is raised zhen an account is already logged.
	ErrMultipleLogin = errors.New("account already logged")
	// ErrInvalidEntityType is raised when an entity doesn't respect the correct type.
	ErrInvalidEntityType = errors.New("invalid entity type")
	// ErrNotFound is raised when a mandatory resource is not found in storage.
	ErrNotFound = errors.New("no results found")
	// ErrMissingTarget is raised when a ability is performed with a missing component targets.
	ErrMissingTarget = errors.New("target missing")
)
