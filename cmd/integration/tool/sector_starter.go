package tool

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// AddSectorStarter add sector starters already existing listed in json filename.
func (s *Service) AddSectorStarter(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "add sector starters")
	}
	defer file.Close()
	resp, err := http.Post(s.url+"/sector/starter", "application/json", file)
	if err != nil {
		return errors.Wrap(err, "add sector starters")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "add sector starters")
	}
	return nil
}
