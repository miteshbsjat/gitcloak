package gitcloak

import (
	tkv "github.com/miteshbsjat/textfilekv"
)

var tkvInitialized bool
var tkvObj *tkv.KeyValueStore

func initTextFileKV() {
	if !tkvInitialized {
		tkvInitialized = true
		tkvObj = tkv.NewKeyValueStore(GetGitCloakBase() + "/ggcmap.txt")
	}
}

func GetTextFileKV() *tkv.KeyValueStore {
	initTextFileKV()
	return tkvObj
}

func PutGitAndGitCloak(gitCommitHash string, gitCloakCommitHash string) {
	initTextFileKV()
	tkvObj.Set(gitCommitHash, gitCloakCommitHash)
}

func GetGitCloakCommitHash(gitCommitHash string) (string, bool) {
	initTextFileKV()
	return tkvObj.Get(gitCommitHash)
}
