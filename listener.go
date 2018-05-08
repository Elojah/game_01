package game

// Listener requires the receiver to create a new listener with subject ID.
type Listener struct {
	ID ID
}

// ListenerService handles send/receive methods for listeners.
type ListenerService interface {
	SendListener(Listener, ID) error
	ReceiveListener(string, int) (Subscription, error)
}
