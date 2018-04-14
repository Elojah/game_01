package game

type Right uint8

const (
	Read Right = iota
	Update
	Create
	Delete
)
