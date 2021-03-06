package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/oklog/ulid"
	perrors "github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// SetPC represents the payload to send to create a new PC.
type SetPC struct {
	Token gulid.ID
	Name  string
	Type  gulid.ID
	Spawn gulid.ID
}

// Check checks setpc validity.
func (spc SetPC) Check() error {
	l := len(spc.Name)
	if l < 4 || l > 24 || strings.IndexFunc(spc.Name, func(r rune) bool {
		return (r < 'A' || r > 'z') && (r < '0' || r > '9') && (r != '_')
	}) != -1 {
		return errors.New("invalid name")
	}
	return nil
}

// ListPC represents the payload to list token PCs.
type ListPC struct {
	Token gulid.ID
}

// DelPC represents the payload to delete a PC.
type DelPC struct {
	Token gulid.ID
	PC    gulid.ID
}

// ConnectPC represents the payload to connect to an existing PC.
type ConnectPC struct {
	Token  gulid.ID
	Target gulid.ID
}

// DisconnectPC represents the payload to disconnect a token.
type DisconnectPC struct {
	Token gulid.ID
}

// EntityPC represents the response when connecting to an existing PC.
type EntityPC struct {
	ID gulid.ID
}

func (h *handler) createPC(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/create").Str("method", "POST").Str("address", r.RemoteAddr).Logger()

	// #Read body
	var setPC SetPC
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
	tok, err := h.account.FetchTokenFromAddr(setPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Check user permission to create a new PC.
	left, err := h.entity.FetchPCLeft(tok.Account)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve left pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if left <= 0 {
		err := gerrors.ErrFullPCCreated{AccountID: tok.Account.String()}
		logger.Error().Err(err).Msg("no more pc left")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Decrease token permission to create a new PC by 1.
	if err := h.entity.UpsertPCLeft(left-1, tok.Account); err != nil {
		logger.Error().Err(err).Msg("failed to decrease left pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("template", setPC.Type.String()).Logger()

	// #Retrieve entity template for new PC.
	template, err := h.entity.FetchTemplate(setPC.Type)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve template")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Create PC from the template.
	pc := template
	pc.Type = pc.ID
	pc.ID = gulid.NewID()
	logger = logger.With().Str("pc", pc.ID.String()).Logger()
	if err := pc.Check(); err != nil {
		logger.Error().Err(err).Msg("wrong pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pc.Name = setPC.Name

	// #Set starter abilities to pc.
	if err := h.ability.SetStarters(pc.ID, pc.Type); err != nil {
		logger.Error().Err(err).Str("pc", pc.ID.String()).Msg("failed to set starter abilities")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Set inventory to PC.
	pc.InventoryID = gulid.NewID()
	// set inventory in mr only, will be written in lru at connection
	if err := h.entity.UpsertMRInventory(pc.ID, entity.Inventory{
		ID:    pc.InventoryID,
		Size_: entity.DefaultInventorySize,
		Items: make(map[string]uint64),
	}); err != nil {
		logger.Error().Err(err).Str("pc", pc.ID.String()).Msg("failed to set inventory")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// #Retrieve pc spawn.
	pc.SpawnID = setPC.Spawn
	logger = logger.With().Str("spawn", pc.SpawnID.String()).Logger()
	sp, err := h.entity.FetchSpawn(pc.SpawnID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve spawn")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Assign pc position at spawn and set pc.
	pc.Position = sp.Position
	if err := h.entity.UpsertPC(pc, tok.Account); err != nil {
		logger.Error().Err(err).Msg("failed to create pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("pc creation success")
}

func (h *handler) listPC(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/list").Str("method", "POST").Str("address", r.RemoteAddr).Logger()

	// #Read body
	var listPC ListPC
	if err := json.NewDecoder(r.Body).Decode(&listPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("token", listPC.Token.String()).Logger()

	// #Get and check token.
	tok, err := h.account.FetchTokenFromAddr(listPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("account", tok.Account.String()).Logger()

	// #Retrieve PCs by account.
	pcs, err := h.entity.ListPC(tok.Account)
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

	// #Write response
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		return
	}

	logger.Info().Msg("pc list success")
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

	logger := log.With().Str("route", "/pc/connect").Str("method", "POST").Str("address", r.RemoteAddr).Logger()

	// #Read body
	var connectPC ConnectPC
	if err := json.NewDecoder(r.Body).Decode(&connectPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().
		Str("token", connectPC.Token.String()).
		Str("pc", connectPC.Target.String()).
		Logger()

	// #Get and check token.
	tok, err := h.account.FetchTokenFromAddr(connectPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !tok.Entity.IsZero() {
		logger.Error().Msg("packet rejected")
		http.Error(w, "token already in use", http.StatusBadRequest)
		return
	}

	// #Retrieve PC for this account.
	pc, err := h.entity.FetchPC(tok.Account, connectPC.Target)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve PC")
		http.Error(w, "failed to connect", http.StatusBadRequest)
		return
	}

	// #Creates entity cloned from pc.
	e := pc
	e.ID = gulid.NewID()
	logger = logger.With().Str("entity", e.ID.String()).Logger()
	if err := h.entity.Upsert(e, ulid.Now()); err != nil {
		logger.Error().Err(err).Msg("failed to create entity from PC")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Copy PC abilities
	if err := h.ability.Copy(pc.ID, e.ID); err != nil {
		logger.Error().Err(err).Msg("failed to copy abilities to entity")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Copy inventory and add inventory in lru
	inv, err := h.entity.FetchMRInventory(pc.ID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve inventory from pc")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}
	if err := h.entity.UpsertMRInventoryWithCache(e.ID, inv); err != nil {
		logger.Error().Err(err).Msg("failed to set inventory to entity")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	logger = logger.With().Str("sector", pc.Position.SectorID.String()).Logger()
	// #Add entity to PC sector.
	if err := h.sector.AddEntityToSector(e.ID, pc.Position.SectorID); err != nil {
		logger.Error().Err(err).Msg("failed to add entity to sector")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Add permission token/entity.
	if err := h.entity.UpsertPermission(entity.Permission{
		ID:     gulid.NewID(),
		Source: tok.ID.String(),
		Target: e.ID.String(),
		Value:  int(account.Owner),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create permissions")
		http.Error(w, "failed to create permissions", http.StatusInternalServerError)
		return
	}

	// #Creates a new sequencer for this entity.
	sequencer, err := h.sequencer.Create(e.ID)
	logger = logger.With().Str("sequencer", sequencer.ID.String()).Logger()
	if err != nil {
		logger.Error().Err(err).Msg("failed to create entity sequencer")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
	}

	// #Creates a new recurrer for this token/entity.
	recurrer, err := h.recurrer.Create(e.ID, tok.ID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create entity recurrer")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
	}

	// #Update token with pool informations.
	tok.CorePool = sequencer.Pool
	tok.SyncPool = recurrer.Pool
	tok.PC = pc.ID
	tok.Entity = e.ID
	if err := h.account.UpsertToken(tok); err != nil {
		logger.Error().Err(err).Str("token", tok.ID.String()).Msg("failed to update token pools")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Marshal response
	raw, err := json.Marshal(EntityPC{ID: e.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal response")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		return
	}

	logger.Info().Msg("connect pc success")
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

	logger := log.With().Str("route", "/pc/disconnect").Str("address", r.RemoteAddr).Logger()

	// #Read body
	var disconnectPC DisconnectPC
	if err := json.NewDecoder(r.Body).Decode(&disconnectPC); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("token", disconnectPC.Token.String()).Logger()

	// #Get and check token.
	tok, err := h.account.FetchTokenFromAddr(disconnectPC.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Disconnect token entities.
	if err := h.account.DisconnectToken(tok.ID); err != nil {
		logger.Error().Err(err).Str("token", tok.ID.String()).Msg("failed to disconnect")
		http.Error(w, "failed to disconnect", http.StatusInternalServerError)
		return
	}

	// #Close token recurrer.
	if err := h.recurrer.Erase(tok.ID); err != nil {
		logger.Error().Err(err).Str("token", tok.ID.String()).Msg("failed to remove recurrer")
		http.Error(w, "failed to remove recurrer", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("disconnect success")
}

// delPC deletes a PC.
func (h *handler) delPC(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/del").Str("address", r.RemoteAddr).Logger()

	// #Read body
	var del DelPC
	if err := json.NewDecoder(r.Body).Decode(&del); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("token", del.Token.String()).Logger()

	// #Get and check token.
	tok, err := h.account.FetchTokenFromAddr(del.Token, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Disconnect token entities.
	// It is not possible to be connected when deleting a PC to avoid a resave (would lead to a wrong pcleft+1).
	if err := h.account.DisconnectToken(tok.ID); err != nil {
		logger.Error().Err(err).Msg("failed to disconnect token")
		http.Error(w, "failed to disconnect token", http.StatusInternalServerError)
		return
	}

	// #Close token recurrer.
	if err := h.recurrer.Erase(tok.ID); err != nil {
		switch perrors.Cause(err).(type) {
		case gerrors.ErrNotFound:
		default:
			logger.Error().Err(err).Msg("failed to remove recurrer")
			http.Error(w, "failed to remove recurrer", http.StatusInternalServerError)
			return
		}
	}

	// #Remove PC through service
	if err := h.entity.ErasePC(tok.Account, del.PC); err != nil {
		logger.Error().Err(err).Msg("failed to remove PC")
		http.Error(w, "failed to remove PC", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("del pc success")
}
