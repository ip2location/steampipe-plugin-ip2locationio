# Table: ip2locationio_whois

Get WHOIS domain information about a domain name.
The `ip2locationio_whois` table requires the `domain` field to be specified in all queries, defining the domain name to lookup.

## Examples

### Info for a specific domain name

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

### Nested info for a specific domain name

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
