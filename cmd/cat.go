// Copyright © 2025 Sam Muller gamedevsam@pm.me
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "List all aliases in '.bash_aliases'",
	Run: func(cmd *cobra.Command, args []string) {
		i := 0
		if group == "" {
			for groupKey, groupValue := range aliases.Groups {
				i = printGroup(i, groupKey, groupValue.Values)
			}
		} else {
			g, ok := aliases.Groups[group]
			if !ok {
				fmt.Printf("Group '%v' does not exist", group)
			}
			printGroup(i, group, g.Values)
		}
	},
}

func printGroup(i int, groupName string, values map[string]string) int {
	fmt.Printf("-- %v --\n", groupName)
	for key, value := range values {
		fmt.Printf("[%v] %v=\"%v\"\n", i, key, value)
		i++
	}
	fmt.Println()
	return i
}

func init() {
	rootCmd.AddCommand(catCmd)
	catCmd.SuggestFor = []string{"list", "show"}
}
