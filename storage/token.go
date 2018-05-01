package storage

import (
	"net"

	"github.com/elojah/game_01"
)

// Domain converts a storage token into a domain token.
func (t *Token) Domain() (game.Token, error) {
	var token game.Token
	var err error

	token.ID = game.ID(t.ID)
	token.Account = game.ID(t.Account)
	if token.IP, err = net.ResolveUDPAddr("udp", t.IP); err != nil {
		return token, nil
	}
	return token, nil
}
