package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/ulid"
)

var (
	testAccount = map[string]string{
		"username": "integration_test_username",
		"password": "integration_test_password",
	}
	testPCName = "testint"
	testPCType = ulid.MustParse("01CE3J5M6QMP5A4C0GTTYFYANP")
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
	Method   string
	Route    string
	Token    string
	Addr     string
	Template string
	PC       string
	Sector   string
}

type listPCLog struct {
	common
	Method  string
	Route   string
	Token   string
	Addr    string
	Account string
}

type createPC struct {
	Token ulid.ID
	Name  string
	Type  ulid.ID
}

type listPC struct {
	Token ulid.ID
}

type connectPC struct {
	Token  ulid.ID
	Target ulid.ID
}

type signoutAccount struct {
	Username string
	Token    ulid.ID
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
	expected := tokenLog{
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
		if actual.common != expected.common {
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
		Name:  testPCName,
		Type:  testPCType,
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
	expected := createPCLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "pc creation success",
		},
		Method:   "POST",
		Route:    "/pc/create",
		Template: testPCType.String(),
	}
	return a.Expect(func(s string) (bool, error) {
		var actual createPCLog
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual.common != expected.common {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		if _, err := ulid.Parse(actual.Token); err != nil {
			return false, fmt.Errorf("invalid token %s", s)
		}
		if _, err := ulid.Parse(actual.PC); err != nil {
			return false, fmt.Errorf("invalid pc %s", s)
		}
		if _, err := ulid.Parse(actual.Sector); err != nil {
			return false, fmt.Errorf("invalid sector %s", s)
		}
		if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
	})
}

func expectListPC(a *LogAnalyzer, tok account.Token) (entity.PC, error) {
	lpc := listPC{
		Token: tok.ID,
	}
	raw, err := json.Marshal(lpc)
	if err != nil {
		return entity.PC{}, err
	}
	resp, err := http.Post("https://localhost:8080/pc/list", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return entity.PC{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return entity.PC{}, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	var pcs []entity.PC
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&pcs); err != nil {
		return entity.PC{}, err
	}
	if len(pcs) == 0 {
		return entity.PC{}, errors.New("no pcs")
	}
	expected := listPCLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "pc list success",
		},
		Method: "POST",
		Route:  "/pc/list",
	}
	return pcs[0], a.Expect(func(s string) (bool, error) {
		var actual listPCLog
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if actual.common != expected.common {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		if _, err := ulid.Parse(actual.Token); err != nil {
			return false, fmt.Errorf("invalid token %s", s)
		}
		if _, err := ulid.Parse(actual.Account); err != nil {
			return false, fmt.Errorf("invalid account %s", s)
		}
		if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
			return false, fmt.Errorf("invalid log addr %s", s)
		}
		return true, nil
	})
}

func expectConnectPC(a *LogAnalyzer, tok account.Token, pc entity.PC) (entity.E, error) {
	cpc := connectPC{
		Token:  tok.ID,
		Target: pc.ID,
	}
	raw, err := json.Marshal(cpc)
	if err != nil {
		return entity.E{}, err
	}
	resp, err := http.Post("https://localhost:8080/pc/connect", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return entity.E{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return entity.E{}, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	var e entity.E
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		return entity.E{}, err
	}
	return e, nil
}

func expectSignout(a *LogAnalyzer, tok account.Token) error {
	sa := signoutAccount{
		Token:    tok.ID,
		Username: testAccount["username"],
	}
	raw, err := json.Marshal(sa)
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
			Message: "signout success",
		},
		Method: "POST",
		Route:  "/signout",
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
	pc, err := expectListPC(a, tok)
	if err != nil {
		return ulid.ID{}, err
	}
	_ = pc
	// e, err := expectConnectPC(a, tok, pc)
	// if err != nil {
	// 	return ulid.ID{}, err
	// }
	// _ = e
	if err := expectSignout(a, tok); err != nil {
		return ulid.ID{}, err
	}
	if err := expectUnsubscribe(a); err != nil {
		return ulid.ID{}, err
	}
	return ulid.ID{}, nil
}
