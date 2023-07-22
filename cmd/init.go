/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	mfs "github.com/miteshbsjat/gitcloak/pkg/fs"

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

	mfs.AddLineToFile("/tmp/test.txt", "DEMO_LINE")
	encAlgo, _ := getEncryptionAlgo()
	fmt.Println(encAlgo)
	// password, _ := getPassword()
	// fmt.Println(password)
	confirmInit()
}

// func initPrompt() {

// }

func getEncryptionAlgo() (string, error) {
	prompt := promptui.Select{
		Label: "Select Encryption Algorithm",
		Items: []string{"aes", "chacha", "xxtea"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "aes", err
	}

	fmt.Printf("You choose %q\n", result)
	return result, nil
}

func getPassword() (string, error) {
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
