package gitcloak

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYamlConfigDump(t *testing.T) {

	gcc := GitCloakConfig{
		EncryptionAlgorithm: "aes",
		EncryptionKey:       "abc",
		Regex:               "*abc.txt",
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

	var data string = gcc.EncryptionAlgorithm
	result := gcc1.EncryptionAlgorithm
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
	data = gcc.EncryptionKey
	result = gcc1.EncryptionKey
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
	data = gcc.Path
	result = gcc1.Path
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
	data = gcc.Regex
	result = gcc1.Regex
	if result != data {
		t.Errorf("%s ! = %s", result, data)
	}
}
