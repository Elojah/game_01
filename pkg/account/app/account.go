package app

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

var _ account.App = (*A)(nil)

// A implements account applications.
type A struct {
	account.Store
	account.TokenStore
	account.TokenHCStore

	Entity entity.App
}

// CreateToken creates app new token from account payload A. Returns an error if the account is invalid.
func (app A) CreateToken(payload account.A, addr string) (account.Token, error) {

	// #Search account in redis
	acc, err := app.Store.FetchAccount(payload.Username)
	if err != nil {
		return account.Token{}, errors.Wrap(err, "create token")
	}
	if acc.Password != payload.Password {
		return account.Token{}, errors.Wrap(gerrors.ErrWrongCredentials{Username: payload.Username}, "create token")
	}
	if !acc.Token.IsZero() {
		return account.Token{}, errors.Wrap(gerrors.ErrMultipleLogin{AccountID: acc.ID.String()}, "create token")
	}

	// #Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return account.Token{}, errors.Wrap(errors.Wrapf(err, "resolve address %s", addr), "create token")
	}

	// #Upsert new token
	t := account.Token{
		ID:      gulid.NewID(),
		Account: acc.ID,
		IP:      ip.String(),
	}
	if err := app.TokenStore.UpsertToken(t); err != nil {
		return account.Token{}, errors.Wrap(err, "create token")
	}
	acc.Token = t.ID
	if err := app.Store.UpsertAccount(acc); err != nil {
		return account.Token{}, errors.Wrap(err, "create token")
	}

	return t, nil
}

// FetchTokenFromAddr retrieves acc token and check IP validity.
func (app A) FetchTokenFromAddr(id gulid.ID, addr string) (account.Token, error) {

	// #Search message UUID in storage.
	t, err := app.TokenStore.FetchToken(id)
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

// DisconnectToken closes acc token and all entities/sequencer/sync associated.
func (app A) DisconnectToken(id gulid.ID) error {

	// Disconnect must be permissive in case of infra failures.
	var result *multierror.Error

	// #Retrieve token
	t, err := app.TokenStore.FetchToken(id)
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
	if err := app.TokenStore.UpsertToken(t); err != nil {
		result = multierror.Append(result, errors.Wrap(err, "disconnect token"))
	}

	// #Retrieve entity
	e, err := app.Entity.Fetch(te, ulid.Now())
	if err != nil {
		// Token is valid but not connected to any entity.
		switch errors.Cause(err).(type) {
		case gerrors.ErrNotFound:
			return result.ErrorOrNil()
		}
		return errors.Wrap(err, "disconnect token")
	}

	// #Save last entity state into PC
	pc := e
	pc.ID = t.PC
	if err := app.Entity.UpsertPC(pc, t.Account); err != nil {
		return errors.Wrap(err, "disconnect token")
	}

	// #Save last inventory state into MR store as pc and remove entity inv in MR store
	inv, err := app.Entity.FetchMRInventoryFromCache(e.InventoryID, e.ID)
	if err != nil {
		return errors.Wrap(err, "disconnect token")
	}
	if err := app.Entity.UpsertMRInventoryWithCache(pc.ID, inv); err != nil {
		return errors.Wrap(err, "disconnect token")
	}
	if err := app.Entity.RemoveMRInventory(e.ID); err != nil {
		return errors.Wrap(err, "disconnect token")
	}

	// #Disconnect all entities associated with token.
	ps, err := app.Entity.ListPermission(t.ID.String())
	if err != nil {
		return errors.Wrap(err, "disconnect token")
	}
	for _, p := range ps {
		targetID := gulid.MustParse(p.Target)
		if err := app.Entity.Disconnect(targetID); err != nil {
			result = multierror.Append(result, errors.Wrap(err, "disconnect token"))
			continue // don't remove permission in error case, it could lead to data loss
		}
		// #Delete token permission on entity
		if err := app.Entity.RemovePermission(p.Source, p.Target); err != nil {
			result = multierror.Append(result, errors.Wrap(err, "disconnect token"))
		}
	}

	return result.ErrorOrNil()
}
