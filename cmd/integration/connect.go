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

func (expected accountLog) Equal(actual accountLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Account); err != nil {
		return fmt.Errorf("invalid account %s", actual.Account)
	}
	if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	return nil
}

type tokenLog struct {
	common
	Method string
	Route  string
	Token  string
	Addr   string
}

func (expected tokenLog) Equal(actual tokenLog) error {
	if actual.Exe != expected.Exe {
		return fmt.Errorf("unexpected exe %s", actual.Exe)
	}
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Token); err != nil {
		return fmt.Errorf("invalid token %s", actual.Token)
	}
	if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	return nil
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

func (expected createPCLog) Equal(actual createPCLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.PC); err != nil {
		return fmt.Errorf("invalid pc %s", actual.PC)
	}
	if _, err := ulid.Parse(actual.Sector); err != nil {
		return fmt.Errorf("invalid sector %s", actual.Sector)
	}
	if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	return nil
}

type listPCLog struct {
	common
	Method  string
	Route   string
	Token   string
	Addr    string
	Account string
}

func (expected listPCLog) Equal(actual listPCLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Account); err != nil {
		return fmt.Errorf("invalid account %s", actual.Account)
	}
	if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	return nil
}

type connectPCLog struct {
	common
	Method    string
	Route     string
	Token     string
	Addr      string
	PC        string
	Entity    string
	Sector    string
	Sequencer string
}

func (expected connectPCLog) Equal(actual connectPCLog) error {
	if actual.Exe != expected.Exe {
		return fmt.Errorf("unexpected exe %s", actual.Exe)
	}
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Entity); err != nil {
		return fmt.Errorf("invalid entity %s", actual.Entity)
	}
	if _, err := ulid.Parse(actual.Sector); err != nil {
		return fmt.Errorf("invalid sector %s", actual.Sector)
	}
	if _, err := ulid.Parse(actual.Sequencer); err != nil {
		return fmt.Errorf("invalid sequencer %s", actual.Sequencer)
	}
	if _, err := net.ResolveTCPAddr("tcp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	return nil
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

type recurrerLog struct {
	common
	Sync     string
	Recurrer string
	Addr     string
	Time     int64
}

func (expected recurrerLog) Equal(actual recurrerLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	// config sync id
	if actual.Sync != expected.Sync {
		return fmt.Errorf("invalid sync %s", actual.Sync)
	}
	if _, err := ulid.Parse(actual.Recurrer); err != nil {
		return fmt.Errorf("invalid recurrer %s", actual.Recurrer)
	}
	if _, err := net.ResolveUDPAddr("udp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	// ignore time
	return nil
}

type sequencerLog struct {
	common
	Core      string
	Sequencer string
	Addr      string
	Time      int64
}

func (expected sequencerLog) Equal(actual sequencerLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	// config core id
	if actual.Core != expected.Core {
		return fmt.Errorf("invalid core %s", actual.Core)
	}
	if _, err := ulid.Parse(actual.Sequencer); err != nil {
		return fmt.Errorf("invalid sequencer %s", actual.Sequencer)
	}
	if _, err := net.ResolveUDPAddr("udp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	// ignore time
	return nil
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
		if err := expected.Equal(actual); err != nil {
			return false, err
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
		if err := expected.Equal(actual); err != nil {
			return false, err
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
		Token:    tok.ID.String(),
	}
	return a.Expect(func(s string) (bool, error) {
		var actual createPCLog
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if err := expected.Equal(actual); err != nil {
			return false, err
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
		Token:  tok.ID.String(),
	}
	return pcs[0], a.Expect(func(s string) (bool, error) {
		var actual listPCLog
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, err
		}
		if err := expected.Equal(actual); err != nil {
			return false, err
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
	expected := connectPCLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_auth",
			Message: "connect pc success",
		},
		Method: "POST",
		Route:  "/pc/connect",
		Token:  tok.ID.String(),
		PC:     pc.ID.String(),
	}
	expectedRec := recurrerLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_sync",
			Message: "recurrer up",
		},
		Sync: "01CGBKPGMC4TG961QMQ71TEV2P",
	}
	expectedSeq := sequencerLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "sequencer up",
		},
		Core: "01CCS1SZ4B20G98XYMFGVC9VS4",
	}
	n := 3
	return e, a.Expect(func(s string) (bool, error) {
		n--
		var com common
		if err := json.Unmarshal([]byte(s), &com); err != nil {
			return false, err
		}
		switch com.Exe {
		case "./bin/game_auth":
			var actual connectPCLog
			if err := json.Unmarshal([]byte(s), &actual); err != nil {
				return n == 0, err
			}
			return n == 0, expected.Equal(actual)
		case "./bin/game_sync":
			var actual recurrerLog
			if err := json.Unmarshal([]byte(s), &actual); err != nil {
				return n == 0, err
			}
			return n == 0, expectedRec.Equal(actual)
		case "./bin/game_core":
			var actual sequencerLog
			if err := json.Unmarshal([]byte(s), &actual); err != nil {
				return n == 0, err
			}
			return n == 0, expectedSeq.Equal(actual)
		default:
			return false, fmt.Errorf("unexpected exe %s", com.Exe)
		}
	})
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
	expected := tokenLog{
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
		if err := expected.Equal(actual); err != nil {
			return false, err
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
		if err := expected.Equal(actual); err != nil {
			return false, err
		}
		return true, nil
	})
}

func expectConnect(a *LogAnalyzer) (account.Token, entity.E, error) {
	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec
	if err := expectSubscribe(a); err != nil {
		return account.Token{}, entity.E{}, err
	}
	tok, err := expectSignin(a)
	if err != nil {
		return account.Token{}, entity.E{}, err
	}
	if err := expectCreatePC(a, tok); err != nil {
		return account.Token{}, entity.E{}, err
	}
	pc, err := expectListPC(a, tok)
	if err != nil {
		return account.Token{}, entity.E{}, err
	}
	e, err := expectConnectPC(a, tok, pc)
	if err != nil {
		return account.Token{}, entity.E{}, err
	}
	return tok, e, nil
}

func expectDisconnect(a *LogAnalyzer, tok account.Token) error {
	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec
	if err := expectSignout(a, tok); err != nil {
		return err
	}
	if err := expectUnsubscribe(a); err != nil {
		return err
	}
	return nil
}
