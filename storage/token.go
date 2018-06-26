package storage

import (
	"net"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/account"
)

// Domain converts a storage token into a domain token.
func (t *Token) Domain(id game.ID) (account.Token, error) {
	var token account.Token
	var err error

	token.ID = id
	token.Account = game.ID(t.Account)
	if token.IP, err = net.ResolveUDPAddr("udp", t.IP); err != nil {
		return token, nil
	}
	return token, nil
}

// NewToken converts a domain token into a storage token.
func NewToken(token account.Token) *Token {
	return &Token{
		IP:      token.IP.String(),
		Account: [16]byte(token.Account),
	}
}
