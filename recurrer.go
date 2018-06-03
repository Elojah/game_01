package game

// Recurrer requires the receiver to create a new recurrer with subject ID.
type Recurrer struct {
	ID ID
}

// QRecurrerMapper handles send/receive methods for recurrers.
type QRecurrerMapper interface {
	SendRecurrer(Recurrer, ID) error
}
