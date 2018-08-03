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
	Method string
	Route  string
	Token  string
	Addr   string
}

type createPCLog struct {
	common
	Method string
	Route  string
	Token  string
	Addr   string
}

type createPC struct {
	Token ulid.ID
	Name  string
	Type  ulid.ID
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
	return tok, a.Expect(func(s string) (bool, error) {
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
		if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
	})
}

func expectCreatePC(a *LogAnalyzer, tok account.Token) error {
	cpc := createPC{
		Token: tok.ID,
		Name:  "testint",
		Type:  ulid.MustParse("01CE3J5M6QMP5A4C0GTTYFYANP"),
	}
	raw, err := json.Marshal(cpc)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8080/pc/create", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	expectedCreatePC := createPCLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "pc creation success",
		},
		Method: "POST",
		Route:  "/pc/create",
	}
	return a.Expect(func(s string) (bool, error) {
		var actual createPCLog
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual.common != expectedCreatePC.common {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		if _, err := ulid.Parse(actual.Token); err != nil {
			return false, fmt.Errorf("invalid token %s", s)
		}
		if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
	})
}

func expectSignout(a *LogAnalyzer) error {
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
	return a.Expect(func(s string) (bool, error) {
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
		if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
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
	if err := expectCreatePC(a, tok); err != nil {
		return ulid.ID{}, err
	}
	if err := expectSignout(a); err != nil {
		return ulid.ID{}, err
	}
	if err := expectUnsubscribe(a); err != nil {
		return ulid.ID{}, err
	}
	return ulid.ID{}, nil
}
