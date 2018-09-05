package main

import (
	"encoding/json"
	"fmt"
)

func expectUp(a *LogAnalyzer) error {
	logup := []common{
		common{Level: "info", Exe: "./bin/game_sync", Message: "sync up"},
		common{Level: "info", Exe: "./bin/game_core", Message: "core up"},
		common{Level: "info", Exe: "./bin/game_api", Message: "api up"},
		common{Level: "info", Exe: "./bin/game_auth", Message: "auth up"},
		common{Level: "info", Exe: "./bin/game_revoker", Message: "revoker up"},
		common{Level: "info", Exe: "./bin/game_tool", Message: "tool up"},
	}

	fup := func(s string) (bool, error) {
		var lup common
		if err := json.Unmarshal([]byte(s), &lup); err != nil {
			return false, err
		}
		found := false
		var i int
		for i = 0; i < len(logup); i++ {
			if lup == logup[i] {
				found = true
				break
			}
		}
		if !found {
			return false, fmt.Errorf("unexpected log %s", s)
		}
		logup[i] = logup[len(logup)-1]
		logup = logup[:len(logup)-1]
		if len(logup) == 0 {
			return true, nil
		}
		return false, nil
	}

	return a.Expect(fup)
}

func expectUpClient(a *LogAnalyzer) error {

	logup := common{Level: "info", Exe: "./bin/game_client", Message: "client up"}

	fup := func(s string) (bool, error) {
		var lup common
		if err := json.Unmarshal([]byte(s), &lup); err != nil {
			return false, err
		}
		if lup != logup {
			return true, fmt.Errorf("unexpected log %s", s)
		}
		return true, nil
	}

	return a.Expect(fup)
}
