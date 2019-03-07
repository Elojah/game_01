package tool

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// AddSector add sectors listed in json filename.
func (s *Service) AddSector(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "add sectors")
	}
	defer file.Close()
	resp, err := http.Post(s.url+"/sector", "application/json", file)
	if err != nil {
		return errors.Wrap(err, "add sectors")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "add sector starters")
	}
	return nil
}
