package api

// Service wraps api helpers.
type Service struct {
	url string
}

// NewService returns a integration service for api.
func NewService(url string) *Service {
	return &Service{url: url}
}
