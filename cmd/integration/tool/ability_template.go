package tool

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// AddAbilityTemplates add ability templates listed in json filename.
func (s *Service) AddAbilityTemplates(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "add ability templates")
	}
	defer file.Close()
	resp, err := http.Post(s.url+"/ability/template", "application/json", file)
	if err != nil {
		return errors.Wrap(err, "add ability templates")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "add ability templates")
	}
	return nil
}
