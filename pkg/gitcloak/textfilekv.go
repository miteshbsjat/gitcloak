package gitcloak

import (
	"fmt"
	"os"
	"sync"

	tkv "github.com/miteshbsjat/textfilekv"
)

var tkvInitialized sync.Once
var tkvObj *tkv.KeyValueStore

var kvStoreCache = struct {
	mu    sync.RWMutex
	cache map[string]*tkv.KeyValueStore
}{
	cache: make(map[string]*tkv.KeyValueStore),
}

// NewKVStore is a factory method that returns the appropriate KVStore instance based on the filename.
func NewKVStore(file string) (*tkv.KeyValueStore, error) {
	kvStoreCache.mu.RLock()
	if kvs, ok := kvStoreCache.cache[file]; ok {
		kvStoreCache.mu.RUnlock()
		return kvs, nil
	}
	kvStoreCache.mu.RUnlock()

	kvStoreCache.mu.Lock()
	defer kvStoreCache.mu.Unlock()

	filePath := GetGitCloakBase() + string(os.PathSeparator) + file + ".txt"
	// Create a new KVStore only if it doesn't exist in the cache
	switch file {
	case "ggcmap":
		kvs, err := tkv.NewKeyValueStore(filePath)
		if err != nil {
			return nil, err
		}
		kvStoreCache.cache[file] = kvs
		return kvs, nil
	case "filestate":
		kvs, err := tkv.NewKeyValueStore(filePath)
		if err != nil {
			return nil, err
		}
		kvStoreCache.cache[file] = kvs
		return kvs, nil
	// Add more cases for different filenames if needed
	default:
		return nil, fmt.Errorf("unsupported filename: %s", file)
	}
}

func initTextFileKV() {
	tkvInitialized.Do(func() {
		tkvObj, _ = tkv.NewKeyValueStore(GetGitCloakBase() + "/ggcmap.txt")
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
