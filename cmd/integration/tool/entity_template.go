package tool

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// AddEntityTemplates add entity templates listed in json filename.
func (s *Service) AddEntityTemplates(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "add entity templates")
	}
	defer file.Close()
	resp, err := http.Post(s.url+"/entity/template", "application/json", file)
	if err != nil {
		return errors.Wrap(err, "add entity templates")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "add entity templates")
	}
	return nil
}
