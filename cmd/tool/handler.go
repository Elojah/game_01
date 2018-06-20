package main

import (
	"context"
	"net/http"

	"github.com/elojah/game_01"
)

type handler struct {
	game.AbilityMapper
	game.AbilityTemplateMapper
	game.AccountMapper
	game.EntityMapper
	game.EntityTemplateMapper
	game.SectorMapper
	game.SectorEntitiesMapper

	srv *http.Server
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ability", h.ability)
	mux.HandleFunc("/entity", h.entity)
	mux.HandleFunc("/sector", h.sector)

	mux.HandleFunc("/ability/template", h.abilityTemplate)
	mux.HandleFunc("/entity/template", h.entityTemplate)

	mux.HandleFunc("/sector/entities", h.sectorEntities)

	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() { _ = h.srv.ListenAndServeTLS(c.Cert, c.Key) }()
	return nil
}

// Close shutdowns the server listening.
func (h *handler) Close() error {
	return h.srv.Shutdown(context.Background())
}
