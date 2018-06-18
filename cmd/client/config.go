package main

import (
	"errors"

	game "github.com/elojah/game_01"
	"github.com/oklog/ulid"
)

// Config is web quic server structure config.
type Config struct {
	Token    game.ID `json:"token"`
	TickRate uint    `json:"tick_rate"`
	Address  string  `json:"address"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if c.Token.Compare(rhs.Token) != 0 {
		return false
	}
	if c.TickRate != rhs.TickRate {
		return false
	}
	if c.Address != rhs.Address {
		return false
	}
	return true
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cTickRate, ok := fconf["tick_rate"]
	if !ok {
		return errors.New("missing key tick_rate")
	}
	cTickRateFloat, ok := cTickRate.(float64)
	if !ok {
		return errors.New("key tick_rate invalid. must be numeric")
	}
	c.TickRate = uint(cTickRateFloat)
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
	return nil
}
