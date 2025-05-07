# Table: ip2locationio_geolocation

Get location and other information about an IP address.
The `ip2locationio_geolocation` table requires the `ip` field to be specified in all queries, defining the IP address to lookup.

## Examples

### Info for a specific IP address

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

### Nested info for specific IP address

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
