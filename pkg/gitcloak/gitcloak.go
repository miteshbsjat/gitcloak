package gitcloak

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	mgit "github.com/miteshbsjat/gitcloak/pkg/git"
	. "github.com/miteshbsjat/gitcloak/pkg/utils"
	"gopkg.in/yaml.v3"
)

// type GitCloakConfig struct {
// 	EncryptionAlgorithm string `yaml:"encryption_algorithm"`
// 	EncryptionKey       string `yaml:"encryption_key"`
// 	Path                string `yaml:"path,omitempty"`
// 	Regex               string `yaml:"path_regex,omitempty"`
// }

type Encryption struct {
	Algorithm string `yaml:"algorithm"`
	Key       string `yaml:"key"`
	IV        string `yaml:"iv,omitempty"`
}

type Rule struct {
	Name       string     `yaml:"name"`
	Encryption Encryption `yaml:"encryption"`
	Path       string     `yaml:"path,omitempty"`
	Regex      string     `yaml:"path_regex,omitempty"`
}

type GitCloakConfig struct {
	Rules []Rule `yaml:"rules"`
}

func NewRule(name, encrAlgo, encrKey, regex, path string) *Rule {
	encr := Encryption{
		Algorithm: encrAlgo,
		Key:       encrKey,
	}

	rule := Rule{
		Name:       name,
		Encryption: encr,
		Regex:      regex,
		Path:       path,
	}
	return &rule
}

func NewGitCloakConfig(name, encrAlgo, encrKey, regex, path string) *GitCloakConfig {
	rule := NewRule(name, encrAlgo, encrKey, regex, path)
	rules := []Rule{*rule}
	gcc := GitCloakConfig{
		Rules: rules,
	}
	return &gcc
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

func GetGitCloakConfigPath() string {
	return GetGitCloakBase() + "/config.yaml"
}

func (gcc *GitCloakConfig) CreateGitCloakConfig() (string, error) {
	// Open the file for writing
	fileName := GetGitCloakConfigPath()
	file, err := os.Create(fileName)
	if err != nil {
		return fileName, err
	}
	defer file.Close()

	// Create a YAML encoder
	encoder := yaml.NewEncoder(file)

	// Encode the struct into YAML and write it to the file
	if err := encoder.Encode(&gcc); err != nil {
		return fileName, err
	}
	return fileName, nil
}

func GitCloakGitInit() {

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

func GitCloakGitCommit(commitMessage string) (commitHash string, err error) {
	dirPath := GetGitCloakBase()

	cwd, err := os.Getwd()
	CheckIfError(err)
	defer os.Chdir(cwd)

	// git commit with the given message
	repo, err := git.PlainOpen(dirPath)
	if err != nil {
		return "", err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	// Add all changes to the repository
	_, err = wt.Add(".")
	if err != nil {
		return "", err
	}

	// Commit the changes
	commit, err := wt.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Mitesh Singh Jat",
			Email: fmt.Sprintf("%s@example.com", "mitesh"),
			When:  time.Now(),
		},
	})
	if err != nil {
		return "", err
	}

	// Print the commit hash
	// fmt.Println("Commit Hash:", commit)
	commitHash = commit.String()
	return commitHash, nil
}
