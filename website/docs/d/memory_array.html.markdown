---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_array"
description: |-
  Holder housing multiple memory units.
---

# Data Source: intersight_memory_array
Holder housing multiple memory units.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_array.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `array_id`:(int) The instance number of the memory array. 
* `cpu_id`:(int) ID of the CPU that access this memory array. 
* `current_capacity`:(string) Current capacity of all the memory units on a server. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `error_correction`:(string) The primary hardware error detection or correction method supported by the memory array. 
* `max_capacity`:(string) Maximum capacity of all the memory units on a server. 
* `max_devices`:(string) The maximum number of slots or sockets available for memory devices in the memory array. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_power_state`:(string) The power state indicator of the memory array. 
* `presence`:(string) The presence of atleast one memory device in the array. Valid values are 'equipped' and 'absent'. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 