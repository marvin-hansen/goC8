package config

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/acc"
)

const (
	defaultEndpoint = "https://api-bigmouth-2265268a-ap-south.paas.macrometa.io/_fabric/"
	defaultFabric   = "SouthEastAsia"
	apiKEY          = acc.KEY
	timeout         = 5 // http connection timeout in seconds
)

func GetDefaultConfig() *goC8.ClientConfig {
	return goC8.NewConfig(apiKEY, defaultEndpoint, defaultFabric, timeout)
}
