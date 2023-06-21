package ip2locationio

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type ip2locationioConfig struct {
	ApiKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &ip2locationioConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) ip2locationioConfig {
	if connection == nil || connection.Config == nil {
		return ip2locationioConfig{}
	}
	config, _ := connection.Config.(ip2locationioConfig)
	return config
}
