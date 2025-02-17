---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_fcoe_uplink_pc_role"
description: |-
  Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
---

# Data Source: intersight_fabric_fcoe_uplink_pc_role
Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_fcoe_uplink_pc_role.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port.* `Auto` - Admin configurable speed AUTO ( default ).* `1Gbps` - Admin configurable speed 1Gbps.* `10Gbps` - Admin configurable speed 10Gbps.* `25Gbps` - Admin configurable speed 25Gbps.* `40Gbps` - Admin configurable speed 40Gbps.* `100Gbps` - Admin configurable speed 100Gbps. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pc_id`:(int) Unique Identifier of the port-channel, local to this switch. 
* `udld_admin_state`:(string) Admin configured state for UDLD for this port.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
 