/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"os"
	"path/filepath"
	"tasks/internal/taskhandler"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A todo list for all your tasks",
	PersistentPreRun: func(cmd *cobra.Command, args []string)  {
		filename := filepath.Join(taskhandler.TaskFileLocation, taskhandler.TasksFile)
		taskhandler.InitializeFile(filename)
		initTasks = taskhandler.ReadTasksFromDisk(filename)

		if initTasks == nil {
			initTasks = &taskhandler.Tasks{}
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var initTasks *taskhandler.Tasks

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tasks.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}


