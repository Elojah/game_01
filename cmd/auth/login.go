package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

func (h handler) login(w http.ResponseWriter, r *http.Request) {
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
	var accountPayload game.Account
	if err = json.Unmarshal(b, &accountPayload); err != nil {
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// Search account in redis
	account, err := h.GetAccount(game.AccountBuilder{
		Username: accountPayload.Username,
	})
	if err != nil && err != storage.ErrNotFound {
		http.Error(w, "failed to retrieve account", http.StatusInternalServerError)
		return
	}
	if err == storage.ErrNotFound || account.Password != accountPayload.Password {
		http.Error(w, "wrong username/password", http.StatusUnauthorized)
		return
	}

	// Identify origin IP
	ip, err := net.ResolveUDPAddr("tcp", r.RemoteAddr)
	if err != nil {
		http.Error(w, "failed to identify ip", http.StatusInternalServerError)
		return
	}

	// Create a new token
	token := game.Token{
		ID:      game.NewULID(),
		Account: account.ID,
		IP:      ip,
	}
	if err := h.CreateToken(token); err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	// Marshal token for response
	raw, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
