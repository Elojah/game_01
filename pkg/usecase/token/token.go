package token

import (
	"net"
	"time"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
	uce "github.com/elojah/game_01/pkg/usecase/entity"
	"github.com/elojah/game_01/pkg/usecase/listener"
	"github.com/elojah/game_01/pkg/usecase/recurrer"
)

// T wraps use cases around token object.
type T struct {
	AccountMapper account.Mapper
	account.TokenMapper

	EntityMapper entity.Mapper
	entity.PCMapper

	listener.L
	recurrer.R

	entity.PermissionMapper

	sector.EntitiesMapper
}

// Get retrieves a token and check IP validity.
func (t T) Get(id ulid.ID, addr string) (account.Token, error) {

	// #Search message UUID in storage.
	tok, err := t.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "get token %s", id.String())
	}

	// #Match message UUID with source IP.
	expected, _, ee := net.SplitHostPort(tok.IP)
	actual, _, ea := net.SplitHostPort(addr)
	if expected != actual || ee != nil || ea != nil {
		return account.Token{}, errors.Wrapf(account.ErrWrongIP, "different ips %s != %s", expected, actual)
	}
	return tok, nil
}

// New creates a new token from account payload A. Returns an error if the account is invalid.
func (t T) New(accountPayload account.A, addr string) (account.Token, error) {

	// #Search account in redis
	a, err := t.AccountMapper.GetAccount(account.Subset{
		Username: accountPayload.Username,
	})
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "get account with username %s", accountPayload.Username)
	}
	if a.Password != accountPayload.Password {
		return account.Token{}, errors.Wrap(account.ErrWrongCredentials, "compare passwords")
	}
	if !ulid.IsZero(a.Token) {
		return account.Token{}, errors.Wrap(account.ErrMultipleLogin, "check account state")
	}

	// #Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return account.Token{}, errors.Wrapf(err, "resolve address %s", addr)
	}

	// #Set a new token
	token := account.Token{
		ID:      ulid.NewID(),
		Account: a.ID,
		IP:      ip.String(),
	}
	if err := t.SetToken(token); err != nil {
		return account.Token{}, errors.Wrapf(err, "set token %s", token.ID.String())
	}
	a.Token = token.ID
	if err := t.AccountMapper.SetAccount(a); err != nil {
		return account.Token{}, errors.Wrapf(err, "set account %s with token %s", a.ID.String(), token.ID.String())
	}

	return token, nil
}

// Disconnect closes a token and all entities/listener/sync associated.
func (t T) Disconnect(id ulid.ID) error {

	var softErr error

	// #Retrieve token
	tok, err := t.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		return errors.Wrap(err, "disconnect")
	}

	// #Close token listener
	if err := t.L.Delete(id); err != nil {
		softErr = errors.Wrap(err, "disconnect")
	}

	// #Close token recurrer
	if err := t.R.Delete(id); err != nil {
		softErr = errors.Wrap(err, "disconnect")
	}

	// #Reset token entity.
	tokEntity := tok.Entity
	tok.Entity = ulid.ID{}
	if err := t.SetToken(tok); err != nil {
		softErr = err
	}

	// #Retrieve entity
	e, err := t.EntityMapper.GetEntity(entity.Subset{
		ID:    tokEntity,
		MaxTS: time.Now().UnixNano(),
	})
	if err != nil {
		if err == storage.ErrNotFound {
			return nil
		}
		return errors.Wrapf(err, "get entity %s", tokEntity.String())
	}

	// #Save last entity state into PC
	pc := entity.PC(e)
	pc.ID = tok.PC
	if err := t.SetPC(pc, tok.Account); err != nil {
		return errors.Wrapf(err, "set pc %s from entity %s", pc.ID.String(), e.ID.String())
	}

	// #For each entity permission
	permissions, err := t.ListPermission(entity.PermissionSubset{Source: tok.ID.String()})
	if err != nil {
		return errors.Wrapf(err, "list permissions for token %s", tok.ID.String())
	}
	ucentity := uce.E{
		EntityMapper:     t.EntityMapper,
		PermissionMapper: t.PermissionMapper,
		EntitiesMapper:   t.EntitiesMapper,
		L:                t.L,
	}
	for _, permission := range permissions {
		targetID := ulid.MustParse(permission.Target)
		if err := ucentity.Disconnect(targetID, tok); err != nil {
			softErr = errors.Wrapf(err, "disconnect entity %s from token %s", targetID.String(), tok.ID.String())
		}
	}

	return softErr
}
