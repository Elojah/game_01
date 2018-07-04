package main

import (
	"context"
	"net/http"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
)

type handler struct {
	AbilityMapper         ability.Mapper
	AbilityTemplateMapper ability.TemplateMapper

	AccountMapper account.Mapper

	EntityMapper         entity.Mapper
	EntityTemplateMapper entity.TemplateMapper

	event.ListenerMapper
	event.QListenerMapper

	SectorMapper sector.Mapper
	sector.EntitiesMapper
	sector.StarterMapper

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
