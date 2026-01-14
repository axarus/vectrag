/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type InitConfig struct {
	Name     string
	Port     int
	Database string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new VectraG project",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		projectPrompt := promptui.Prompt{
			Label:   "Project name",
			Default: "my-app",
		}

		projectName, err := projectPrompt.Run()
		if err != nil {
			return err
		}

		portPrompt := promptui.Prompt{
			Label:   "Port",
			Default: "3000",
			Validate: func(input string) error {
				if input == "" {
					return fmt.Errorf("the port cannot be empty")
				}
				return nil
			},
		}
		port, err := portPrompt.Run()
		if err != nil {
			return err
		}
		dbPrompt := promptui.Select{
			Label:     "Database",
			Items:     []string{"PostgreSQL", "MySQL", "SQLite"},
			CursorPos: 2,
		}

		_, database, err := dbPrompt.Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
