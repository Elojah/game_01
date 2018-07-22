package dto

// ACK represents a ACK value sent from server to client to assert a received packet/input.
type ACK struct {
	ID [16]byte
}
