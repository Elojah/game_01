package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/elojah/game_01"
)

func (h handler) subscribe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// Unmarshal payload
	var account game.Account
	if err = json.Unmarshal(b, &account); err != nil {
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	account.ID = game.NewULID()

	// Create account in redis
	if err = h.CreateAccount(account); err != nil {
		http.Error(w, "failed to create account", http.StatusInternalServerError)
		return
	}

	// Marshal token for response
	raw, err := json.Marshal(account)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
