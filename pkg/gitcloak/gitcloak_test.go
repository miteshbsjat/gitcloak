package gitcloak

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYamlConfigDump(t *testing.T) {

	encr := Encryption{
		Algorithm: "aes",
		Key:       "abc",
	}

	rule := Rule{
		Name:       "initrule",
		Encryption: encr,
		Regex:      "abc.txt",
	}
	rules := []Rule{rule}
	gcc := GitCloakConfig{
		Rules: rules,
	}

	// Open the file for writing
	fileName := "/tmp/testgcc.yaml"
	file, err := os.Create(fileName)
	if err != nil {
		// log.Error(err)
		t.Error(err)
	}
	defer file.Close()

	// Create a YAML encoder
	encoder := yaml.NewEncoder(file)

	// Encode the struct into YAML and write it to the file
	if err := encoder.Encode(&gcc); err != nil {
		// log.Fatal(err)
		t.Error(err)
	}

	file1, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
	}
	defer file1.Close()

	// Create a YAML decoder
	decoder := yaml.NewDecoder(file1)

	// Create a struct instance to hold the data
	var gcc1 GitCloakConfig

	// Decode the YAML data into the struct
	if err := decoder.Decode(&gcc1); err != nil {
		t.Error(err)
	}

	var data string = gcc.Rules[0].Encryption.Algorithm
	result := gcc1.Rules[0].Encryption.Algorithm
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
	data = gcc.Rules[0].Encryption.Key
	result = gcc1.Rules[0].Encryption.Key
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
	data = gcc.Rules[0].Path
	result = gcc1.Rules[0].Path
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
	data = gcc.Rules[0].Regex
	result = gcc1.Rules[0].Regex
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
}

func TestGitCloakInit(t *testing.T) {
	gcdir := GetGitCloakBase()
	os.Remove(gcdir)
	GitCloakGitInit()
	dirEntry, err := os.ReadDir(gcdir)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(gcdir)

	dotGitPresent := false
	for _, file := range dirEntry {
		if file.Name() == ".git" {
			dotGitPresent = true
			break
		}
	}
	if !dotGitPresent {
		t.Errorf("%s/.git is not present", gcdir)
	}

	GitCloakGitInit()
}
