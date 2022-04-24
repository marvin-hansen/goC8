package goC8

const (
	defaultEndpoint = "https://api-bigmouth-2265268a-ap-south.paas.macrometa.io/_fabric/"
	defaultFabric   = "SouthEastAsia"
	apiKEY          = KEY
)

type ClientConfig struct {
	apiKey           string
	Fabric           string
	Timeout          int
	QueryTTL         int
	connectionString string
}

func NewDefaultConfig() *ClientConfig {
	host := defaultEndpoint
	fabric := defaultFabric
	return &ClientConfig{
		apiKey:           apiKEY,
		Fabric:           fabric,
		Timeout:          5,
		connectionString: host + fabric,
	}
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
