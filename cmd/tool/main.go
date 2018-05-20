package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/elojah/game_01"
)

func main() {
	var t template

	var root = &cobra.Command{
		Use:   "game_tool [template]",
		Short: "game_tool for data processing",
		Long:  "game_tool provides multiple commands/utils/helpers for data processing and ops requirement.",
		Args:  cobra.MinimumNArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	var templateCmd = &cobra.Command{
		Use:   "template --skills=skills.json --entities=entities.json",
		Short: "add new entity/skill templates",
		Long: `template creates new templates from JSON files. e.g:
			entities:
			[{
				"hp"       : 142,
				"mp"       : 142,
				"type"     : "01CDSTJRVK0HMG6TREBJR7FG1N",
				"name"     : "scavenger"
			}]

			skills:
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
		Run: t.run,
	}
	templateCmd.Flags().StringVar(&t.config, "config", "", "config file for DB connections")
	templateCmd.MarkFlagRequired("config")
	templateCmd.Flags().StringVar(&t.entities, "entities", "", "file where entities are represented in JSON")
	templateCmd.Flags().StringVar(&t.skills, "skills", "", "file where skills are represented in JSON")

	var idCmd = &cobra.Command{
		Use:   "id [no options!]",
		Short: "returns a valid ULID",
		Long:  `id returns a the string representation of a valid ULID set at current timestamp.`,
		Run:   func(cmd *cobra.Command, args []string) { fmt.Println(game.NewULID().String()) },
	}

	root.AddCommand(templateCmd)
	root.AddCommand(idCmd)
	root.Execute()
}
