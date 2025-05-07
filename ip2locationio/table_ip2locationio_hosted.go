package ip2locationio

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"net"
)

func tableIP2LocationIOHosted(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ip2locationio_hosted",
		Description: "Hosted domains information for the IP address.",
		List: &plugin.ListConfig{
			Hydrate: listHosted,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "ip", Require: plugin.Required, CacheMatch: "exact"},
				{Name: "page", Require: plugin.Optional, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_INET, Description: "IP address to query."},
			{Name: "total_domains", Type: proto.ColumnType_INT, Description: "Total hosted domains for that IP address."},
			{Name: "page", Type: proto.ColumnType_INT, Description: "Page of the result."},
			{Name: "per_page", Type: proto.ColumnType_INT, Description: "Number of domains to display per page."},
			{Name: "total_pages", Type: proto.ColumnType_STRING, Description: "Total pages of the result."},
			{Name: "domains", Type: proto.ColumnType_JSON, Description: "The list of hosted domains."},
		},
	}
}

func listHosted(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connectHosted(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ip2locationio_hosted.listHosted", "hosted_connection_error", err)
		return nil, err
	}

	ipStr := ""
	if d.EqualsQuals["ip"] != nil {
		ipStr = d.EqualsQuals["ip"].GetInetValue().GetAddr()
		plugin.Logger(ctx).Debug("ip2locationio_hosted.listHosted", "ipStr", ipStr)
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		plugin.Logger(ctx).Warn("ip2locationio_hosted.listHosted", "invalid_ip", ip, "status", "Must supply a valid IP address.")
	}

	page := 1
	if d.EqualsQuals["page"] != nil {
		page = int(d.EqualsQuals["page"].GetInt64Value())
		plugin.Logger(ctx).Debug("ip2locationio_hosted.listHosted", "page", page)
	}

	res, err := conn.LookUp(ip.String(), page)

	plugin.Logger(ctx).Debug("ip2locationio_hosted.listHosted", "results", res)
	if err != nil {
		plugin.Logger(ctx).Error("ip2locationio_hosted.listHosted", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, res)

	return nil, nil
}
