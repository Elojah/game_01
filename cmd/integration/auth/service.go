package auth

// Service wraps account helpers.
type Service struct {
	url string
}

func NewService(url string) *Service {
	return &Service{url: url}
}
