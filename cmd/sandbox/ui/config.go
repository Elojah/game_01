package ui

import "errors"

type Config struct {
	LevelFile string
	FPS       float64
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cFPS, ok := fconf["fps"]
	if !ok {
		return errors.New("missing key fps")
	}
	c.FPS, ok = cFPS.(float64)
	if !ok {
		return errors.New("key fps invalid. must be numeric")
	}
	cLevelFile, ok := fconf["level_file"]
	if !ok {
		return errors.New("missing key level_file")
	}
	c.LevelFile, ok = cLevelFile.(string)
	if !ok {
		return errors.New("key level_file invalid. must be string")
	}

	return nil
}
