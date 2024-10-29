/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string)  {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.PrintErrln("Invalid task ID:", err)
		}
		if id <= 0 || id > len(*initTasks) {
			cmd.PrintErrln("Invalid task ID")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.PrintErrln("Invalid task ID:", err)
		}
		initTasks.Complete(id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
