package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/elojah/game_01"
)

func main() {
	var t template
	var s ability

	var root = &cobra.Command{
		Use:   "game_tool [add-template] [show-template]",
		Short: "game_tool for data processing",
		Long:  "game_tool provides multiple commands/utils/helpers for data processing and ops requirement.",
		Args:  cobra.MinimumNArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	var addTemplateCmd = &cobra.Command{
		Use:   "add-template --config=bin/config_core.json --abilities=abilities.json --entities=entities.json",
		Short: "add new entity/ability templates",
		Long: `add-template creates new templates from JSON files. e.g:
			entities:
			[{
				"hp"       : 142,
				"mp"       : 142,
				"type"     : "01CDSTJRVK0HMG6TREBJR7FG1N",
				"name"     : "scavenger"
			}]

			abilities:
			[{
				"type"          : "01CDSTJRVK0HMG6TREBJR7FG1N",
				"name"          : "fireball",
				"mp_consumption" : 30,
				"direct_damage"  : 12,
				"direct_heal"    : 0,
				"cd"            : 4,
				"current_cd"     : 0
			}]
		`,
		Run: t.AddTemplates,
	}
	addTemplateCmd.Flags().StringVar(&t.config, "config", "", "config file for DB connections")
	addTemplateCmd.MarkFlagRequired("config")
	addTemplateCmd.Flags().StringVar(&t.entities, "entities", "", "file where entities are represented in JSON")
	addTemplateCmd.Flags().StringVar(&t.abilities, "abilities", "", "file where abilities are represented in JSON")

	var showTemplateCmd = &cobra.Command{
		Use:   "show-template --config=bin/config_core.json [abilities] [entities]",
		Short: "show entity/ability templates",
		Long:  `show-template list all templates defined in redis namespaces`,
		Args:  cobra.MinimumNArgs(1),
		Run:   t.ShowTemplates,
	}
	showTemplateCmd.Flags().StringVar(&t.config, "config", "", "config file for DB connections")
	showTemplateCmd.MarkFlagRequired("config")

	var abilitiesCmd = &cobra.Command{
		Use:   "add-abilities --config=bin/config_core.json --abilities=abilities.json",
		Short: "add new abilities linked to an entity",
		Long: `add-abilities creates new abilities from JSON files. e.g:
			abilities:
			[{
				"entity_id"      : "01CDSTJRVK0HMG6TREBJR7FG1N",
				"type"           : "01CDSTJRVK0HMG6TREBJR7FG1N",
				"name"           : "fireball",
				"mp_consumption" : 30,
				"direct_damage"  : 12,
				"direct_heal"    : 0,
				"cd"             : 4,
				"current_cd"     : 0
			}]
		`,
		Run: s.AddAbilities,
	}
	abilitiesCmd.Flags().StringVar(&s.config, "config", "", "config file for DB connections")
	abilitiesCmd.MarkFlagRequired("config")
	abilitiesCmd.Flags().StringVar(&s.abilities, "abilities", "", "file where abilities are represented in JSON")

	var idCmd = &cobra.Command{
		Use:   "id [no options!]",
		Short: "returns a valid ULID",
		Long:  `id returns a the string representation of a valid ULID set at current timestamp.`,
		Run:   func(cmd *cobra.Command, args []string) { fmt.Println(game.NewULID().String()) },
	}

	root.AddCommand(addTemplateCmd)
	root.AddCommand(showTemplateCmd)
	root.AddCommand(abilitiesCmd)
	root.AddCommand(idCmd)
	root.Execute()
}
