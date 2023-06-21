package ip2locationio

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableIP2LocationIOWhois(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ip2locationio_whois",
		Description: "Whois information for the domain name.",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain", Require: plugin.Required, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromQual("domain").NullIfZero(), Description: "Domain name to query."},
			{Name: "domain_id", Type: proto.ColumnType_STRING, Description: "Domain name ID."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Domain name status."},
			{Name: "create_date", Type: proto.ColumnType_TIMESTAMP, Description: "Domain name creation date."},
			{Name: "update_date", Type: proto.ColumnType_TIMESTAMP, Description: "Domain name updated date."},
			{Name: "expire_date", Type: proto.ColumnType_TIMESTAMP, Description: "Domain name expiration date."},
			{Name: "domain_age", Type: proto.ColumnType_INT, Description: "Domain name age in day(s)."},
			{Name: "whois_server", Type: proto.ColumnType_STRING, Description: "WHOIS server name."},

			// Nested columns
			{Name: "registrar", Type: proto.ColumnType_JSON, Description: "Registrar details."},
			{Name: "registrant", Type: proto.ColumnType_JSON, Description: "Registrant details."},
			{Name: "admin", Type: proto.ColumnType_JSON, Description: "Admin details."},
			{Name: "tech", Type: proto.ColumnType_JSON, Description: "Tech details."},
			{Name: "billing", Type: proto.ColumnType_JSON, Description: "Billing details."},
			{Name: "nameservers", Type: proto.ColumnType_JSON, Description: "Name servers."},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connectWhois(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ip2locationio_whois.listDomain", "whois_connection_error", err)
		return nil, err
	}

	domain := d.EqualsQuals["domain"].GetStringValue()
	plugin.Logger(ctx).Debug("ip2locationio_whois.listDomain", "domain", domain)

	res, err := conn.LookUp(domain)

	plugin.Logger(ctx).Debug("ip2locationio_whois.listDomain", "results", res)
	if err != nil {
		plugin.Logger(ctx).Error("ip2locationio_whois.listDomain", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, res)

	return nil, nil
}
