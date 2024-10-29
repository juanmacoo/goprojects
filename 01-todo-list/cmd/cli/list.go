/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cli

import (
	"github.com/spf13/cobra"
)

var allFlag bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		initTasks.List(allFlag)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Show all tasks, including non-completed.")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
