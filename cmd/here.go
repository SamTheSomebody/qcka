// Copyright © 2025 Sam Muller gamedevsam@pm.me
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// hereCmd represents the here command
var hereCmd = &cobra.Command{
	Use:   "here [\"name\"]",
	Short: "Add a change directory [cd] alias for the current directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value, err := os.Getwd()
		if err != nil {
			fmt.Printf("[ERROR] Couldn't get working directory: %v\n", err)
			os.Exit(1)
		}
		value = fmt.Sprintf("cd %v", value)
		err = aliases.Add("BOOKMARKS", args[0], value)
		if err != nil {
			fmt.Printf("[ERROR] Couldn't add here alias: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(hereCmd)
}
