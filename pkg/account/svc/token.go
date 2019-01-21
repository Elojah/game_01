package svc

import (
	"net"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/oklog/ulid"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/infra"
	gulid "github.com/elojah/game_01/pkg/ulid"
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
	a, err := s.Account.GetAccount(payload.Username)
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "get account with username %s", payload.Username)
	}
	if a.Password != payload.Password {
		return account.Token{}, errors.Wrap(gerrors.ErrWrongCredentials{Username: payload.Username}, "compare passwords")
	}
	if !a.Token.IsZero() {
		return account.Token{}, errors.Wrap(gerrors.ErrMultipleLogin{a.ID.String()}, "check existing account token")
	}

	// #Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "resolve address %s", addr)
	}

	// #Set a new token
	t := account.Token{
		ID:      gulid.NewID(),
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
func (s TokenService) Access(id gulid.ID, addr string) (account.Token, error) {

	// #Search message UUID in storage.
	t, err := s.AccountToken.GetToken(id)
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "get token %s", id.String())
	}

	// #Match message UUID with source IP.
	expected, _, ee := net.SplitHostPort(t.IP)
	actual, _, ea := net.SplitHostPort(addr)
	if expected != actual || ee != nil || ea != nil {
		return account.Token{}, gerrors.ErrWrongIP{TokenID: id.String(), Expected: expected, Actual: actual}
	}
	return t, nil
}

// Disconnect closes a token and all entities/sequencer/sync associated.
func (s TokenService) Disconnect(id gulid.ID) error {

	// Disconnect must be permissive in case of infra failures.
	var result *multierror.Error

	// #Retrieve token
	t, err := s.AccountToken.GetToken(id)
	if err != nil {
		return errors.Wrapf(err, "get token %s", id.String())
	}

	// #Close token recurrer
	if err := s.InfraRecurrerService.Remove(id); err != nil {
		result = multierror.Append(result, errors.Wrapf(err, "remove recurrer %s", id.String()))
	}

	// #Reset token entity.
	te := t.Entity
	t.Entity = gulid.Zero()
	if err := s.AccountToken.SetToken(t); err != nil {
		result = multierror.Append(result, errors.Wrapf(err, "set token %s", id.String()))
	}

	// #Retrieve entity
	e, err := s.Entity.GetEntity(te, ulid.Now())
	if err != nil {
		// Token is valid but not connected to any entity.
		switch errors.Cause(err).(type) {
		case gerrors.ErrNotFound:
			return result.ErrorOrNil()
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
	ps, err := s.EntityPermission.ListPermission(t.ID.String())
	if err != nil {
		return errors.Wrapf(err, "list permissions for token %s", t.ID.String())
	}
	for _, p := range ps {
		targetID := gulid.MustParse(p.Target)
		if err := s.EntityService.Disconnect(targetID, t); err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "disconnect entity %s from token %s", targetID.String(), t.ID.String()))
		}
	}

	return result.ErrorOrNil()
}
