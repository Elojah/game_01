package main

import (
	"context"
	"net/http"

	"github.com/elojah/game_01/pkg/account"
	accountapp "github.com/elojah/game_01/pkg/account/app"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	infraapp "github.com/elojah/game_01/pkg/infra/app"
	"github.com/elojah/game_01/pkg/sector"
)

type handler struct {
	AccountStore account.Store

	entity.PCStore
	entity.PCLeftStore
	entity.PermissionStore
	entity.TemplateStore

	event.QStore

	infra.SyncStore

	sector.EntitiesStore
	sector.StarterStore
	SectorStore sector.Store

	accountapp.TokenService
	infraapp.ListenerService
	infraapp.RecurrerService

	srv *http.Server
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/signin", h.signin)
	mux.HandleFunc("/signout", h.signout)

	mux.HandleFunc("/subscribe", h.subscribe)
	mux.HandleFunc("/unsubscribe", h.unsubscribe)

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
