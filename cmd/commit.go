/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/miteshbsjat/gitcloak/pkg/git"
	"github.com/miteshbsjat/gitcloak/pkg/gitcloak"
	. "github.com/miteshbsjat/gitcloak/pkg/utils"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Info("gitcloak commit called")
		commitMessage, err := cmd.Flags().GetString("message")
		CheckIfError(err)
		if commitMessage == "" {
			commitMessage, err = getCommitMessage()
			CheckIfError(err)
		}

		// commit config file
		gcCommitHash, err := gitcloak.GitCloakGitCommit("gitcloak commit " + commitMessage)
		CheckIfError(err)
		Info("gitcloak Commit Hash = %s", gcCommitHash)
		pwd, err := os.Getwd()
		CheckIfError(err)
		gCommitHash, err := git.GetGitCommitHash(pwd)
		CheckIfError(err)
		Info("git Commit Hash = %s", gCommitHash)
		gckv, err := gitcloak.NewKVStore("ggcmap")
		CheckIfError(err)
		err = gckv.Set(gCommitHash, gcCommitHash)
		CheckIfError(err)
		_, err = gitcloak.GitCloakGitCommit("gitcloak commit mapped commit hashes")
		CheckIfError(err)
		Info("gitcloak commit completed with %s message", commitMessage)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	commitCmd.Flags().StringP("message", "m", "", "Message for gitcloak commit")
}

func getCommitMessage() (string, error) {
	validate := func(input string) error {
		if len(input) < 10 {
			return errors.New("Commit Message should have more than 10 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Commit Message",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
