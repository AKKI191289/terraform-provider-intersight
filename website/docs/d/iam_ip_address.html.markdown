---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_ip_address"
description: |-
  Add an IP address to enable IP address based access management.
---

# Data Source: intersight_iam_ip_address
Add an IP address to enable IP address based access management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_ip_address.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `address`:(string) The Trusted IP range's address. IP address, CIDR range, and IP address range formats are supported. For example '12.13.14.15', '12.13.14.0/24', and '12.13.14.15-12.13.14.200'. Reserved IP ranges '127.0.0.1', '10.0.0.0/8', '172.16.0.0/12', and '192.168.0.0/16' are not allowed. 
* `description`:(string) Description of Trusted IP address range. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 