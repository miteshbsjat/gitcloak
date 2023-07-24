/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/miteshbsjat/gitcloak/pkg/encrypt"
	"github.com/miteshbsjat/gitcloak/pkg/git"
	"github.com/miteshbsjat/gitcloak/pkg/gitcloak"
	. "github.com/miteshbsjat/gitcloak/pkg/utils"

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
		encAlgo, err := cmd.Flags().GetString("encryption-algorithm")
		CheckIfError(err)
		if encAlgo == "" {
			encAlgo, err = getEncryptionAlgo()
			CheckIfError(err)
		}
		// fmt.Println(encAlgo)

		encKey, err := cmd.Flags().GetString("key")
		CheckIfError(err)
		if encKey == "" {
			encKey, err = getEncryptionKey()
			CheckIfError(err)
		}
		// fmt.Println(encKey)

		regex, err := cmd.Flags().GetString("regex")
		CheckIfError(err)
		// fmt.Println(regex)
		path := ""
		if regex == "" {
			path1, err := cmd.Flags().GetString("path")
			CheckIfError(err)
			path = path1
			// fmt.Println(path)
		}

		// Create gitcloak
		gitcloak.GitCloakGitInit()
		gitCloakConfig := gitcloak.GitCloakConfig{
			EncryptionAlgorithm: encAlgo,
			EncryptionKey:       encKey,
			Path:                path,
			Regex:               regex,
		}
		_, err = gitCloakConfig.CreateGitCloakConfig()
		CheckIfError(err)
		// commit config file
		gcCommitHash, err := gitcloak.GitCloakGitCommit("gitcloak init commit")
		CheckIfError(err)
		fmt.Println(gcCommitHash)
		pwd, err := os.Getwd()
		CheckIfError(err)
		gCommitHash, err := git.GetGitCommitHash(pwd)
		CheckIfError(err)
		fmt.Println(gCommitHash)
		// gitcloak.PutGitAndGitCloak(gCommitHash, gcCommitHash)
		gckv, err := gitcloak.NewKVStore("ggcmap")
		CheckIfError(err)
		err = gckv.Set(gCommitHash, gcCommitHash)
		CheckIfError(err)
		_, err = gitcloak.GitCloakGitCommit("gitcloak init mapped commit hashes")
		CheckIfError(err)
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
	initCmd.Flags().StringP("encryption-algorithm", "e", "",
		"Encryption Algorithm to select from "+strings.Join(encrypt.ENCRYPTION_ALGORITHMS, ", "))
	initCmd.Flags().StringP("key", "k", "", "Encryption Key 16 characters")
	initCmd.Flags().StringP("path", "p", "",
		"Relative File path for encryption; if -r is given then it is preferred.")
	initCmd.Flags().StringP("regex", "r", "",
		"Regex Pattern for files for encryption like: \"*secret.txt\"")
}

// func initPrompt() {

// }

func getEncryptionAlgo() (string, error) {
	prompt := promptui.Select{
		Label: "Select Encryption Algorithm",
		// Items: []string{"aes", "chacha", "xxtea"},
		Items: encrypt.ENCRYPTION_ALGORITHMS,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "aes", err
	}

	fmt.Printf("You choose %q\n", result)
	return result, nil
}

func getEncryptionKey() (string, error) {
	validate := func(input string) error {
		if len(input) < 16 {
			return errors.New("encryption key must have more than 15 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Encryption Key",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	// fmt.Printf("Your password is %q\n", result)
	return result, nil
}

func confirmInit() {
	{
		prompt := promptui.Prompt{
			Label:     "Init Confirm",
			IsConfirm: true,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)
	}
}
