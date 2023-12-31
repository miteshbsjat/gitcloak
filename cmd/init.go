/*
Copyright © 2023 Mitesh Singh Jat <mitesh.singh.jat@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/miteshbsjat/gitcloak/pkg/encrypt"
	"github.com/miteshbsjat/gitcloak/pkg/fs"
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
		Info("gitcloak init started")
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
		lineRandom, err := cmd.Flags().GetBool("line-random")
		CheckIfError(err)
		// fmt.Println(lineRandom)
		encSeed, err := cmd.Flags().GetInt64("seed")
		CheckIfError(err)
		if encSeed == encrypt.SEED_DEFAULT {
			encSeed, err = getEncryptionSeed()
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
		// Create the git commit hook script
		err = createGitCommitHook()
		CheckIfError2(err, "Failed to create git pre-commit hook %v", "")

		// Create gitcloak
		gitcloak.GitCloakGitInit()
		gitCloakConfig := gitcloak.NewGitCloakConfig("gcinit", encAlgo, encKey, encSeed, regex, path, lineRandom)
		_, err = gitCloakConfig.CreateGitCloakConfig()
		CheckIfError(err)
		// commit config file
		gcCommitHash, err := gitcloak.GitCloakGitCommit("gitcloak init commit")
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
		_, err = gitcloak.GitCloakGitCommit("gitcloak init mapped commit hashes")
		CheckIfError(err)

		// Add .gitcloak in .gitignore
		gitBase, err := git.GetGitBaseDir()
		CheckIfError(err)
		fs.AddLineToFile(gitBase+"/.gitignore", ".gitcloak")

		Info("gitcloak init completed with %s algorithm", encAlgo)
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
	initCmd.Flags().BoolP("line-random", "l", false, "randomize IV every line in file")
	initCmd.Flags().StringP("path", "p", "",
		"Relative File path for encryption; if -r is given then it is preferred.")
	initCmd.Flags().StringP("regex", "r", "",
		"Regex Pattern for files for encryption like: \".*secret.txt$\"")
	initCmd.Flags().Int64P("seed", "s", encrypt.SEED_DEFAULT,
		"Seed Integer for IV randomization")
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

func getEncryptionSeed() (int64, error) {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("encryption seed must have more than 6 digits")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Encryption Seed for IV",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, err
	}

	fmt.Printf("Your password is %q\n", result)
	seed, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		Warn("Prompt failed %v\n", err)
		return -1, err
	}
	return seed, nil
}

// func confirmInit() {
// 	{
// 		prompt := promptui.Prompt{
// 			Label:     "Init Confirm",
// 			IsConfirm: true,
// 		}

// 		result, err := prompt.Run()

// 		if err != nil {
// 			fmt.Printf("Prompt failed %v\n", err)
// 			return
// 		}

// 		fmt.Printf("You choose %q\n", result)
// 	}
// }

func createGitCommitHook() error {
	scriptPath, err := git.GetGitBaseDir()
	if err != nil {
		Warn("Error in getting git base dir in current dir")
		return err
	}
	scriptPath = scriptPath + "/.git/hooks/pre-commit"

	scriptLines := []string{
		"#!/usr/bin/env bash",
		"echo 'gitcloak git pre-commit hook'",
		"",
		"prog_name=\"$(basename $0)\"",
		"",
		"## Check if gitcloak is present",
		"which gitcloak >/dev/null 2>&1",
		"retval=$?",
		"if [ $retval -ne 0 ]; then",
		"  echo \"$prog_name: gitcloak is not installed.\" >&2",
		"  exit 1",
		"fi",
		"",
		"## find filestate.txt file",
		"dn=$(dirname $0)",
		"filestate=$(find $dn/../.. -name \"filestate.txt\" | grep \".gitcloak\")",
		"if [ \"x$filestate\" == \"x\" ]; then",
		"  echo \"Cannot find file .gitcloak/filestate.txt file\" >&2",
		"  exit 1",
		"fi",
		"any_decrypted=$(grep -c \"=decrypted$\" $filestate)",
		"if [ $any_decrypted -gt 0 ]; then",
		"  echo \"Some file matching rule is in decrypted state in .gitcloak/filestate.txt file\" >&2",
		"  grep \"=decrypted$\" $filestate ",
		"  echo \"\"",
		"  echo \"Please run 'gitcloak encrypt' command\"",
		"  exit 1",
		"fi",
		"echo \"All files are in encrypted state\"",
		"",
		"exit 0",
		"",
	}
	if fs.IsPresent(scriptPath) {
		Warn("git hook %s is already present", scriptPath)
		Info("Please add following lines in %s\n%v", scriptPath, strings.Join(scriptLines, LineBreak()))
		return fmt.Errorf("git hook %s is already present", scriptPath)
	}

	err = fs.CreateShellScript(scriptPath, scriptLines)
	if err != nil {
		Warn("Error in creating git hook file %s", scriptPath)
		return err
	}
	Info("Create git hook %s", scriptPath)
	return nil
}
