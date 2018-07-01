package main

import (
	"context"
	"net/http"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
)

type handler struct {
	account.TokenMapper
	AccountMapper account.Mapper

	entity.PCLeftMapper
	entity.PCMapper
	entity.TemplateMapper
	EntityMapper entity.Mapper

	event.QListenerMapper
	event.QMapper
	event.QRecurrerMapper

	infra.CoreMapper
	infra.SyncMapper

	entity.PermissionMapper

	sector.EntitiesMapper
	sector.StarterMapper
	SectorMapper sector.Mapper

	srv *http.Server
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", h.login)
	mux.HandleFunc("/subscribe", h.subscribe)
	mux.HandleFunc("/pc/create", h.createPC)
	mux.HandleFunc("/pc/list", h.listPC)
	mux.HandleFunc("/pc/connect", h.connectPC)
	mux.HandleFunc("/pc/disconnect", h.disconnectPC)
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
