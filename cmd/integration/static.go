package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type entityTemplate struct {
	common
	Method          string
	Route           string
	EntityTemplates int `json:"entity_templates"`
}

type abilityTemplate struct {
	common
	Method           string
	Route            string
	AbilityTemplates int `json:"ability_templates"`
}

type sector struct {
	common
	Method  string
	Route   string
	Sectors int
}

type sectorStarter struct {
	common
	Method   string
	Route    string
	Starters int
}

func expectEntityTemplates(a *LogAnalyzer) error {
	file, err := os.Open("./static/entity_templates.json")
	if err != nil {
		return err
	}
	defer file.Close()
	resp, err := http.Post("https://localhost:8081/entity/template", "application/json", file)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expected := entityTemplate{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "found",
		},
		Method:          "POST",
		Route:           "/entity/template",
		EntityTemplates: 5,
	}
	return a.Expect(func(s string) (bool, error) {
		var actual entityTemplate
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual != expected {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		return true, nil
	})
}

func expectAbilityTemplates(a *LogAnalyzer) error {
	file, err := os.Open("./static/ability_templates.json")
	if err != nil {
		return err
	}
	defer file.Close()
	resp, err := http.Post("https://localhost:8081/ability/template", "application/json", file)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expected := abilityTemplate{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "found",
		},
		Method:           "POST",
		Route:            "/ability/template",
		AbilityTemplates: 1,
	}
	return a.Expect(func(s string) (bool, error) {
		var actual abilityTemplate
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual != expected {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		return true, nil
	})
}

func expectSector(a *LogAnalyzer) error {
	file, err := os.Open("./static/sector.json")
	if err != nil {
		return err
	}
	defer file.Close()
	resp, err := http.Post("https://localhost:8081/sector", "application/json", file)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expected := sector{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "found",
		},
		Method:  "POST",
		Route:   "/sector",
		Sectors: 2,
	}
	return a.Expect(func(s string) (bool, error) {
		var actual sector
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual != expected {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		return true, nil
	})
}

func expectSectorStarter(a *LogAnalyzer) error {
	file, err := os.Open("./static/sector_starter.json")
	if err != nil {
		return err
	}
	defer file.Close()
	resp, err := http.Post("https://localhost:8081/sector/starter", "application/json", file)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expected := sectorStarter{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "found",
		},
		Method:   "POST",
		Route:    "/sector/starter",
		Starters: 1,
	}
	return a.Expect(func(s string) (bool, error) {
		var actual sectorStarter
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual != expected {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		return true, nil
	})
}

func expectStatic(a *LogAnalyzer) error {
	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec
	if err := expectEntityTemplates(a); err != nil {
		return err
	}
	if err := expectAbilityTemplates(a); err != nil {
		return err
	}
	if err := expectSector(a); err != nil {
		return err
	}
	if err := expectSectorStarter(a); err != nil {
		return err
	}
	return nil
}
