package main

import (
	"errors"
	"time"

	"github.com/elojah/game_01/pkg/ulid"
)

// Config is web quic server structure config.
type Config struct {
	Token     ulid.ID       `json:"token"`
	Address   string        `json:"address"`
	Tolerance time.Duration `json:"tolerance"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if ulid.Compare(c.Token, rhs.Token) != 0 {
		return false
	}
	return (c.Address != rhs.Address &&
		c.Tolerance == rhs.Tolerance)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cToken, ok := fconf["token"]
	if !ok {
		return errors.New("missing key token")
	}
	cTokenString, ok := cToken.(string)
	if !ok {
		return errors.New("key token invalid. must be string")
	}
	var err error
	if c.Token, err = ulid.Parse(cTokenString); err != nil {
		return errors.New("key token invalid. must be ulid")
	}
	cAddress, ok := fconf["address"]
	if !ok {
		return errors.New("missing key address")
	}
	c.Address, ok = cAddress.(string)
	if !ok {
		return errors.New("key address invalid. must be string")
	}

	cTolerance, ok := fconf["tolerance"]
	if !ok {
		return errors.New("missing key tolerance")
	}
	cToleranceString, ok := cTolerance.(string)
	if !ok {
		return errors.New("key tolerance invalid. must be string")
	}
	c.Tolerance, err = time.ParseDuration(cToleranceString)
	if err != nil {
		return err
	}

	return nil
}
