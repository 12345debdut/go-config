package client

type ServerConfigClient interface {
	Get(config, key string) (interface{}, error)
	GetInt(config, key string) (int64, error)
	GetFloat(config, key string) (float64, error)
	GetString(config, key string) (string, error)
	GetBool(config, key string) (bool, error)
	GetSlice(config, key string) ([]interface{}, error)
	GetIntSlice(config, key string) ([]int64, error)
	GetFloatSlice(config, key string) ([]float64, error)
	GetStringSlice(config, key string) ([]string, error)
	GetBoolSlice(config, key string) ([]bool, error)
	GetMap(config, key string) (map[string]interface{}, error)
	GetIntMap(config, key string) (map[string]int64, error)
	GetFloatMap(config, key string) (map[string]float64, error)
	GetStringMap(config, key string) (map[string]string, error)
	GetBoolMap(config, key string) (map[string]bool, error)
	// Close is used to perform any closing actions on the config client
	Close() error
}
