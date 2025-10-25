package client

import (
	"slices"
	"sync"

	"github.com/spf13/cast"
)

type localServerConfigClient struct {
	config      map[string]struct{}
	configNames []string
	configMap   map[string]*lockedConfigValue
	mutex       *sync.RWMutex
}

func newLocalServerConfigClient() *localServerConfigClient {
	return &localServerConfigClient{
		config:      make(map[string]struct{}),
		configNames: make([]string, 0),
		configMap:   make(map[string]*lockedConfigValue),
	}
}

type lockedConfigValue struct {
	value interface{}
	lock  *sync.RWMutex
}

func (c *localServerConfigClient) Get(config, key string) (interface{}, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	effectiveKey := config + "." + key
	if _, ok := c.config[effectiveKey]; ok {
		return c.configMap[effectiveKey].value, nil
	}
	return nil, &ConfigKeyMissingError{key: key, config: config}
}

func (c *localServerConfigClient) GetInt(config, key string) (int64, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return 0, err
	}
	if configValue == nil {
		return 0, &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToInt64E(configValue)
	if err != nil {
		return 0, err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetFloat(config, key string) (float64, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return 0, err
	}
	if configValue == nil {
		return 0, &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToFloat64E(configValue)
	if err != nil {
		return 0, err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetString(config, key string) (string, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return "", err
	}
	if configValue == nil {
		return "", &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToStringE(configValue)
	if err != nil {
		return "", err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetBool(config, key string) (bool, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return false, err
	}
	if configValue == nil {
		return false, &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToBoolE(configValue)
	if err != nil {
		return false, err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetSlice(config, key string) ([]interface{}, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make([]interface{}, 0), err
	}
	if configValue == nil {
		return make([]interface{}, 0), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToSliceE(configValue)
	if err != nil {
		return make([]interface{}, 0), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetIntSlice(config, key string) ([]int64, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make([]int64, 0), err
	}
	if configValue == nil {
		return make([]int64, 0), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToInt64SliceE(configValue)
	if err != nil {
		return make([]int64, 0), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetFloatSlice(config, key string) ([]float64, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make([]float64, 0), err
	}
	if configValue == nil {
		return make([]float64, 0), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToFloat64SliceE(configValue)
	if err != nil {
		return make([]float64, 0), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetStringSlice(config, key string) ([]string, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make([]string, 0), err
	}
	if configValue == nil {
		return make([]string, 0), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToStringSliceE(configValue)
	if err != nil {
		return make([]string, 0), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetBoolSlice(config, key string) ([]bool, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make([]bool, 0), err
	}
	if configValue == nil {
		return make([]bool, 0), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToBoolSliceE(configValue)
	if err != nil {
		return make([]bool, 0), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetIntMap(config, key string) (map[string]int64, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make(map[string]int64), err
	}
	if configValue == nil {
		return make(map[string]int64), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToStringMapInt64E(configValue)
	if err != nil {
		return make(map[string]int64), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetFloatMap(config, key string) (map[string]float64, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make(map[string]float64), err
	}
	if configValue == nil {
		return make(map[string]float64), &ConfigValueMissingError{}
	}
	configValueFloat, err := toStringMapFloat64E(configValue)
	if err != nil {
		return make(map[string]float64), err
	}
	return configValueFloat, nil
}

func (c *localServerConfigClient) GetStringMap(config, key string) (map[string]string, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make(map[string]string), err
	}
	if configValue == nil {
		return make(map[string]string), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToStringMapStringE(configValue)
	if err != nil {
		return make(map[string]string), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetBoolMap(config, key string) (map[string]bool, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make(map[string]bool), err
	}
	if configValue == nil {
		return make(map[string]bool), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToStringMapBoolE(configValue)
	if err != nil {
		return make(map[string]bool), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) GetMap(config, key string) (map[string]interface{}, error) {
	configValue, err := c.Get(config, key)
	if err != nil {
		return make(map[string]interface{}), err
	}
	if configValue == nil {
		return make(map[string]interface{}), &ConfigValueMissingError{}
	}
	configValueInt, err := cast.ToStringMapE(configValue)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return configValueInt, nil
}

func (c *localServerConfigClient) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k := range c.configMap {
		delete(c.configMap, k)
	}
	for key, _ := range c.config {
		delete(c.config, key)
	}
	slices.Delete(c.configNames, 0, len(c.configNames))
	return nil
}
