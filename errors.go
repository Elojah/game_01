package game

import (
	"errors"
)

var (
	// ErrWrongIP is raised when a source IP doesn't match with token-associated IP.
	ErrWrongIP = errors.New("ip don't match")
	// ErrInvalidTS is raised when a packet has a TS out of valid range.
	ErrInvalidTS = errors.New("packet TS is out of valid range")
)
