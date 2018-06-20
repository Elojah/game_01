package main

import (
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
)

func (h *handler) pcCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/create").Logger()

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// # Unmarshal message.
	msg := dto.Message{}
	if _, err := msg.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// # Parse message UUID.
	tokenID := game.ID(msg.Token)

	// # Search message UUID in storage.
	token, err := h.GetToken(tokenID)
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("tokenID", tokenID.String()).Msg("packet rejected")
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

	spc, ok := msg.Action.(game.SetPC)
	if !ok {
		logger.Error().Err(err).Str("status", "wrongtyped").Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Check user permission to create a new PC.
	left, err := h.GetPCLeft(game.PCLeftSubset{AccountID: token.Account})
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

	// #Decrease token permission to create a new PC by 1.
	if err := h.SetPCLeft(left-1, token.Account); err != nil {
		logger.Error().Err(err).Msg("failed to decrease left pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Retrieve template for new PC.
	template, err := h.GetEntityTemplate(game.EntityTemplateSubset{Type: spc.Type})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve template")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Create PC from the template.
	pc := game.PC(template)
	pc.ID = game.NewID()
	pc.Position = game.Position{
		// TODO list of positions config ? Areas config + random ? Define spawn
		SectorID: ulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
		Coord:    game.Vec3{X: 100 * rand.Float64(), Y: 100 * rand.Float64(), Z: 100 * rand.Float64()},
	}
	if err := pc.Check(); err != nil {
		// #TODO delete pc
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

func (h *handler) pcList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/pc/list").Logger()

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// # Unmarshal message.
	msg := dto.Message{}
	if _, err := msg.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// # Parse message UUID.
	tokenID := game.ID(msg.Token)

	// # Search message UUID in storage.
	token, err := h.GetToken(tokenID)
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("tokenID", tokenID.String()).Msg("packet rejected")
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

	lpc, ok := msg.Action.(game.ListPC)
	if !ok {
		logger.Error().Err(err).Str("status", "wrongtyped").Msg("packet rejected")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Check user permission to create a new PC.
	left, err := h.GetPCLeft(game.PCLeftSubset{AccountID: token.Account})
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

	// #Decrease token permission to create a new PC by 1.
	if err := h.SetPCLeft(left-1, token.Account); err != nil {
		logger.Error().Err(err).Msg("failed to decrease left pc")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Retrieve template for new PC.
	template, err := h.GetEntityTemplate(game.EntityTemplateSubset{Type: lpc.Type})
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve template")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Create PC from the template.
	pc := game.PC(template)
	pc.ID = game.NewID()
	pc.Position = game.Position{
		// TODO list of positions config ? Areas config + random ? Define spawn
		SectorID: ulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
		Coord:    game.Vec3{X: 100 * rand.Float64(), Y: 100 * rand.Float64(), Z: 100 * rand.Float64()},
	}
	if err := pc.Check(); err != nil {
		// #TODO delete pc
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
