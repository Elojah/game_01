package main

import (
	"errors"
)

// Config wraps window configuration.
type Config struct {
	Title       string  `json:"title"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	FullScreen  bool    `json:"full_screen"`
	Resizable   bool    `json:"resizable"`
	Undecorated bool    `json:"undecorated"`
	VSync       bool    `json:"v_sync"`

	TickRate uint32 `json:"tick_rate"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c == rhs
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cTitle, ok := fconf["title"]
	if !ok {
		return errors.New("missing key title")
	}
	c.Title, ok = cTitle.(string)
	if !ok {
		return errors.New("key title invalid. must be string")
	}

	cWidth, ok := fconf["width"]
	if !ok {
		return errors.New("missing key width")
	}
	c.Width, ok = cWidth.(float64)
	if !ok {
		return errors.New("key width invalid. must be numeric")
	}

	cHeight, ok := fconf["height"]
	if !ok {
		return errors.New("missing key height")
	}
	c.Height, ok = cHeight.(float64)
	if !ok {
		return errors.New("key height invalid. must be numeric")
	}

	cFullScreen, ok := fconf["full_screen"]
	if !ok {
		return errors.New("missing key full_screen")
	}
	c.FullScreen, ok = cFullScreen.(bool)
	if !ok {
		return errors.New("key full_screen invalid. must be string")
	}

	cResizable, ok := fconf["resizable"]
	if !ok {
		return errors.New("missing key resizable")
	}
	c.Resizable, ok = cResizable.(bool)
	if !ok {
		return errors.New("key resizable invalid. must be string")
	}

	cUndecorated, ok := fconf["undecorated"]
	if !ok {
		return errors.New("missing key undecorated")
	}
	c.Undecorated, ok = cUndecorated.(bool)
	if !ok {
		return errors.New("key undecorated invalid. must be string")
	}

	cVSync, ok := fconf["vsync"]
	if !ok {
		return errors.New("missing key vsync")
	}
	c.VSync, ok = cVSync.(bool)
	if !ok {
		return errors.New("key vsync invalid. must be string")
	}

	cTickRate, ok := fconf["tick_rate"]
	if !ok {
		return errors.New("missing key tick_rate")
	}
	cTickRateFloat64, ok := cTickRate.(float64)
	if !ok {
		return errors.New("key tick_rate invalid. must be numeric")
	}
	c.TickRate = uint32(cTickRateFloat64)

	return nil
}
