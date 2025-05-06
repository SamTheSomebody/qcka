// Copyright © 2025 Sam Muller gamedevsam@pm.me
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an alias",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := aliases.Remove(args[0])
		if err != nil {
			fmt.Printf("[ERROR] Couldn't remove %v: %v", args[0], err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.SuggestFor = []string{"rm", "delete"}
}
