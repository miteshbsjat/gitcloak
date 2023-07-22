/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitcloak",
	Short: "Securing sensitive file(s) in git diff compatible manner.",
	Long: `This tool secures the sensitive file(s) in your git repository.
	
This tool achives this in a simple and innovative manner of encrypting given file line by line.
This helps in making diff work while remaining compatible to git diff.
There are several usecases where this tool can be used.
  * Obsidian: markdown based note taking app (Second Brain),
  * Public Password store without taking any third party paid service,
  * All files related to a service in one repository, hence easier and simpler CI/CD.
  * Please add more if you can think of more usecases.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitcloak.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
