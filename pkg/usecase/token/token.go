package token

import (
	"net"
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	uce "github.com/elojah/game_01/pkg/usecase/entity"
	"github.com/elojah/game_01/pkg/usecase/listener"
	"github.com/elojah/game_01/pkg/usecase/recurrer"
	"github.com/pkg/errors"
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
	expected, _, ee := net.SplitHostPort(tok.IP.String())
	actual, _, ea := net.SplitHostPort(addr)
	if expected != actual || ee != nil || ea != nil {
		err := account.ErrWrongIP
		return account.Token{}, errors.Wrapf(err, "different ips %s != %s", expected, actual)
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
		err := account.ErrWrongCredentials
		return account.Token{}, errors.Wrap(err, "passwords don't match")
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
		IP:      ip,
	}
	if err := t.SetToken(token); err != nil {
		return account.Token{}, errors.Wrapf(err, "set token %s", token.ID.String())
	}

	return token, nil
}

// Disconnect closes a token and all entities/listener/sync associated.
func (t T) Disconnect(id ulid.ID) error {

	// #Retrieve token
	tok, err := t.GetToken(account.TokenSubset{ID: id})
	if err != nil {
		return errors.Wrap(err, "disconnect")
	}

	// #Close token listener
	if err := t.L.Delete(id); err != nil {
		return errors.Wrap(err, "disconnect")
	}

	// #Close token recurrer
	if err := t.R.Delete(id); err != nil {
		return errors.Wrap(err, "disconnect")
	}

	// #Retrieve entity
	e, err := t.EntityMapper.GetEntity(entity.Subset{
		ID:    tok.Entity,
		MaxTS: time.Now().UnixNano(),
	})
	if err != nil {
		return errors.Wrapf(err, "get entity %s", tok.Entity.String())
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
			return errors.Wrapf(err, "disconnect entity %s from token %s", targetID.String(), tok.ID.String())
		}
	}

	if err := t.DelToken(account.TokenSubset{ID: id}); err != nil {
		return errors.Wrapf(err, "delete token %s", id.String())
	}

	return nil
}
