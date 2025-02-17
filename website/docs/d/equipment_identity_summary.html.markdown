---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_identity_summary"
description: |-
  Consolidated view of all equipment identities.
---

# Data Source: intersight_equipment_identity_summary
Consolidated view of all equipment identities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_identity_summary.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `adapter_serial`:(string) Serial Identifier of an adapter participating in SWM. 
* `admin_action`:(string) Updated by UI/API to trigger specific chassis action type.* `None` - No operation value for maintenance actions on an equipment.* `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight.* `Recommission` - Recommission the equipment.* `Reack` - Reacknowledge the equipment and discover it again.* `Remove` - Remove the equipment permanently from Intersight management. 
* `admin_action_state`:(string) The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `chassis_id`:(int) Chassis Identifier of a blade server. 
* `firmware_supportability`:(string) Describes whether the running CIMC version supports Intersight managed mode.* `Unknown` - The running firmware version is unknown.* `Supported` - The running firmware version is known and supports IMM mode.* `NotSupported` - The running firmware version is known and does not support IMM mode. 
* `identifier`:(int) Numeric Identifier assigned by the management system to the equipment. 
* `nr_lifecycle`:(string) The equipment's lifecycle status.* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.* `Active` - Default Lifecycle State for a physical entity.* `Decommissioned` - Decommission Lifecycle state.* `DecommissionInProgress` - Decommission Inprogress Lifecycle state.* `RecommissionInProgress` - Recommission Inprogress Lifecycle state.* `OperationFailed` - Failed Operation Lifecycle state.* `ReackInProgress` - ReackInProgress Lifecycle state.* `RemoveInProgress` - RemoveInProgress Lifecycle state.* `Discovered` - Discovered Lifecycle state.* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity.* `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity. 
* `model`:(string) The vendor provided model name for the equipment. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presence`:(string) The presence state of the blade server.* `Unknown` - The default presence state.* `Equipped` - The server is equipped in the slot.* `EquippedMismatch` - The slot is equipped, but there is another server currently inventoried in the slot.* `Missing` - The server is not present in the given slot. 
* `serial`:(string) The serial number of the equipment. 
* `slot_id`:(int) Chassis slot number of a blade server. 
* `source_object_type`:(string) The source object type of this view MO. 
* `switch_id`:(int) Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B. 
* `vendor`:(string) The manufacturer of the equipment. 
 