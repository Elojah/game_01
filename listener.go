package game

// Listener requires the receiver to create a new listener with subject ID.
type Listener struct {
	ID ID
}

// QListenerService handles send/receive methods for listeners.
type QListenerService interface {
	SendListener(Listener, ID) error
}
