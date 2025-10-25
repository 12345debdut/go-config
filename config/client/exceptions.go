package client

type ConfigKeyMissingError struct {
	key    string
	config string
}

type ConfigValueMissingError struct {
	key string
}

func (e *ConfigValueMissingError) Error() string {
	return "config value " + e.key + " is missing"
}

func (e *ConfigKeyMissingError) Error() string {
	return "config key " + e.key + " is missing"
}
