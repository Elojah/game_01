package storage

import (
	"net"

	"github.com/elojah/game_01"
)

// Domain converts a storage token into a domain token.
func (t *Token) Domain(id game.ID) (game.Token, error) {
	var token game.Token
	var err error

	token.ID = id
	token.Account = game.ID(t.Account)
	if token.IP, err = net.ResolveUDPAddr("udp", t.IP); err != nil {
		return token, nil
	}
	return token, nil
}

// NewToken converts a domain token into a storage token.
func NewToken(token game.Token) *Token {
	return &Token{
		IP:      token.IP.String(),
		Account: [16]byte(token.Account),
	}
}
