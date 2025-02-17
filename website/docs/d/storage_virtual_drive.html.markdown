---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_virtual_drive"
description: |-
  A Virtual Disk Drive or Logical Unit Number.
---

# Data Source: intersight_storage_virtual_drive
A Virtual Disk Drive or Logical Unit Number.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_virtual_drive.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_policy`:(string) The access policy of the virtual drive. 
* `actual_write_cache_policy`:(string) The current write cache policy of the virtual drive. 
* `available_size`:(string) Available storage capacity of the virtual drive. 
* `block_size`:(string) Block size of the virtual drive. 
* `bootable`:(string) The virtual drive bootable state. 
* `config_state`:(string) The configuration state of the virtual drive. 
* `configured_write_cache_policy`:(string) The requested write cache policy of the virtual drive. 
* `connection_protocol`:(string) The connection protocol of the virtual drive. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `drive_cache`:(string) The state of the drive cache of the virtual drive. 
* `drive_security`:(string) The driveSecurity state of the Virtual drive. 
* `drive_state`:(string) The state of the Virtual drive. 
* `io_policy`:(string) The Input/Output Policy defined on the Virtual drive. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Virtual drive. 
* `num_blocks`:(string) Number of Blocks on the Physical Disk. 
* `oper_state`:(string) The current operational state of Virtual drive. 
* `operability`:(string) The current operability state of Virtual drive. 
* `physical_block_size`:(string) The block size of the the virtual drive. 
* `presence`:(string) The presence status of the virtual drive. 
* `read_policy`:(string) The read-ahead cache mode of the virtual drive. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `security_flags`:(string) The security flags set for this virtual drive. 
* `serial`:(string) This field identifies the serial of the given component. 
* `size`:(string) The size of the virtual drive in MB. 
* `strip_size`:(string) The strip size is the portion of a stripe that resides on a single drive in the drive group, this is measured in KB. 
* `type`:(string) The raid type of the virtual drive. 
* `uuid`:(string) The uuid of the virtual drive. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vendor_uuid`:(string) The UUID value of the vendor. 
* `virtual_drive_id`:(string) The identifier for this Virtual drive. 
 