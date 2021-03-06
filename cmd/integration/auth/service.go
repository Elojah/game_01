package auth

// Service wraps account helpers.
type Service struct {
	url string
}

// NewService returns a integration service for auth.
func NewService(url string) *Service {
	return &Service{url: url}
}
