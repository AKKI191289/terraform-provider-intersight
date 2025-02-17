---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_supported_version"
description: |-
  The supported operating system version by SCU. The API is mainly for UI operation. In the software management page,
operating system configuration will be created by providing the vendor and version. The version will be filtered
based on vendor.
---

# Data Source: intersight_os_supported_version
The supported operating system version by SCU. The API is mainly for UI operation. In the software management page,
operating system configuration will be created by providing the vendor and version. The version will be filtered
based on vendor.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_supported_version.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `version_name`:(string) The OsInstall Supported Operating System Version Name. 
 