package client

// Service wraps client helpers.
type Service struct {
	url string
}

// NewService returns a integration service for client.
func NewService(url string) *Service {
	return &Service{url: url}
}
