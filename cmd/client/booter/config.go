package booter

import (
	"errors"
	"time"
)

// Config is web quic server structure config.
type Config struct {
	Address   string        `json:"address"`
	Tolerance time.Duration `json:"tolerance"`

	Title       string `json:"title"`
	Width       int32  `json:"width"`
	Height      int32  `json:"height"`
	FullScreen  bool   `json:"full_screen"`
	Resizable   bool   `json:"resizable"`
	Undecorated bool   `json:"undecorated"`
	VSync       bool   `json:"v_sync"`
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
	cWidthFloat64, ok := cWidth.(float64)
	if !ok {
		return errors.New("key width invalid. must be numeric")
	}
	c.Width = int32(cWidthFloat64)

	cHeight, ok := fconf["height"]
	if !ok {
		return errors.New("missing key height")
	}
	cHeightFloat64, ok := cHeight.(float64)
	if !ok {
		return errors.New("key height invalid. must be numeric")
	}
	c.Height = int32(cHeightFloat64)

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

	cVSync, ok := fconf["v_sync"]
	if !ok {
		return errors.New("missing key v_sync")
	}
	c.VSync, ok = cVSync.(bool)
	if !ok {
		return errors.New("key v_sync invalid. must be string")
	}

	return nil
}
