package main

import (
	"context"
	"net/http"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
)

type handler struct {
	AbilityStore         ability.Store
	AbilityTemplateStore ability.TemplateStore

	AccountStore account.Store

	EntityStore         entity.Store
	EntityTemplateStore entity.TemplateStore

	infra.SequencerService

	SectorStore sector.Store
	sector.EntitiesStore
	sector.StarterStore

	srv *http.Server
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ability/template", h.abilityTemplate)
	mux.HandleFunc("/entity/template", h.entityTemplate)

	mux.HandleFunc("/sector/entities", h.sectorEntities)
	mux.HandleFunc("/sector/starter", h.sectorStarter)

	mux.HandleFunc("/ability", h.ability)
	mux.HandleFunc("/entity", h.entity)
	mux.HandleFunc("/sector", h.sector)

	mux.HandleFunc("/sequencer", h.sequencer)

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
