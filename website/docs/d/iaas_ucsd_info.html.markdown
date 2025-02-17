---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_info"
description: |-
  UCS Director accounts managed by Intersight.
---

# Data Source: intersight_iaas_ucsd_info
UCS Director accounts managed by Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_ucsd_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_id`:(string) Moid of the UCS Director device connector's asset.DeviceRegistration. 
* `guid`:(string) Unique ID of UCS Director getting registerd with Intersight. 
* `host_name`:(string) The UCS Director hostname for management. 
* `ip`:(string) The UCS Director IP address for management. 
* `last_backup`:(string) Last successful backup created for this UCS Director appliance if backup is configured. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_type`:(string) NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node. 
* `product_name`:(string) The UCS Director product name. 
* `product_vendor`:(string) The UCS Director product vendor. 
* `product_version`:(string) The UCS Director product/platform version. 
* `status`:(string) The UCS Director status. Possible values are Active, Inactive, Unknown. 
 