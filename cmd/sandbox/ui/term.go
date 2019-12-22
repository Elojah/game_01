package ui

type Term struct {
}

func (t *Term) Close() error {
	return nil
}

// Dial initialize a Term.
func (t *Term) Dial(cfg Config) error {
	return nil
}
