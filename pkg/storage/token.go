package storage

import (
	"net"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Domain converts a storage token into a domain token.
func (t *Token) Domain() (account.Token, error) {
	var tok account.Token
	var err error

	tok.ID = ulid.ID(t.ID)
	if tok.IP, err = net.ResolveUDPAddr("udp", t.IP); err != nil {
		return tok, err
	}
	tok.Account = ulid.ID(t.Account)
	tok.Ping = t.Ping
	tok.CorePool = ulid.ID(t.CorePool)
	tok.SyncPool = ulid.ID(t.SyncPool)
	tok.PC = ulid.ID(t.PC)
	tok.Entity = ulid.ID(t.Entity)
	return tok, nil
}

// NewToken converts a domain token into a storage token.
func NewToken(tok account.Token) *Token {
	return &Token{
		ID:       [16]byte(tok.ID),
		IP:       tok.IP.String(),
		Account:  [16]byte(tok.Account),
		Ping:     tok.Ping,
		CorePool: [16]byte(tok.CorePool),
		SyncPool: [16]byte(tok.SyncPool),
		PC:       [16]byte(tok.PC),
		Entity:   [16]byte(tok.Entity),
	}
}
