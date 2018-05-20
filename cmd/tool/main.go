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

type template struct {
	game.EntityTemplateService

	config    string
	templates string
}

// run template tool.
func (t *template) run(cmd *cobra.Command, args []string) {

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	launchers := services.Launchers{}

	// redis
	rd := redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers = append(launchers, rdl)
	rdx := redisx.NewService(&rd)

	t.EntityTemplateService = rdx

	if err := launchers.Up(t.config); err != nil {
		logger.Error().Err(err).Str("filename", t.config).Msg("failed to start")
		return
	}

	raw, err := ioutil.ReadFile(t.templates)
	if err != nil {
		logger.Error().Err(err).Str("templates", t.templates).Msg("failed to read templates file")
		return
	}
	var templates []game.EntityTemplate
	if err := json.Unmarshal(raw, &templates); err != nil {
		logger.Error().Err(err).Str("templates", t.templates).Msg("invalid JSON")
		return
	}

	logger.Info().Int("templates", len(templates)).Msg("found")

	for _, tpl := range templates {
		if err := t.SetEntityTemplate(tpl); err != nil {
			logger.Error().Err(err).Str("template", tpl.ID.String()).Msg("failed to set template")
			return
		}
	}

	logger.Info().Msg("done")
}

func main() {
	var t template

	var root = &cobra.Command{
		Use:   "game_tool [template]",
		Short: "game_tool for data processing",
		Long:  "game_tool provides multiple commands/utils/helpers for data processing and ops requirement.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	var templateCmd = &cobra.Command{
		Use:   "template [no options!]",
		Short: "add new entity templates",
		Long: `template creates new templates entities from a JSON file. e.g:
			[
			{
				ID       : "01CDSTJRVK0HMG6TREBJR7FG1N",
				HP       : 142,
				MP       : 142,
				Position : {},
				Type     : 2
			}
			]
		`,
		Run: t.run,
	}
	templateCmd.Flags().StringVar(&t.config, "config", "", "config file for DB connections")
	templateCmd.MarkFlagRequired("config")
	templateCmd.Flags().StringVar(&t.templates, "templates", "", "templates where templates are represented in JSON")
	templateCmd.MarkFlagRequired("templates")

	root.AddCommand(templateCmd)
	root.Execute()
}
