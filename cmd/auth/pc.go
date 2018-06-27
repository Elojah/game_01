package main

import (
	"encoding/json"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/oklog/ulid"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/dto"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
)

func (h *handler) createPC(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/create").Logger()

	// # Read body
	var setPC dto.SetPC
	if err := json.NewDecoder(r.Body).Decode(&setPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// # Search message UUID in storage.
	token, err := h.GetToken(setPC.Token)
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("token", setPC.Token.String()).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// # Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(r.RemoteAddr)
	if expected != actual {
		err := game.ErrWrongIP
		logger.Error().Err(err).Str("status", "hijacked").Str("expected", expected).Str("actual", actual).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Check user permission to create h new PC.
	left, err := h.GetPCLeft(entity.PCLeftSubset{AccountID: token.Account})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve left pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if left <= 0 {
		err := game.ErrInvalidAction
		logger.Error().Err(err).Msg("no more pc left")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Decrease token permission to create h new PC by 1.
	if err := h.SetPCLeft(left-1, token.Account); err != nil {
		logger.Error().Err(err).Msg("failed to decrease left pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Retrieve template for new PC.
	template, err := h.GetEntityTemplate(entity.TemplateSubset{Type: setPC.Type})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve template")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Create PC from the template.
	pc := entity.PC(template)
	pc.ID = ulid.NewID()
	pc.Position = entity.Position{
		// TODO list of positions config ? Areas config + random ? Define spawn
		SectorID: ulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
		Coord:    geometry.Vec3{X: 100 * rand.Float64(), Y: 100 * rand.Float64(), Z: 100 * rand.Float64()},
	}
	if err := pc.Check(); err != nil {
		logger.Error().Err(err).Msg("wrong pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.SetPC(pc, token.Account); err != nil {
		logger.Error().Err(err).Msg("failed to create pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
}

func (h *handler) listPC(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/list").Logger()

	// # Read body
	var listPC dto.ListPC
	if err := json.NewDecoder(r.Body).Decode(&listPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// # Search message UUID in storage.
	token, err := h.GetToken(listPC.Token)
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("token", listPC.Token.String()).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// # Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(r.RemoteAddr)
	if expected != actual {
		err := game.ErrWrongIP
		logger.Error().Err(err).Str("status", "hijacked").Str("expected", expected).Str("actual", actual).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Retrieve PCs by account.
	pcs, err := h.ListPC(entity.PCSubset{AccountID: token.Account})
	if err != nil {
		logger.Error().Err(err).Str("account", token.Account.String()).Msg("failed to retrieve PCs")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Marshal results.
	raw, err := json.Marshal(pcs)
	if err != nil {
		logger.Error().Err(err).Str("account", token.Account.String()).Msg("failed to marshal PCs")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

// ConnectPC creates an entity from h PC.
func (h *handler) connectPC(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/list").Logger()

	// # Read body
	var connectPC dto.ConnectPC
	if err := json.NewDecoder(r.Body).Decode(&connectPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// # Search message UUID in storage.
	token, err := h.GetToken(connectPC.Token)
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("token", connectPC.Token.String()).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// # Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(r.RemoteAddr)
	if expected != actual {
		err := game.ErrWrongIP
		logger.Error().Err(err).Str("status", "hijacked").Str("expected", expected).Str("actual", actual).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Retrieve PC for this account.
	pc, err := h.GetPC(entity.PCSubset{
		AccountID: token.Account,
		ID:        connectPC.Target,
	})
	if err != nil {
		logger.Error().Err(err).Str("account", token.Account.String()).Str("id", connectPC.Target.String()).Msg("failed to retrieve PC")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Creates entity cloned from pc.
	entity := entity.E(pc)
	entity.ID = ulid.NewID()
	if err := h.EntityMapper.SetEntity(entity, time.Now().UnixNano()); err != nil {
		logger.Error().Err(err).Str("id", entity.ID.String()).Msg("failed to create entity from PC")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Add entity to PC sector.
	if err := h.AddEntityToSector(entity.ID, pc.Position.SectorID); err != nil {
		logger.Error().Err(err).Str("id", entity.ID.String()).Str("sector", pc.Position.SectorID.String()).Msg("failed to add entity to sector")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Creates h new listener for this entity.
	core := h.cores[rand.Intn(len(h.cores))]
	listener := event.Listener{ID: entity.ID}
	if err := h.SendListener(listener, core); err != nil {
		logger.Error().Err(err).Str("core", core.String()).Str("id", entity.ID.String()).Msg("failed to add listener to entity")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Creates h new synchronizer for this token/entity.
	sync := h.syncs[rand.Intn(len(h.syncs))]
	if err := h.SendRecurrer(event.Recurrer{
		ID:       ulid.NewID(),
		EntityID: entity.ID,
		TokenID:  token.ID,
		Action:   event.OpenRec,
	}, sync); err != nil {
		logger.Error().Err(err).Str("sync", sync.String()).Str("id", entity.ID.String()).Msg("failed to add sync for entity")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Marshal response
	raw, err := json.Marshal(dto.Entity{ID: entity.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
