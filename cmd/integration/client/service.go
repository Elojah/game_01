package client

import "io"

// Service wraps client helpers.
type Service struct {
	in io.WriteCloser
}

// NewService returns a integration service for client.
func NewService(in io.WriteCloser) *Service {
	return &Service{in: in}
}
