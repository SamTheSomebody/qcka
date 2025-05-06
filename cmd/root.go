// Copyright © 2025 Sam Muller gamedevsam@pm.me
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/samthesomebody/qcka/files"
	"github.com/samthesomebody/qcka/settings"
)

var (
	group   string
	aliases *files.Aliases
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qcka",
	Short: "Quick alias is a tool for managaing aliases in bash.",
	Long: `Quick alias allows you to add and manage aliases
	without having to edit your .bashrc or .bash_aliases.`,
}

func Execute() {
	isAliasSourced, err := files.IsAliasSourced()
	if err != nil {
		fmt.Printf("[ERROR] Couldn't check if alias was sourced: %v\n", err)
		os.Exit(1)
	}

	if !isAliasSourced {
		err = files.SourceAlias()
		if err != nil {
			fmt.Printf("[ERROR] Couldn't source aliases: %v\n", err)
			os.Exit(1)
		}
	}

	aliases, err = files.GetAliases()
	if err != nil {
		fmt.Printf("[ERROR] Couldn't read aliases: %v\n", err)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	settings.New()
	rootCmd.PersistentFlags().StringVarP(&group, "group", "g", "", "Define a group for the command")
	rootCmd.PersistentFlags().BoolVarP(&settings.Settings.Force, "force", "f", false, "Ignore any warnings")
	rootCmd.PersistentFlags().BoolVarP(&settings.Settings.Quiet, "quiet", "q", false, "Run in quiet mode")
	rootCmd.PersistentFlags().BoolVarP(&settings.Settings.Verbose, "verbose", "v", false, "Run in verbose mode")
	rootCmd.MarkFlagsMutuallyExclusive("quiet", "verbose")
}
