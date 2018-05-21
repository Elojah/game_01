package game

// Listener requires the receiver to create a new listener with subject ID.
type Listener struct {
	ID ID
}

// QListenerMapper handles send/receive methods for listeners.
type QListenerMapper interface {
	SendListener(Listener, ID) error
}
