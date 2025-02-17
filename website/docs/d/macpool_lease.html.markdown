---
subcategory: "macpool"
layout: "intersight"
page_title: "Intersight: intersight_macpool_lease"
description: |-
  Lease represents a single MAC address that is part of the universe, allocated either from a pool or through static assignment.
---

# Data Source: intersight_macpool_lease
Lease represents a single MAC address that is part of the universe, allocated either from a pool or through static assignment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_macpool_lease.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `allocation_type`:(string) Type of the lease allocation either static or dynamic (i.e via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `mac_address`:(string) MAC address allocated for pool-based allocation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 