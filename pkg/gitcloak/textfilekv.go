package gitcloak

import (
	"sync"

	tkv "github.com/miteshbsjat/textfilekv"
)

var tkvInitialized sync.Once
var tkvObj *tkv.KeyValueStore

func initTextFileKV() {
	tkvInitialized.Do(func() {
		tkvObj = tkv.NewKeyValueStore(GetGitCloakBase() + "/ggcmap.txt")
	})
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
