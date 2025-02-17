---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_boot_device"
description: |-
  Actual boot devices of the system as enumerated by BIOS.
---

# Data Source: intersight_bios_boot_device
Actual boot devices of the system as enumerated by BIOS.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bios_boot_device.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_name`:(string) Name of the Configured Boot Device. 
* `device_type`:(string) Type of the Configured Boot Device. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 