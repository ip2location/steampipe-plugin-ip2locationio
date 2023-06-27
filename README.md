![image](https://hub.steampipe.io/images/plugins/ip2location/ip2locationio-social-graphic.png)

# IP2Location.io Plugin for Steampipe

Use SQL to retrieve geolocation or WHOIS info in minutes using [ip2location.io](https://ip2location.io).

- **[Get started →](https://hub.steampipe.io/plugins/ip2location/ip2locationio)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/ip2location/ip2locationio/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/ip2location/steampipe-plugin-ip2locationio/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```sh
steampipe plugin install ip2location/ip2locationio
```

Configure the server address in `~/.steampipe/config/ip2locationio.spc`:

```hcl
connection "ip2locationio" {
  plugin = "ip2location/ip2locationio"

  # API key for requests. Required.
  # Sign up for a free key at https://www.ip2location.io/pricing
  # This can also be set via the `IP2LOCATIONIO_API_KEY` environment variable.
  # api_key = "Q5Z8QS544RKC2VK4P3ZH7YW3C16MDCBW"
}
```

Or through environment variables:

```sh
export IP2LOCATIONIO_API_KEY=Q5Z8QS544RKC2VK4P3ZH7YW3C16MDCBW
```

Run steampipe:

```sh
steampipe query
```

Query IP geolocation:

```sql
select
   country_code,
   country_name,
   region_name,
   city_name 
from
   ip2locationio_geolocation 
where
   ip = '8.8.8.8';
```

```
+--------------+--------------------------+-------------+---------------+
| country_code | country_name             | region_name | city_name     |
+--------------+--------------------------+-------------+---------------+
| US           | United States of America | California  | Mountain View |
+--------------+--------------------------+-------------+---------------+
```

Query WHOIS data:

```sql
select
   domain,
   domain_id,
   status,
   create_date 
from
   ip2locationio_whois 
where
   domain = 'google.com';
```

```
+------------+-------------------------+---------------------------------------------------------------------------+----------------------+
| domain     | domain_id               | status                                                                    | create_date          |
+------------+-------------------------+---------------------------------------------------------------------------+----------------------+
| google.com | 2138514_DOMAIN_COM-VRSN | clientUpdateProhibited (https://www.icann.org/epp#clientUpdateProhibited) | 1997-09-15T07:00:00Z |
+------------+-------------------------+---------------------------------------------------------------------------+----------------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/ip2location/steampipe-plugin-ip2locationio.git
cd steampipe-plugin-ip2locationio
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
nano ~/.steampipe/config/ip2locationio.spc
```

Try it!

```
steampipe query
> .inspect ip2locationio
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/ip2location/steampipe-plugin-ip2locationio/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [ip2location.io Plugin](https://github.com/ip2location/steampipe-plugin-ip2locationio/labels/help%20wanted)
