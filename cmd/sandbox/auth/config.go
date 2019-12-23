package auth

import (
	"errors"
)

// Config is client config.
// Tolerance in milliseconds.
type Config struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return (c.Username != rhs.Username &&
		c.Password != rhs.Password)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cAddress, ok := fconf["address"]
	if !ok {
		return errors.New("missing key address")
	}
	c.Address, ok = cAddress.(string)
	if !ok {
		return errors.New("key address invalid. must be string")
	}

	cUsername, ok := fconf["username"]
	if !ok {
		return errors.New("missing key username")
	}
	c.Username, ok = cUsername.(string)
	if !ok {
		return errors.New("key username invalid. must be string")
	}

	cPassword, ok := fconf["password"]
	if !ok {
		return errors.New("missing key password")
	}
	c.Password, ok = cPassword.(string)
	if !ok {
		return errors.New("key password invalid. must be string")
	}

	return nil
}
