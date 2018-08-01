package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/elojah/game_01/pkg/ulid"
)

var (
	testAccount = map[string]string{
		"username": "integration_test_username",
		"password": "integration_test_password",
	}
)

type account struct {
	common
	Account string
	Method  string
	Route   string
	Addr    string
}

type token struct {
	common
	Method   string
	Route    string
	Token    string
	Listener string
	Addr     string
}

func expectSubscribe(a *LogAnalyzer) error {
	raw, err := json.Marshal(testAccount)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8080/subscribe", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expected := account{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "subscribe success",
		},
		Method: "POST",
		Route:  "/subscribe",
	}
	return a.Expect(func(s string) (bool, error) {
		var at account
		if err := json.Unmarshal([]byte(s), &at); err != nil {
			return false, err
		}
		if at.common != expected.common {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		if _, err := ulid.Parse(at.Account); err != nil {
			return false, fmt.Errorf("invalid log account %s", s)
		}
		if _, err := net.ResolveTCPAddr("tcp", at.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
	})
}

func expectSignin(a *LogAnalyzer) error {
	raw, err := json.Marshal(testAccount)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8080/signin", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expected := token{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "signin success",
		},
		Method: "POST",
		Route:  "/signin",
	}
	count := 0
	return a.Expect(func(s string) (bool, error) {
		defer func() { count++ }()
		switch count {
		case 0:
			var tt token
			if err := json.Unmarshal([]byte(s), &tt); err != nil {
				return false, err
			}
			if tt.common != expected.common {
				return false, fmt.Errorf("unexpected log %s", s)
			}
			if _, err := ulid.Parse(tt.Token); err != nil {
				return false, fmt.Errorf("invalid token %s", s)
			}
			if _, err := ulid.Parse(tt.Listener); err != nil {
				return false, fmt.Errorf("invalid listener %s", s)
			}
			if _, err := net.ResolveTCPAddr("tcp", tt.Addr); err != nil {
				return false, fmt.Errorf("invalid log addr %s", s)
			}
			return true, nil
		case 1:
		}
	})
}

func expectUnsubscribe(a *LogAnalyzer) error {
	raw, err := json.Marshal(testAccount)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8080/unsubscribe", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expected := account{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "unsubscribe success",
		},
		Method: "POST",
		Route:  "/unsubscribe",
	}
	return a.Expect(func(s string) (bool, error) {
		var at account
		if err := json.Unmarshal([]byte(s), &at); err != nil {
			return false, err
		}
		if at.common != expected.common {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		if _, err := ulid.Parse(at.Account); err != nil {
			return false, fmt.Errorf("invalid log account %s", s)
		}
		if _, err := net.ResolveTCPAddr("tcp", at.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
	})
}

func expectAuth(a *LogAnalyzer) (ulid.ID, error) {
	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	if err := expectSubscribe(a); err != nil {
		return ulid.ID{}, err
	}
	if err := expectSignin(a); err != nil {
		return ulid.ID{}, err
	}
	if err := expectUnsubscribe(a); err != nil {
		return ulid.ID{}, err
	}
	return ulid.ID{}, nil
}
