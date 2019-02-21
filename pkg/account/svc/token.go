package svc

import (
	"net"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/oklog/ulid"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TokenService represents token usecases.
type TokenService struct {
	AccountStore          account.Store
	AccountTokenStore     account.TokenStore
	EntityStore           entity.Store
	EntityPCStore         entity.PCStore
	EntityPermissionStore entity.PermissionStore

	EntityService entity.Service
}

// New creates a new token from account payload A. Returns an error if the account is invalid.
func (s TokenService) New(payload account.A, addr string) (account.Token, error) {

	// #Search account in redis
	a, err := s.AccountStore.GetAccount(payload.Username)
	if err != nil {
		return account.Token{}, errors.Wrap(err, "new token")
	}
	if a.Password != payload.Password {
		return account.Token{}, errors.Wrap(gerrors.ErrWrongCredentials{Username: payload.Username}, "new token")
	}
	if !a.Token.IsZero() {
		return account.Token{}, errors.Wrap(gerrors.ErrMultipleLogin{AccountID: a.ID.String()}, "new token")
	}

	// #Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return account.Token{}, errors.Wrap(errors.Wrapf(err, "resolve address %s", addr), "new token")
	}

	// #Set a new token
	t := account.Token{
		ID:      gulid.NewID(),
		Account: a.ID,
		IP:      ip.String(),
	}
	if err := s.AccountTokenStore.SetToken(t); err != nil {
		return account.Token{}, errors.Wrap(err, "new token")
	}
	a.Token = t.ID
	if err := s.AccountStore.SetAccount(a); err != nil {
		return account.Token{}, errors.Wrap(err, "new token")
	}

	return t, nil
}

// Access retrieves a token and check IP validity.
func (s TokenService) Access(id gulid.ID, addr string) (account.Token, error) {

	// #Search message UUID in storage.
	t, err := s.AccountTokenStore.GetToken(id)
	if err != nil {
		return account.Token{}, errors.Wrap(err, "access token")
	}

	// #Match message UUID with source IP.
	expected, _, ee := net.SplitHostPort(t.IP)
	actual, _, ea := net.SplitHostPort(addr)
	if expected != actual || ee != nil || ea != nil {
		return account.Token{}, errors.Wrap(
			gerrors.ErrWrongIP{TokenID: id.String(), Expected: expected, Actual: actual},
			"access token",
		)
	}
	return t, nil
}

// Disconnect closes a token and all entities/sequencer/sync associated.
func (s TokenService) Disconnect(id gulid.ID) error {

	// Disconnect must be permissive in case of infra failures.
	var result *multierror.Error

	// #Retrieve token
	t, err := s.AccountTokenStore.GetToken(id)
	if err != nil {
		return errors.Wrap(err, "disconnect token")
	}

	// #Check if token is connected
	if t.Entity.IsZero() {
		return nil
	}

	// #Reset token entity.
	te := t.Entity
	t.Entity = gulid.Zero()
	if err := s.AccountTokenStore.SetToken(t); err != nil {
		result = multierror.Append(result, errors.Wrap(err, "disconnect token"))
	}

	// #Retrieve entity
	e, err := s.EntityStore.GetEntity(te, ulid.Now())
	if err != nil {
		// Token is valid but not connected to any entity.
		switch errors.Cause(err).(type) {
		case gerrors.ErrNotFound:
			return result.ErrorOrNil()
		}
		return errors.Wrap(err, "disconnect token")
	}

	// #Save last entity state into PC
	pc := entity.PC(e)
	pc.ID = t.PC
	if err := s.EntityPCStore.SetPC(pc, t.Account); err != nil {
		return errors.Wrap(err, "disconnect token")
	}

	// # Disconnect all entities associated with token.
	ps, err := s.EntityPermissionStore.ListPermission(t.ID.String())
	if err != nil {
		return errors.Wrap(err, "disconnect token")
	}
	for _, p := range ps {
		targetID := gulid.MustParse(p.Target)
		if err := s.EntityService.Disconnect(targetID, t); err != nil {
			result = multierror.Append(result, errors.Wrap(err, "disconnect token"))
		}
	}

	return result.ErrorOrNil()
}
