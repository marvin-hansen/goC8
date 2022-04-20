package c8

type ClientConfig struct {
	Fabric           string
	Host             string
	Timeout          int
	connectionString string
}

func NewClientDefaultConfig() *ClientConfig {
	host := defaultEndpoint
	fabric := defaultFabric
	return &ClientConfig{
		Fabric:           fabric,
		Host:             host,
		Timeout:          5,
		connectionString: host + fabric,
	}
}

func NewClientConfig(host, fabric string, timeout int) *ClientConfig {
	return &ClientConfig{
		Host:             host,
		Fabric:           fabric,
		Timeout:          timeout,
		connectionString: host + fabric,
	}
}

func (c ClientConfig) GetConnectionString() string {
	return c.connectionString
}
