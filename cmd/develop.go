/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// developCmd represents the develop command
var developCmd = &cobra.Command{
	Use:   "develop",
	Short: "Run VectraG in development mode",
	Long: `Run VectraG in development mode.

The develop command starts VectraG with development-friendly settings and tools, allowing you to:

- Create, update, and delete content models
- Experiment with schema changes safely
- Iterate quickly during local development

This mode is intended for local environments and early-stage development.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(developCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// developCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// developCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
