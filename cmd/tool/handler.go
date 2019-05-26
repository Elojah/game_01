package main

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/item"
	"github.com/elojah/game_01/pkg/sector"
)

type handler struct {
	ability   ability.App
	account   account.App
	entity    entity.App
	event     event.App
	item      item.App
	sector    sector.App
	sequencer infra.SequencerApp

	srv *http.Server
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ability", h.abilityHandle)
	mux.HandleFunc("/ability/template", h.abilityTemplateHandle)
	mux.HandleFunc("/ability/starter", h.abilityStarterHandle)

	mux.HandleFunc("/entity", h.entityHandle)
	mux.HandleFunc("/entity/template", h.entityTemplateHandle)
	mux.HandleFunc("/entity/move", h.entityMoveHandle)

	mux.HandleFunc("/inventory", h.inventoryHandle)
	mux.HandleFunc("/item", h.itemHandle)

	mux.HandleFunc("/loot", h.lootHandle)

	mux.HandleFunc("/sector", h.sectorHandle)
	mux.HandleFunc("/sector/entities", h.sectorEntitiesHandle)

	mux.HandleFunc("/sequencer", h.sequencerHandle)

	mux.HandleFunc("/spawn", h.spawnHandle)

	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() {
		if err := h.srv.ListenAndServeTLS(c.Cert, c.Key); err != nil {
			log.Error().Err(err).Msg("failed to start server")
		}
	}()
	return nil
}

// Close shutdowns the server listening.
func (h *handler) Close() error {
	return h.srv.Shutdown(context.Background())
}
