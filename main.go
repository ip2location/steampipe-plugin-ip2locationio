package main

import (
	"github.com/ip2location/steampipe-plugin-ip2locationio/ip2locationio"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: ip2locationio.Plugin})
}
