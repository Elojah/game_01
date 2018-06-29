package main

import (
	"encoding/json"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/dto"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/perm"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
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

	// #Read body
	var setPC dto.SetPC
	if err := json.NewDecoder(r.Body).Decode(&setPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// #Search message UUID in storage.
	token, err := h.GetToken(account.TokenSubset{ID: setPC.Token})
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("token", setPC.Token.String()).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(r.RemoteAddr)
	if expected != actual {
		err := account.ErrWrongIP
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
		err := account.ErrInvalidAction
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

	// #Retrieve a random starter sector.
	start, err := h.GetRandomStarter(sector.StarterSubset{})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve starter sector")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sec, err := h.SectorMapper.GetSector(sector.Subset{ID: start.SectorID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve starter sector")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Create PC from the template and put it in a random starter sector.
	pc := entity.PC(template)
	pc.ID = ulid.NewID()
	pc.Position = entity.Position{
		SectorID: sec.ID,
		Coord:    geometry.Vec3{X: sec.Size.X * rand.Float64(), Y: sec.Size.Y * rand.Float64(), Z: sec.Size.Z * rand.Float64()},
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

	// #Read body
	var listPC dto.ListPC
	if err := json.NewDecoder(r.Body).Decode(&listPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// #Search message UUID in storage.
	token, err := h.GetToken(account.TokenSubset{ID: listPC.Token})
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("token", listPC.Token.String()).Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(r.RemoteAddr)
	if expected != actual {
		err := account.ErrWrongIP
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

	logger := log.With().Str("route", "/pc/connect").Logger()

	// #Read body
	var connectPC dto.ConnectPC
	if err := json.NewDecoder(r.Body).Decode(&connectPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// #Search message UUID in storage.
	token, err := h.GetToken(account.TokenSubset{ID: connectPC.Token})
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("token", connectPC.Token.String()).Msg("packet rejected")
		http.Error(w, "wrong token id", http.StatusBadRequest)
		return
	}

	// #Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(r.RemoteAddr)
	if expected != actual {
		err := account.ErrWrongIP
		logger.Error().Err(err).Str("status", "hijacked").Str("expected", expected).Str("actual", actual).Msg("packet rejected")
		http.Error(w, "unrecognized ip", http.StatusBadRequest)
		return
	}

	if token.Entity.Time() != 0 {
		logger.Error().Str("entity", token.Entity.String()).Str("pc", token.PC.String()).Msg("packet rejected")
		http.Error(w, "token already in use", http.StatusBadRequest)
		return
	}

	// #Retrieve PC for this account.
	pc, err := h.GetPC(entity.PCSubset{
		AccountID: token.Account,
		ID:        connectPC.Target,
	})
	if err != nil {
		logger.Error().Err(err).Str("account", token.Account.String()).Str("id", connectPC.Target.String()).Msg("failed to retrieve PC")
		http.Error(w, "failed to connect", http.StatusBadRequest)
		return
	}

	// #Creates entity cloned from pc.
	entity := entity.E(pc)
	entity.ID = ulid.NewID()
	if err := h.EntityMapper.SetEntity(entity, time.Now().UnixNano()); err != nil {
		logger.Error().Err(err).Str("id", entity.ID.String()).Msg("failed to create entity from PC")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Add entity to PC sector.
	if err := h.AddEntityToSector(entity.ID, pc.Position.SectorID); err != nil {
		logger.Error().Err(err).Str("id", entity.ID.String()).Str("sector", pc.Position.SectorID.String()).Msg("failed to add entity to sector")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Add permission token/entity.
	if err := h.PermMapper.SetPermission(perm.P{
		ID:     ulid.NewID(),
		Source: token.ID.String(),
		Target: entity.ID.String(),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create permissions")
		http.Error(w, "failed to create permissions", http.StatusInternalServerError)
		return
	}

	// #Creates a new listener for this entity.
	// Set a new listener for this token
	core, err := h.GetRandomCore(infra.CoreSubset{})
	if err != nil {
		if err == storage.ErrNotFound {
			logger.Error().Err(err).Msg("no core available")
			http.Error(w, "failed to create listener", http.StatusInternalServerError)
			return
		}
		logger.Error().Err(err).Msg("failed to get a core")
		http.Error(w, "failed to create listener", http.StatusInternalServerError)
		return
	}
	if err := h.SendListener(event.Listener{ID: entity.ID, Action: event.Open}, core.ID); err != nil {
		logger.Error().Err(err).Str("core", core.ID.String()).Str("id", entity.ID.String()).Msg("failed to add listener to entity")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Creates a new synchronizer for this token/entity.
	// Set a new recurrer for this token/entity.
	sync, err := h.GetRandomSync(infra.SyncSubset{})
	if err != nil {
		if err == storage.ErrNotFound {
			logger.Error().Err(err).Msg("no sync available")
			http.Error(w, "failed to create recurrer", http.StatusInternalServerError)
			return
		}
		logger.Error().Err(err).Msg("failed to get a sync")
		http.Error(w, "failed to create recurrer", http.StatusInternalServerError)
		return
	}
	if err := h.SendRecurrer(event.Recurrer{
		ID:       ulid.NewID(),
		EntityID: entity.ID,
		TokenID:  token.ID,
		Action:   event.Open,
	}, sync.ID); err != nil {
		logger.Error().Err(err).Str("sync", sync.ID.String()).Str("id", entity.ID.String()).Msg("failed to add sync for entity")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Update token with pool informations.
	token.CorePool = core.ID
	token.SyncPool = sync.ID
	token.PC = pc.ID
	token.Entity = entity.ID
	if err := h.SetToken(token); err != nil {
		logger.Error().Err(err).Str("token", token.ID.String()).Msg("failed to update token pools")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Marshal response
	raw, err := json.Marshal(dto.Entity{ID: entity.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal response")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
