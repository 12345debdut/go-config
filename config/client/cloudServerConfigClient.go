package client

type cloudServerConfigClient struct{}

func newCloudServerConfigClient() *cloudServerConfigClient {
	return &cloudServerConfigClient{}
}

func (c *cloudServerConfigClient) Get(config, key string) (interface{}, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetInt(config, key string) (int64, error) {
	return 0, nil
}

func (c *cloudServerConfigClient) GetFloat(config, key string) (float64, error) {
	return 0, nil
}

func (c *cloudServerConfigClient) GetString(config, key string) (string, error) {
	return "", nil
}

func (c *cloudServerConfigClient) GetBool(config, key string) (bool, error) {
	return false, nil
}

func (c *cloudServerConfigClient) GetSlice(config, key string) ([]interface{}, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetIntSlice(config, key string) ([]int64, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetFloatSlice(config, key string) ([]float64, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetStringSlice(config, key string) ([]string, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetBoolSlice(config, key string) ([]bool, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetIntMap(config, key string) (map[string]int64, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetFloatMap(config, key string) (map[string]float64, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetStringMap(config, key string) (map[string]string, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetBoolMap(config, key string) (map[string]bool, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) GetMap(config, key string) (map[string]interface{}, error) {
	return nil, nil
}

func (c *cloudServerConfigClient) Unmarshal(config, key string, value interface{}) error {
	return nil
}

func (c *cloudServerConfigClient) Close() error {
	return nil
}
