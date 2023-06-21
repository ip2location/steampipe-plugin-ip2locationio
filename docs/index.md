---
organization: IP2Location
category: ["saas"]
icon_url: "/images/plugins/ip2location/ip2locationio.svg"
brand_color: "#535eab"
display_name: "ip2location.io"
short_name: "ip2locationio"
description: "Steampipe plugin to query IP geolocation or WHOIS information from ip2location.io."
og_description: "Query ip2location.io with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/ip2location/ip2locationio-social-graphic.png"
---

# ip2location.io + Steampipe

[ip2location.io](https://ip2location.io) is an API for IP address information (e.g. location) or WHOIS data (domain registration info).

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

Query IP geolocation nested data:

```sql
select
   country_code,
   country ->> 'capital' as capital_city,
   country['translation'] as translation 
from
   ip2locationio_geolocation 
where
   ip = '8.8.8.8' 
   and lang = 'es';
```

```
+--------------+------------------+---------------------------------------------------------+
| country_code | capital_city     | translation                                             |
+--------------+------------------+---------------------------------------------------------+
| US           | Washington, D.C. | {"lang":"es","value":"Estados Unidos de América (los)"} |
+--------------+------------------+---------------------------------------------------------+
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

Query WHOIS nested data:

```sql
select
   domain,
   domain_id,
   registrar ->> 'name' as registrar_name,
   nameservers 
from
   ip2locationio_whois 
where
   domain = 'google.com';
```

```
+------------+-------------------------+-------------------+-----------------------------------------------------------------------+
| domain     | domain_id               | registrar_name    | nameservers                                                           |
+------------+-------------------------+-------------------+-----------------------------------------------------------------------+
| google.com | 2138514_DOMAIN_COM-VRSN | MarkMonitor, Inc. | ["ns2.google.com","ns4.google.com","ns3.google.com","ns1.google.com"] |
+------------+-------------------------+-------------------+-----------------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/ip2location/ip2locationio/tables)**

## Get started

### Install

Download and install the latest ip2location.io plugin:

```bash
steampipe plugin install ip2locationio
```

### Configuration

Installing the latest ip2locationio plugin will create a config file (`~/.steampipe/config/ip2locationio.spc`) with a single connection named `ip2locationio`:

```hcl
connection "ip2locationio" {
  plugin = "ip2location/ip2locationio"

  # Required: Set your IP2Location.io API key.
  # Sign up for a free key at https://www.ip2location.io/pricing
  # api_key = "Q5Z8QS544RKC2VK4P3ZH7YW3C16MDCBW"
}
```

- `api_key` - Required API key from ip2location.io.

## Get involved

- Open source: https://github.com/ip2location/steampipe-plugin-ip2locationio
- Community: [Slack Channel](https://steampipe.io/community/join)
