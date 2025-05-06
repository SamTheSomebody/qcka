/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/samthesomebody/qcka/files"
	"github.com/spf13/cobra"
)

var addGroup string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [value]",
	Short: "Add an alias to '.bash_aliases'",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		//TODO check if alias is valid
		files.Write()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addGroup, "group", "g", "", "Add the alias to a group")
}
