package storage

import (
	"net"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Domain converts a storage token into a domain token.
func (t *Token) Domain() (account.Token, error) {
	var token account.Token
	var err error

	token.ID = ulid.ID(t.ID)
	if token.IP, err = net.ResolveUDPAddr("udp", t.IP); err != nil {
		return token, err
	}
	token.Account = ulid.ID(t.Account)
	token.Ping = t.Ping
	token.CorePool = ulid.ID(t.CorePool)
	token.SyncPool = ulid.ID(t.SyncPool)
	token.PC = ulid.ID(t.PC)
	token.Entity = ulid.ID(t.Entity)
	return token, nil
}

// NewToken converts a domain token into a storage token.
func NewToken(token account.Token) *Token {
	return &Token{
		ID:       [16]byte(token.ID),
		IP:       token.IP.String(),
		Account:  [16]byte(token.Account),
		Ping:     token.Ping,
		CorePool: [16]byte(token.CorePool),
		SyncPool: [16]byte(token.SyncPool),
		PC:       [16]byte(token.PC),
		Entity:   [16]byte(token.Entity),
	}
}
