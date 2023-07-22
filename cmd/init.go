/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup and initialize the gitcloak in current git repository.",
	Long: `Setup and initialize the gitcloak in current git repository.
* It creates .gitcloak/ directory.
* Fixes .gitignore for gitcloak.
* Creates .gitcloak/config.yaml based on arguments passed.
* Creates commit-version:config-version map textfilekv store.
* Initialize git repo in .gitcloak/ .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
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

	// gitcloak.addLineToFile("/tmp/test.txt", "DEMO_LINE")
}
