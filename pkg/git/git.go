package git

import (
	"fmt"
	"os"
	"path/filepath"
)

var gitBaseDir string = ""

func GetGitBaseDir() (string, error) {
	if gitBaseDir == "" {
		gitBaseD, err := findGitRootDir()
		gitBaseDir = gitBaseD
		if err != nil {
			return gitBaseDir, nil
		}
	}
	return gitBaseDir, nil
}

func findGitRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, ".git")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("git repository not found in the current or parent directories")
		}
		currentDir = parentDir
	}
}
