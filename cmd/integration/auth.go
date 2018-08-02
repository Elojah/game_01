package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

var (
	testAccount = map[string]string{
		"username": "integration_test_username",
		"password": "integration_test_password",
	}
)

type accountLog struct {
	common
	Account string
	Method  string
	Route   string
	Addr    string
}

type tokenLog struct {
	common
	Method   string
	Route    string
	Token    string
	Listener string
	Addr     string
}

type listenerLog struct {
	common
	Core     string
	Listener string
	Action   uint8
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
	expected := accountLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "subscribe success",
		},
		Method: "POST",
		Route:  "/subscribe",
	}
	return a.Expect(func(s string) (bool, error) {
		var actual accountLog
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual.common != expected.common {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		if _, err := ulid.Parse(actual.Account); err != nil {
			return false, fmt.Errorf("invalid log account %s", s)
		}
		if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
	})
}

func expectSignin(a *LogAnalyzer) (account.Token, error) {
	var tok account.Token
	raw, err := json.Marshal(testAccount)
	if err != nil {
		return tok, err
	}
	resp, err := http.Post("https://localhost:8080/signin", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return tok, err
	}
	if resp.StatusCode != http.StatusOK {
		return tok, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return tok, err
	}
	expectedToken := tokenLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "signin success",
		},
		Method: "POST",
		Route:  "/signin",
	}
	expectedListener := listenerLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "listening",
		},
		Action: 0,
	}
	count := 0
	return tok, a.Expect(func(s string) (bool, error) {
		defer func() { count++ }()
		switch count {
		case 0:
			var actual tokenLog
			if err := json.Unmarshal([]byte(s), &actual); err != nil {
				return false, err
			}
			if actual.common != expectedToken.common {
				return false, fmt.Errorf("unexpected log %s", s)
			}
			if _, err := ulid.Parse(actual.Token); err != nil {
				return false, fmt.Errorf("invalid token %s", s)
			}
			if _, err := ulid.Parse(actual.Listener); err != nil {
				return false, fmt.Errorf("invalid listener %s", s)
			}
			if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
				return false, fmt.Errorf("invalid log addr %s", s)
			}
			return false, nil
		case 1:
			var actual listenerLog
			if err := json.Unmarshal([]byte(s), &actual); err != nil {
				return false, err
			}
			if actual.common != expectedListener.common {
				return false, fmt.Errorf("unexpected log %s", s)
			}
			if _, err := ulid.Parse(actual.Core); err != nil {
				return false, fmt.Errorf("invalid core %s", s)
			}
			if _, err := ulid.Parse(actual.Listener); err != nil {
				return false, fmt.Errorf("invalid listener %s", s)
			}
			if actual.Action != expectedListener.Action {
				return false, fmt.Errorf("invalid action %s", s)
			}
			return true, nil
		default:
			return false, fmt.Errorf("additional log %s", s)
		}
	})
}

func expectSignout(a *LogAnalyzer, tok account.Token) error {
	return nil
	raw, err := json.Marshal(testAccount)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8080/signout", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expectedToken := tokenLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "signin success",
		},
		Method: "POST",
		Route:  "/signin",
	}
	expectedListener := listenerLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "listening",
		},
		Action: 0,
	}
	count := 0
	return a.Expect(func(s string) (bool, error) {
		defer func() { count++ }()
		switch count {
		case 0:
			var actual tokenLog
			if err := json.Unmarshal([]byte(s), &actual); err != nil {
				return false, err
			}
			if actual.common != expectedToken.common {
				return false, fmt.Errorf("unexpected log %s", s)
			}
			if _, err := ulid.Parse(actual.Token); err != nil {
				return false, fmt.Errorf("invalid tokenLog %s", s)
			}
			if _, err := ulid.Parse(actual.Listener); err != nil {
				return false, fmt.Errorf("invalid listenerLog %s", s)
			}
			if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
				return false, fmt.Errorf("invalid log addr %s", s)
			}
			return false, nil
		case 1:
			var actual listenerLog
			if err := json.Unmarshal([]byte(s), &actual); err != nil {
				return false, err
			}
			if actual.common != expectedListener.common {
				return false, fmt.Errorf("unexpected log %s", s)
			}
			if _, err := ulid.Parse(actual.Core); err != nil {
				return false, fmt.Errorf("invalid core %s", s)
			}
			if _, err := ulid.Parse(actual.Listener); err != nil {
				return false, fmt.Errorf("invalid listenerLog %s", s)
			}
			if actual.Action != expectedListener.Action {
				return false, fmt.Errorf("invalid action %s", s)
			}
			return true, nil
		default:
			return false, fmt.Errorf("additional log %s", s)
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
	expected := accountLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "unsubscribe success",
		},
		Method: "POST",
		Route:  "/unsubscribe",
	}
	return a.Expect(func(s string) (bool, error) {
		var actual accountLog
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual.common != expected.common {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		if _, err := ulid.Parse(actual.Account); err != nil {
			return false, fmt.Errorf("invalid log account %s", s)
		}
		if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
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
	tok, err := expectSignin(a)
	if err != nil {
		return ulid.ID{}, err
	}
	if err := expectSignout(a, tok); err != nil {
		return ulid.ID{}, err
	}
	if err := expectUnsubscribe(a); err != nil {
		return ulid.ID{}, err
	}
	return ulid.ID{}, nil
}
