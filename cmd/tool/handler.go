package main

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/item"
	"github.com/elojah/game_01/pkg/sector"
)

type handler struct {
	AbilityStore         ability.Store
	AbilityStarterStore  ability.StarterStore
	AbilityTemplateStore ability.TemplateStore

	AccountStore account.Store

	EntityStore          entity.Store
	EntityTemplateStore  entity.TemplateStore
	EntityInventoryStore entity.InventoryStore
	EntitySpawnStore     entity.SpawnStore

	ItemStore     item.Store
	ItemLootStore item.LootStore

	InfraSequencerService infra.SequencerService

	SectorStore         sector.Store
	SectorEntitiesStore sector.EntitiesStore

	srv *http.Server
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ability", h.ability)
	mux.HandleFunc("/ability/template", h.abilityTemplate)
	mux.HandleFunc("/ability/starter", h.abilityStarter)

	mux.HandleFunc("/entity", h.entity)
	mux.HandleFunc("/entity/template", h.entityTemplate)
	mux.HandleFunc("/entity/move", h.entityMove)

	mux.HandleFunc("/inventory", h.inventory)
	mux.HandleFunc("/item", h.item)

	mux.HandleFunc("/loot", h.loot)

	mux.HandleFunc("/sector", h.sector)
	mux.HandleFunc("/sector/entities", h.sectorEntities)

	mux.HandleFunc("/sequencer", h.sequencer)

	mux.HandleFunc("/spawn", h.spawn)

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
