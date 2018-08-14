package svc

import (
	"net"
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	serrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// TokenService represents token usecases.
type TokenService struct {
	Account          account.Store
	AccountToken     account.TokenStore
	Entity           entity.Store
	EntityPC         entity.PCStore
	EntityPermission entity.PermissionStore

	EntityService        entity.Service
	InfraRecurrerService infra.RecurrerService
}

// New creates a new token from account payload A. Returns an error if the account is invalid.
func (s TokenService) New(payload account.A, addr string) (account.Token, error) {

	// #Search account in redis
	a, err := s.Account.GetAccount(account.Subset{
		Username: payload.Username,
	})
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "get account with username %s", payload.Username)
	}
	if a.Password != payload.Password {
		return account.Token{}, errors.Wrap(account.ErrWrongCredentials, "compare passwords")
	}
	if !a.Token.IsZero() {
		return account.Token{}, errors.Wrap(account.ErrMultipleLogin, "check existing account token")
	}

	// #Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "resolve address %s", addr)
	}

	// #Set a new token
	t := account.Token{
		ID:      ulid.NewID(),
		Account: a.ID,
		IP:      ip.String(),
	}
	if err := s.AccountToken.SetToken(t); err != nil {
		return account.Token{}, errors.Wrapf(err, "set token %s", t.ID.String())
	}
	a.Token = t.ID
	if err := s.Account.SetAccount(a); err != nil {
		return account.Token{}, errors.Wrapf(err, "set account %s with token %s", a.ID.String(), t.ID.String())
	}

	return t, nil
}

// Access retrieves a token and check IP validity.
func (s TokenService) Access(id ulid.ID, addr string) (account.Token, error) {

	// #Search message UUID in storage.
	t, err := s.AccountToken.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "get token %s", id.String())
	}

	// #Match message UUID with source IP.
	expected, _, ee := net.SplitHostPort(t.IP)
	actual, _, ea := net.SplitHostPort(addr)
	if expected != actual || ee != nil || ea != nil {
		return account.Token{}, errors.Wrapf(account.ErrWrongIP, "different ips %s != %s", expected, actual)
	}
	return t, nil
}

// Disconnect closes a token and all entities/listener/sync associated.
func (s TokenService) Disconnect(id ulid.ID) error {

	// Disconnect must be permissive in case of infra failures.
	var nonblockErr error

	// #Retrieve token
	t, err := s.AccountToken.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		return errors.Wrapf(err, "get token %s", id.String())
	}

	// #Close token recurrer
	if err := s.InfraRecurrerService.Remove(id); err != nil {
		nonblockErr = errors.Wrapf(err, "remove recurrer %s", id.String())
	}

	// #Reset token entity.
	te := t.Entity
	t.Entity = ulid.ID{}
	if err := s.AccountToken.SetToken(t); err != nil {
		nonblockErr = errors.Wrapf(err, "set token %s", id.String())
	}

	// #Retrieve entity
	e, err := s.Entity.GetEntity(entity.Subset{
		ID:    te,
		MaxTS: time.Now().UnixNano(),
	})
	if err != nil {
		// Token is valid but not connected to any entity.
		if err == serrors.ErrNotFound {
			return nil
		}
		return errors.Wrapf(err, "get entity %s", te.String())
	}

	// #Save last entity state into PC
	pc := entity.PC(e)
	pc.ID = t.PC
	if err := s.EntityPC.SetPC(pc, t.Account); err != nil {
		return errors.Wrapf(err, "set pc %s with entity %s", pc.ID.String(), e.ID.String())
	}

	// # Disconnect all entitis associated with token.
	ps, err := s.EntityPermission.ListPermission(entity.PermissionSubset{Source: t.ID.String()})
	if err != nil {
		return errors.Wrapf(err, "list permissions for token %s", t.ID.String())
	}
	for _, p := range ps {
		targetID := ulid.MustParse(p.Target)
		if err := s.EntityService.Disconnect(targetID, t); err != nil {
			nonblockErr = errors.Wrapf(err, "disconnect entity %s from token %s", targetID.String(), t.ID.String())
		}
	}

	return nonblockErr
}
