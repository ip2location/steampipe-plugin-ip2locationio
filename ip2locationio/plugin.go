package ip2locationio

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-ip2locationio",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		// DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"ip2locationio_geolocation": tableIP2LocationIOGeolocation(ctx),
			"ip2locationio_whois":       tableIP2LocationIOWhois(ctx),
			"ip2locationio_hosted":       tableIP2LocationIOHosted(ctx),
		},
	}
	return p
}
