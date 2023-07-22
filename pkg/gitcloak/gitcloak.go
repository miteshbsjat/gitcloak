package gitcloak

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	mgit "github.com/miteshbsjat/gitcloak/pkg/git"
	. "github.com/miteshbsjat/gitcloak/pkg/utils"
)

type GitCloakConfig struct {
	EncryptionAlgorithm string `yaml:"encryption_algorithm"`
	EncryptionKey       string `yaml:"encryption_key"`
	Path                string `yaml:"path,omitempty"`
	Regex               string `yaml:"path_regex,omitempty"`
}

func gitInit() (*git.Repository, error) {
	repo, err := git.PlainInit(".", false)
	if err != nil {
		return nil, err
	}
	return repo, nil

	// wt, err := repo.Worktree()
	// if err != nil {
	// 	return err
	// }

	// _, err = wt.Commit("Initial commit", &git.CommitOptions{
	// 	Author: &object.Signature{
	// 		Name:  "Mitesh Singh Jat",
	// 		Email: "mitesh.singh.jat@gmail.com",
	// 	},
	// })

	// if err != nil {
	// 	return err
	// }
}

var GITCLOAK_BASE = ""

func GetGitCloakBase() string {
	if GITCLOAK_BASE == "" {
		dirPath, err := mgit.GetGitBaseDir()
		CheckIfError(err)
		dirPath = dirPath + "/.gitcloak"
		GITCLOAK_BASE = dirPath
	}
	return GITCLOAK_BASE
}

func gitCloakGitInit() {

	dirPath := GetGitCloakBase()
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create the directory if it does not exist
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	} else {
		// Do nothing already initialized
		return
	}

	cwd, err := os.Getwd()
	CheckIfError(err)
	defer os.Chdir(cwd)

	// Change the working directory to the specified directory
	if err := os.Chdir(dirPath); err != nil {
		log.Fatal(err)
	}

	// Initialize git repository in the current directory
	_, err = gitInit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Git repository initialized successfully.")
}
