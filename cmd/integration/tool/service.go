package tool

// Service wraps tool helpers.
type Service struct {
	url string
}

// NewService returns a integration service for tool.
func NewService(url string) *Service {
	return &Service{url: url}
}
