---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_pc_operation"
description: |-
  PcOperation objects allows the user to alter the state of the port channel.
---

# Data Source: intersight_fabric_pc_operation
PcOperation objects allows the user to alter the state of the port channel.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_pc_operation.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(string) Admin configured state to disable the port channel.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `config_state`:(string) The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pc_id`:(int) Port Channel Identifier for the collection of ports. 
 