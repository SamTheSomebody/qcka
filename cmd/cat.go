/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/samthesomebody/qcka/files"
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "List all aliases in '.bash_aliases'",
	Run: func(cmd *cobra.Command, args []string) {
		aliases, err := files.GetAliases()
		if err != nil {
			fmt.Printf("[ERROR] unable to get aliases: %v\n", err)
			os.Exit(1)
		}

		i := 0
		for key, value := range(aliases) {
			fmt.Printf("[%v] %v: %v\n", i, key, value)
			i++
		}
	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
