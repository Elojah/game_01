package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

func (h *handler) createPC(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/create").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var setPC dto.SetPC
	if err := json.NewDecoder(r.Body).Decode(&setPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	if err := setPC.Check(); err != nil {
		logger.Error().Err(err).Msg("name invalid")
		http.Error(w, "name invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("token", setPC.Token.String()).Logger()

	// #Get and check token.
	tok, err := h.T.Get(setPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Check user permission to create a new PC.
	left, err := h.GetPCLeft(entity.PCLeftSubset{AccountID: tok.Account})
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

	// #Decrease token permission to create a new PC by 1.
	if err := h.SetPCLeft(left-1, tok.Account); err != nil {
		logger.Error().Err(err).Msg("failed to decrease left pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("template", setPC.Type.String()).Logger()

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
		logger.Error().Err(err).Msg("failed to pick random starter")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger = logger.With().Str("sector", start.SectorID.String()).Logger()
	sec, err := h.SectorMapper.GetSector(sector.Subset{ID: start.SectorID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve starter sector")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Create PC from the template and put it in a random starter sector.
	pc := entity.PC(template)
	pc.Type = pc.ID
	pc.ID = ulid.NewID()
	pc.Name = setPC.Name
	logger = logger.With().Str("pc", pc.ID.String()).Logger()
	pc.Position = entity.Position{
		SectorID: sec.ID,
		Coord:    geometry.Vec3{X: sec.Size.X * rand.Float64(), Y: sec.Size.Y * rand.Float64(), Z: sec.Size.Z * rand.Float64()},
	}
	if err := pc.Check(); err != nil {
		logger.Error().Err(err).Msg("wrong pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.SetPC(pc, tok.Account); err != nil {
		logger.Error().Err(err).Msg("failed to create pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info().Msg("pc creation success")

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

	logger := log.With().Str("route", "/pc/list").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var listPC dto.ListPC
	if err := json.NewDecoder(r.Body).Decode(&listPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("token", listPC.Token.String()).Logger()

	// #Get and check token.
	tok, err := h.T.Get(listPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("account", tok.Account.String()).Logger()

	// #Retrieve PCs by account.
	pcs, err := h.ListPC(entity.PCSubset{AccountID: tok.Account})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve PCs")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Marshal results.
	raw, err := json.Marshal(pcs)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal PCs")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info().Msg("pc list success")

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

	logger := log.With().Str("route", "/pc/connect").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var connectPC dto.ConnectPC
	if err := json.NewDecoder(r.Body).Decode(&connectPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("token", connectPC.Token.String()).Logger()
	logger = logger.With().Str("pc", connectPC.Target.String()).Logger()

	// #Get and check token.
	tok, err := h.T.Get(connectPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if tok.Entity.Time() != 0 {
		logger.Error().Msg("packet rejected")
		http.Error(w, "token already in use", http.StatusBadRequest)
		return
	}

	// #Retrieve PC for this account.
	pc, err := h.GetPC(entity.PCSubset{
		AccountID: tok.Account,
		ID:        connectPC.Target,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve PC")
		http.Error(w, "failed to connect", http.StatusBadRequest)
		return
	}

	// #Creates entity cloned from pc.
	e := entity.E(pc)
	e.ID = ulid.NewID()
	logger = logger.With().Str("entity", e.ID.String()).Logger()
	if err := h.EntityMapper.SetEntity(e, time.Now().UnixNano()); err != nil {
		logger.Error().Err(err).Msg("failed to create entity from PC")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	logger = logger.With().Str("sector", pc.Position.SectorID.String()).Logger()
	// #Add entity to PC sector.
	if err := h.AddEntityToSector(e.ID, pc.Position.SectorID); err != nil {
		logger.Error().Err(err).Msg("failed to add entity to sector")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Add permission token/entity.
	if err := h.SetPermission(entity.Permission{
		ID:     ulid.NewID(),
		Source: tok.ID.String(),
		Target: e.ID.String(),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create permissions")
		http.Error(w, "failed to create permissions", http.StatusInternalServerError)
		return
	}

	// #Creates a new listener for this entity.
	listener, err := h.L.New(e.ID)
	logger = logger.With().Str("listener", listener.ID.String()).Logger()
	if err != nil {
		logger.Error().Err(err).Msg("failed to create entity listener")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
	}

	// #Creates a new recurrer for this token/entity.
	recurrer, err := h.R.New(e.ID, tok.ID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create entity recurrer")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
	}

	// #Update token with pool informations.
	tok.CorePool = listener.Pool
	tok.SyncPool = recurrer.Pool
	tok.PC = pc.ID
	tok.Entity = e.ID
	if err := h.SetToken(tok); err != nil {
		logger.Error().Err(err).Str("token", tok.ID.String()).Msg("failed to update token pools")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Marshal response
	raw, err := json.Marshal(dto.EntityPC{ID: e.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal response")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	logger.Info().Msg("signin success")

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

// disconnectPC disconnects a PC.
func (h *handler) disconnectPC(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/disconnect").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var disconnectPC dto.DisconnectPC
	if err := json.NewDecoder(r.Body).Decode(&disconnectPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("token", disconnectPC.Token.String()).Logger()

	// #Get and check token.
	tok, err := h.T.Get(disconnectPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.T.Disconnect(tok.ID); err != nil {
		logger.Error().Err(err).Str("token", tok.ID.String()).Msg("failed to disconnect")
		http.Error(w, "failed to disconnect", http.StatusInternalServerError)
		return
	}

	logger.Info().Msg("disconnect success")

	// #Write response
	w.WriteHeader(http.StatusOK)
}
