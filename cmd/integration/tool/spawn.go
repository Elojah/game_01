package tool

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// AddSpawn add spawns listed in json filename.
func (s *Service) AddSpawn(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "add spawns")
	}
	defer file.Close()
	resp, err := http.Post(s.url+"/spawn", "application/json", file)
	if err != nil {
		return errors.Wrap(err, "add spawns")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "add spawn starters")
	}
	return nil
}
