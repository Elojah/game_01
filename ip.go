package game

import (
	"net"
)

// IP represents a client/server address.
type IP *net.UDPAddr

// IPService handles IP interactions.
type IPService interface {
	CheckIP(IP, ID) error
}
