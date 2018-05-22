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
		Use:   "add-template --config=bin/config_core.json --abilitys=abilitys.json --entities=entities.json",
		Short: "add new entity/ability templates",
		Long: `add-template creates new templates from JSON files. e.g:
			entities:
			[{
				"hp"       : 142,
				"mp"       : 142,
				"type"     : "01CDSTJRVK0HMG6TREBJR7FG1N",
				"name"     : "scavenger"
			}]

			abilitys:
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
	addTemplateCmd.Flags().StringVar(&t.abilitys, "abilitys", "", "file where abilitys are represented in JSON")

	var showTemplateCmd = &cobra.Command{
		Use:   "show-template --config=bin/config_core.json [abilitys] [entities]",
		Short: "show entity/ability templates",
		Long:  `show-template list all templates defined in redis namespaces`,
		Args:  cobra.MinimumNArgs(1),
		Run:   t.ShowTemplates,
	}
	showTemplateCmd.Flags().StringVar(&t.config, "config", "", "config file for DB connections")
	showTemplateCmd.MarkFlagRequired("config")

	var abilitysCmd = &cobra.Command{
		Use:   "add-abilitys --config=bin/config_core.json --abilitys=abilitys.json",
		Short: "add new abilitys linked to an entity",
		Long: `add-abilitys creates new abilitys from JSON files. e.g:
			abilitys:
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
		Run: s.AddAbilitys,
	}
	abilitysCmd.Flags().StringVar(&s.config, "config", "", "config file for DB connections")
	abilitysCmd.MarkFlagRequired("config")
	abilitysCmd.Flags().StringVar(&s.abilitys, "abilitys", "", "file where abilitys are represented in JSON")

	var idCmd = &cobra.Command{
		Use:   "id [no options!]",
		Short: "returns a valid ULID",
		Long:  `id returns a the string representation of a valid ULID set at current timestamp.`,
		Run:   func(cmd *cobra.Command, args []string) { fmt.Println(game.NewULID().String()) },
	}

	root.AddCommand(addTemplateCmd)
	root.AddCommand(showTemplateCmd)
	root.AddCommand(abilitysCmd)
	root.AddCommand(idCmd)
	root.Execute()
}
