package login

import (
	"errors"
	"time"
)

// Config is web quic server structure config.
type Config struct {
	SigninURL string        `json:"signin_url"`
	Tolerance time.Duration `json:"tolerance"`

	Background string `json:"background"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c == rhs
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	var err error
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cSigninURL, ok := fconf["signin_url"]
	if !ok {
		return errors.New("missing key signin_url")
	}
	c.SigninURL, ok = cSigninURL.(string)
	if !ok {
		return errors.New("key signin_url invalid. must be string")
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

	cBackground, ok := fconf["background"]
	if !ok {
		return errors.New("missing key background")
	}
	c.Background, ok = cBackground.(string)
	if !ok {
		return errors.New("key background invalid. must be string")
	}

	return nil
}
