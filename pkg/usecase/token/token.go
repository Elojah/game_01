package token

import (
	"net"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	uce "github.com/elojah/game_01/pkg/usecase/entity"
	"github.com/elojah/game_01/pkg/usecase/listener"
)

// T wraps use cases around token object.
type T struct {
	AccountMapper account.Mapper
	account.TokenMapper

	EntityMapper entity.Mapper
	entity.PCMapper

	listener.L

	entity.PermissionMapper

	sector.EntitiesMapper
}

func (t T) New(accountPayload account.A, addr string) (account.Token, error) {

	logger := log.With().
		Str("account", accountPayload.ID.String()).
		Str("action", "token").
		Logger()

	// #Search account in redis
	a, err := t.AccountMapper.GetAccount(account.Subset{
		Username: accountPayload.Username,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to get account")
		return account.Token{}, err
	}
	if a.Password != accountPayload.Password {
		err := account.ErrWrongCredentials
		logger.Error().Err(err).Msg("failed to authenticate")
		return account.Token{}, err
	}

	// #Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		logger.Error().Err(err).Str("address", addr).Msg("failed to get valid IP")
		return account.Token{}, err
	}

	// #Set a new token
	token := account.Token{
		ID:      ulid.NewID(),
		Account: a.ID,
		IP:      ip,
	}
	if err := t.SetToken(token); err != nil {
		logger.Error().Err(err).Msg("failed to create token")
		return account.Token{}, err
	}

	return token, nil
}

// Disconnect closes a token and all entities/listener/sync associated.
func (t T) Disconnect(id ulid.ID) error {
	logger := log.With().
		Str("token", id.String()).
		Str("action", "disconnect").
		Logger()

	// #Close token listener
	go func() {
		if err := t.L.Delete(id); err != nil {
			logger.Error().Err(err).Msg("failed to close listener")
		}
	}()

	// #Close token recurrer
	go func() {
		if err := t.PublishRecurrer(event.Recurrer{ID: tok.ID, Action: event.Close}, tok.SyncPool); err != nil {
			logger.Error().Err(err).Msg("failed to close recurrer")
		}
	}()

	// #Retrieve entity
	e, err := t.EntityMapper.GetEntity(entity.Subset{
		ID:    tok.Entity,
		MaxTS: time.Now().UnixNano(),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve entity")
		return err
	}

	// #Save last entity state into PC
	pc := entity.PC(e)
	pc.ID = tok.PC
	if err := t.SetPC(pc, tok.Account); err != nil {
		logger.Error().Err(err).Msg("failed to save pc")
		return err
	}

	// #For each entity permission
	permissions, err := t.ListPermission(entity.PermissionSubset{Source: tok.ID.String()})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve permissions")
		return err
	}
	ucentity := uce.E{
		EntityMapper:     t.EntityMapper,
		QRecurrerMapper:  t.QRecurrerMapper,
		QListenerMapper:  t.QListenerMapper,
		PermissionMapper: t.PermissionMapper,
		EntitiesMapper:   t.EntitiesMapper,
	}
	for _, permission := range permissions {
		targetID := ulid.MustParse(permission.Target)
		if err := ucentity.Disconnect(targetID, tok); err != nil {
			logger.Error().Err(err).Str("entity", targetID.String()).Msg("failed to disconnect entity")
		}
	}

	if err := t.DelToken(account.TokenSubset{ID: id}); err != nil {
		logger.Error().Err(err).Str("token", id.String()).Msg("failed to delete token")
		return err
	}

	return nil
}
