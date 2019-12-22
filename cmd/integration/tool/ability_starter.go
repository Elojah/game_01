package tool

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// AddAbilityStarter add ability starters listed in json filename.
func (s *Service) AddAbilityStarter(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "add ability starters")
	}
	defer file.Close()
	resp, err := http.Post(s.url+"/ability/starter", "application/json", file)
	if err != nil {
		return errors.Wrap(err, "add ability starters")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "add ability starters")
	}
	_ = resp.Body.Close()
	return nil
}
