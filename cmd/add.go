// Copyright © 2025 Sam Muller gamedevsam@pm.me
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [\"name\"] [\"value\"]",
	Short: "Add an alias to '.bash_aliases'",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if group == "" {
			group = "DEFAULT"
		}

		err := aliases.Add(group, args[0], args[1])
		if err != nil {
			fmt.Printf("[ERROR] Couldn't add alias: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
