package main

import (
	"encoding/json"
	"fmt"

	"github.com/elojah/game_01/pkg/entity"
)

func expectState(a *LogAnalyzer, ent entity.E) (entity.E, error) {
	var e entity.E
	err := a.Expect(func(s string) (bool, error) {
		if err := json.Unmarshal([]byte(s), &e); err != nil {
			return false, fmt.Errorf("invalid entity %s", s)
		}
		if ent.ID.Compare(e.ID) != 0 {
			return false, nil
		}
		return true, nil
	})
	return e, err
}
