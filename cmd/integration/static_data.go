package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type entityTemplate struct {
	Level           string
	Exe             string
	Method          string
	Route           string
	EntityTemplates int `json:"entity_templates"`
	Message         string
}

type abilityTemplate struct {
	Level            string
	Exe              string
	Method           string
	Route            string
	AbilityTemplates int `json:"ability_templates"`
	Message          string
}

type sector struct {
	Level   string
	Exe     string
	Method  string
	Route   string
	Sectors int
	Message string
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
		Level:           "info",
		Exe:             "./bin/game_tool",
		Method:          "POST",
		Route:           "/entity/template",
		EntityTemplates: 5,
		Message:         "found",
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
		Level:            "info",
		Exe:              "./bin/game_tool",
		Method:           "POST",
		Route:            "/ability/template",
		AbilityTemplates: 2,
		Message:          "found",
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
		Level:   "info",
		Exe:     "./bin/game_tool",
		Method:  "POST",
		Route:   "/sector",
		Sectors: 2,
		Message: "found",
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
	return nil
}
