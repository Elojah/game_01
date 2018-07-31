package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type common struct {
	Level   string
	Exe     string
	Method  string
	Route   string
	Message string
}

type entityTemplate struct {
	common
	EntityTemplates int `json:"entity_templates"`
}

type abilityTemplate struct {
	common
	AbilityTemplates int `json:"ability_templates"`
}

type sector struct {
	common
	Sectors int
}

type sectorStarter struct {
	common
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
			Method:  "POST",
			Route:   "/entity/template",
			Message: "found",
		},
		EntityTemplates: 5,
	}
	return a.Expect(func(s string) (bool, error) {
		var et entityTemplate
		if err := json.Unmarshal([]byte(s), &et); err != nil {
			return false, err
		}
		if et != expected {
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
			Method:  "POST",
			Route:   "/ability/template",
			Message: "found",
		},
		AbilityTemplates: 2,
	}
	return a.Expect(func(s string) (bool, error) {
		var et abilityTemplate
		if err := json.Unmarshal([]byte(s), &et); err != nil {
			return false, err
		}
		if et != expected {
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
			Method:  "POST",
			Route:   "/sector",
			Message: "found",
		},
		Sectors: 2,
	}
	return a.Expect(func(s string) (bool, error) {
		var et sector
		if err := json.Unmarshal([]byte(s), &et); err != nil {
			return false, err
		}
		if et != expected {
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
			Method:  "POST",
			Route:   "/sector/starter",
			Message: "found",
		},
		Starters: 1,
	}
	return a.Expect(func(s string) (bool, error) {
		var et sectorStarter
		if err := json.Unmarshal([]byte(s), &et); err != nil {
			return false, err
		}
		if et != expected {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		return true, nil
	})
}

func expectStaticData(a *LogAnalyzer) error {
	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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
