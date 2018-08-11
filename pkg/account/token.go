package account

import (
	"net"
	"time"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

// TokenStore is the service gate for Token resource.
type TokenStore interface {
	SetToken(Token) error
	GetToken(TokenSubset) (Token, error)
	DelToken(TokenSubset) error
}

// TokenSubset retrieves a token per ID.
type TokenSubset struct {
	ID ulid.ID
}

// TokenHCStore is the service gate for Token health check.
type TokenHCStore interface {
	SetTokenHC(ulid.ID, int64) error
	ListTokenHC(TokenHCSubset) ([]ulid.ID, error)
}

// TokenHCSubset retrieves token healthchecks based on last tick.
type TokenHCSubset struct {
	MaxTS int64
}

// TokenService represents token usecases.
type TokenService struct {
	AccountStore          Store
	TokenStore            TokenStore
	EntityStore           entity.Store
	EntityPCStore         entity.PCStore
	EntityPermissionStore entity.PermissionStore

	EntityService        entity.Service
	InfraRecurrerService infra.RecurrerService
}

// New creates a new token from account payload A. Returns an error if the account is invalid.
func (s TokenService) New(accountPayload A, addr string) (Token, error) {

	// #Search account in redis
	a, err := s.AccountStore.GetAccount(Subset{
		Username: accountPayload.Username,
	})
	if err != nil {
		return Token{}, errors.Wrapf(err, "get account with username %s", accountPayload.Username)
	}
	if a.Password != accountPayload.Password {
		return Token{}, errors.Wrap(ErrWrongCredentials, "compare passwords")
	}
	if !a.Token.IsZero() {
		return Token{}, errors.Wrap(ErrMultipleLogin, "check account state")
	}

	// #Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return Token{}, errors.Wrapf(err, "resolve address %s", addr)
	}

	// #Set a new token
	t := Token{
		ID:      ulid.NewID(),
		Account: a.ID,
		IP:      ip.String(),
	}
	if err := s.TokenStore.SetToken(t); err != nil {
		return Token{}, errors.Wrapf(err, "set token %s", t.ID.String())
	}
	a.Token = t.ID
	if err := s.AccountStore.SetAccount(a); err != nil {
		return Token{}, errors.Wrapf(err, "set account %s with token %s", a.ID.String(), token.ID.String())
	}

	return t, nil
}

// Access retrieves a token and check IP validity.
func (s TokenService) Access(id ulid.ID, addr string) (Token, error) {

	// #Search message UUID in storage.
	t, err := s.TokenStore.GetToken(TokenSubset{ID: id})
	if err != nil {
		return Token{}, errors.Wrapf(err, "get token %s", id.String())
	}

	// #Match message UUID with source IP.
	expected, _, ee := net.SplitHostPort(t.IP)
	actual, _, ea := net.SplitHostPort(addr)
	if expected != actual || ee != nil || ea != nil {
		return Token{}, errors.Wrapf(ErrWrongIP, "different ips %s != %s", expected, actual)
	}
	return t, nil
}

// Disconnect closes a token and all entities/listener/sync associated.
func (s TokenService) Disconnect(id ulid.ID) error {

	var softErr error

	// #Retrieve token
	t, err := s.TokenStore.GetToken(TokenSubset{ID: id})
	if err != nil {
		return errors.Wrap(err, "disconnect")
	}

	// #Close token recurrer
	if err := s.InfraRecurrerService.Remove(id); err != nil {
		softErr = errors.Wrap(err, "disconnect")
	}

	// #Reset token entity.
	te := t.Entity
	t.Entity = ulid.ID{}
	if err := s.TokenService.SetToken(t); err != nil {
		softErr = err
	}

	// #Retrieve entity
	e, err := s.EntityStore.GetEntity(entity.Subset{
		ID:    te,
		MaxTS: time.Now().UnixNano(),
	})
	if err != nil {
		if err == storage.ErrNotFound {
			return nil
		}
		return errors.Wrapf(err, "get entity %s", te.String())
	}

	// #Save last entity state into PC
	pc := entity.PC(e)
	pc.ID = t.PC
	if err := s.EntityPCStore.SetPC(pc, t.Account); err != nil {
		return errors.Wrapf(err, "set pc %s from entity %s", pc.ID.String(), e.ID.String())
	}

	// #For each entity permission
	ps, err := s.EntityPermissionStore.ListPermission(entity.PermissionSubset{Source: t.ID.String()})
	if err != nil {
		return errors.Wrapf(err, "list permissions for token %s", t.ID.String())
	}
	for _, p := range ps {
		targetID := ulid.MustParse(p.Target)
		if err := s.EntityService.Disconnect(targetID, t); err != nil {
			softErr = errors.Wrapf(err, "disconnect entity %s from token %s", targetID.String(), t.ID.String())
		}
	}

	return softErr
}
