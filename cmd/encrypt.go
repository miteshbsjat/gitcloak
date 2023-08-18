/*
Copyright © 2023 Mitesh Singh Jat <mitesh.singh.jat@gmail.com>
*/
package cmd

import (
	"github.com/miteshbsjat/gitcloak/pkg/encrypt"
	"github.com/miteshbsjat/gitcloak/pkg/gitcloak"
	. "github.com/miteshbsjat/gitcloak/pkg/utils"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt the files/files-regex given in .gitcloak/config.yaml",
	Long: `Encrypts the files/files-regex given in .gitcloak/config.yaml

* All set of files given in rules will be encrypted`,
	Run: func(cmd *cobra.Command, args []string) {
		Info("gitcloak encrypt started")

		gitcloakConfigFile, err := cmd.Flags().GetString("configuration")
		CheckIfError(err)
		// Read the given configuration file into struct
		gcc, err := gitcloak.ReadGitCloakConfig(gitcloakConfigFile)
		CheckIfError(err)

		// Loop through each given rules
		for ruleId := range gcc.Rules {
			Info("Processing Rule : %d", ruleId)
			err := encrypt.ProcessRuleForEncryption(gcc.Rules[ruleId])
			CheckIfError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	encryptCmd.Flags().StringP("configuration", "c",
		gitcloak.GetGitCloakConfigPath(), "gitcloak config file")
}
