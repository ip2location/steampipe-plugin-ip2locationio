package ip2locationio

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"net"
)

func tableIP2LocationIOGeolocation(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ip2locationio_geolocation",
		Description: "Geolocation information for the IP address.",
		List: &plugin.ListConfig{
			Hydrate: listIP,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "ip", Require: plugin.Required, CacheMatch: "exact"},
				{Name: "lang", Require: plugin.Optional, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_INET, Description: "IP address to query."},
			{Name: "lang", Type: proto.ColumnType_STRING, Transform: transform.FromQual("lang").NullIfZero(), Description: "Translation language (only for Plus and Security plans)."},
			{Name: "country_code", Type: proto.ColumnType_STRING, Description: "ISO-3166 Country code."},
			{Name: "country_name", Type: proto.ColumnType_STRING, Description: "Country name."},
			{Name: "region_name", Type: proto.ColumnType_STRING, Description: "Region name."},
			{Name: "city_name", Type: proto.ColumnType_STRING, Description: "City name."},
			{Name: "latitude", Type: proto.ColumnType_DOUBLE, Description: "Latitude."},
			{Name: "longitude", Type: proto.ColumnType_DOUBLE, Description: "Longitude."},
			{Name: "zip_code", Type: proto.ColumnType_STRING, Description: "Postal code."},
			{Name: "time_zone", Type: proto.ColumnType_STRING, Description: "Time zone."},
			{Name: "asn", Type: proto.ColumnType_STRING, Description: "Autonomous System Number (ASN)."},
			{Name: "as", Type: proto.ColumnType_STRING, Transform: transform.FromField("AS"), Description: "Autonomous System (AS)."}, // does not work without transform
			{Name: "is_proxy", Type: proto.ColumnType_BOOL, Description: "True if IP a proxy."},

			// Other columns
			{Name: "isp", Type: proto.ColumnType_STRING, Description: "Internet Service Provider (ISP)."},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Internet domain name."},
			{Name: "net_speed", Type: proto.ColumnType_STRING, Description: "Internet connection type."},
			{Name: "idd_code", Type: proto.ColumnType_STRING, Description: "IDD code."},
			{Name: "area_code", Type: proto.ColumnType_STRING, Description: "Area code."},
			{Name: "weather_station_code", Type: proto.ColumnType_STRING, Description: "Weather station code."},
			{Name: "weather_station_name", Type: proto.ColumnType_STRING, Description: "Weather station name."},
			{Name: "elevation", Type: proto.ColumnType_INT, Description: "Elevation above sea level in meters."},
			{Name: "usage_type", Type: proto.ColumnType_STRING, Description: "Usage type classification of ISP or company."},
			{Name: "mcc", Type: proto.ColumnType_STRING, Description: "Mobile Country Codes (MCC) as defined in ITU E.212."},
			{Name: "mnc", Type: proto.ColumnType_STRING, Description: "Mobile Network Code (MNC)."},
			{Name: "mobile_brand", Type: proto.ColumnType_STRING, Description: "Commercial brand associated with the mobile carrier."},
			{Name: "address_type", Type: proto.ColumnType_STRING, Description: "Anycast/Unicast/Multicast/Broadcast."},
			{Name: "district", Type: proto.ColumnType_STRING, Description: "District name."},
			{Name: "ads_category", Type: proto.ColumnType_STRING, Description: "Domain category code based on IAB Tech Lab Content Taxonomy."},
			{Name: "ads_category_name", Type: proto.ColumnType_STRING, Description: "Domain category based on IAB Tech Lab Content Taxonomy."},

			// Nested columns
			{Name: "continent", Type: proto.ColumnType_JSON, Description: "Continent details."},
			{Name: "country", Type: proto.ColumnType_JSON, Description: "Country details."},
			{Name: "region", Type: proto.ColumnType_JSON, Description: "Region details."},
			{Name: "city", Type: proto.ColumnType_JSON, Description: "City details."},
			{Name: "time_zone_info", Type: proto.ColumnType_JSON, Description: "Time zone details."},
			{Name: "geotargeting", Type: proto.ColumnType_JSON, Description: "Metro code."},
			{Name: "proxy", Type: proto.ColumnType_JSON, Description: "Proxy details."},
		},
	}
}

func listIP(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connectGeolocation(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ip2locationio_geolocation.listIP", "geolocation_connection_error", err)
		return nil, err
	}

	ipStr := ""
	if d.EqualsQuals["ip"] != nil {
		ipStr = d.EqualsQuals["ip"].GetInetValue().GetAddr()
		plugin.Logger(ctx).Debug("ip2locationio_geolocation.listIP", "ipStr", ipStr)
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		plugin.Logger(ctx).Warn("ip2locationio_geolocation.listIP", "invalid_ip", ip, "status", "Must supply a valid IP address.")
	}

	lang := ""
	if d.EqualsQuals["lang"] != nil {
		lang = d.EqualsQuals["lang"].GetStringValue()
		plugin.Logger(ctx).Debug("ip2locationio_geolocation.listIP", "lang", lang)
	}

	res, err := conn.LookUp(ip.String(), lang) // language parameter only available with Plus and Security plans

	plugin.Logger(ctx).Debug("ip2locationio_geolocation.listIP", "results", res)
	if err != nil {
		plugin.Logger(ctx).Error("ip2locationio_geolocation.listIP", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, res)

	return nil, nil
}
