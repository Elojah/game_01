package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/elojah/game_01"
	redisx "github.com/elojah/game_01/storage/redis"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

type skillWithEntity struct {
	game.Skill
	EntityID game.ID `json:"entity_id"`
}

type skill struct {
	game.SkillMapper

	config string
	skills string

	logger zerolog.Logger
}

// run skill tool.
func (s *skill) init() error {

	s.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	launchers := services.Launchers{}

	// redis
	rd := redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers = append(launchers, rdl)
	rdx := redisx.NewService(&rd)

	s.SkillMapper = rdx

	if err := launchers.Up(s.config); err != nil {
		s.logger.Error().Err(err).Str("filename", s.config).Msg("failed to start")
		return err
	}
	return nil
}

func (s *skill) AddSkills(cmd *cobra.Command, args []string) {

	if err := s.init(); err != nil {
		return
	}

	s.CreateSkills()

	s.logger.Info().Msg("done")
}

func (s *skill) CreateSkills() {

	raw, err := ioutil.ReadFile(s.skills)
	if err != nil {
		s.logger.Error().Err(err).Str("skills", s.skills).Msg("failed to read skills file")
		return
	}
	var skills []skillWithEntity
	if err := json.Unmarshal(raw, &skills); err != nil {
		s.logger.Error().Err(err).Str("skills", s.skills).Msg("invalid JSON")
		return
	}

	s.logger.Info().Int("skills", len(skills)).Msg("found")

	for _, sk := range skills {
		if err := s.SetSkill(sk.Skill, sk.EntityID); err != nil {
			s.logger.Error().Err(err).Str("skill", sk.ID.String()).Msg("failed to set skill")
			return
		}
	}
}
