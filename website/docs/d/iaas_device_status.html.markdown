---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_device_status"
description: |-
  List of infra accounts managed by UCSD.
---

# Data Source: intersight_iaas_device_status
List of infra accounts managed by UCSD.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_device_status.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_name`:(string) The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD. 
* `account_type`:(string) The UCSD Infra Account type. 
* `category`:(string) The UCSD Infra Account category. 
* `claim_status`:(string) Describes if the device is claimed in Intersight or not.* `Unknown` - If UCS Director managed account claim status information is unknown.* `Yes` - If UCS Director managed account is claimed in Intersight.* `No` - If UCS Director managed account is not claimed in Intersight.* `Not Applicable` - If UCS Director managed account is not capable of providing claim status information. 
* `connection_status`:(string) Describes about the connection status between the UCSD and the actual end device. 
* `device_model`:(string) Describes about the device model. 
* `device_vendor`:(string) Describes about the device vendor/manufacturer of the device. 
* `device_version`:(string) Describes about the current firmware version running on the device. 
* `ip_address`:(string) The IPAddress of the device. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pod`:(string) Describes about the pod to which this device belongs to in UCSD. 
* `pod_type`:(string) Describes about the podType of Pod to which this device belongs to in UCSD. 
 