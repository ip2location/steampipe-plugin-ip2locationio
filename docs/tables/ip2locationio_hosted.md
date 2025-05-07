# Table: ip2locationio_hosted

Get the list of domains hosted by an IP address.
The `ip2locationio_hosted` table requires the `ip` field to be specified in all queries, defining the IP address to lookup.

## Examples

### Domains hosted by a specific IP address (default is to return the first page of the result)

```sql
select
  total_domains,
  page,
  per_page,
  total_pages,
  domains 
from
  ip2locationio_hosted 
where
  ip = '8.8.8.8';
```

### Domains hosted by a specific IP address by result page

```sql
select
  total_domains,
  page,
  per_page,
  total_pages,
  domains 
from
  ip2locationio_hosted 
where
  ip = '8.8.8.8' 
  and page = 2;
```
