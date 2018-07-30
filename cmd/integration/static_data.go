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

func expectStaticData(a *LogAnalyzer) error {
	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	file, err := os.Open("./static/entity_templates.json")
	if err != nil {
		return err
	}
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
