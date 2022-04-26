package goC8

type ClientConfig struct {
	apiKey           string
	Fabric           string
	Timeout          int
	QueryTTL         int
	connectionString string
}

func NewConfig(apiKey, endpoint, fabric string, timeout int) *ClientConfig {
	return &ClientConfig{
		apiKey:           apiKey,
		Fabric:           fabric,
		Timeout:          timeout,
		connectionString: endpoint + fabric,
	}
}

func (c ClientConfig) GetConnectionString() string {
	return c.connectionString
}

func (c ClientConfig) GetApiKey() string {
	return c.apiKey
}

func (c ClientConfig) GetQueryTTL() int {
	return c.QueryTTL
}
