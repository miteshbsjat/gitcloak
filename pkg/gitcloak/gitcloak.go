package gitcloak

type GitCloakConfig struct {
	EncryptionAlgorithm string `yaml:"encryption_algorithm"`
	EncryptionKey       string `yaml:"encryption_key"`
	Path                string `yaml:"path,omitempty"`
	Regex               string `yaml:"path_regex,omitempty"`
}
