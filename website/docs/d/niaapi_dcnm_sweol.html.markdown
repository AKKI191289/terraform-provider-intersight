---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_sweol"
description: |-
  The software end of life notice for DCNM.
---

# Data Source: intersight_niaapi_dcnm_sweol
The software end of life notice for DCNM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_dcnm_sweol.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `affected_versions`:(string) String contains the Release versions affected by this notice, seperated by comma. 
* `announcement_date`:(string) Date time of this notice Announced. 
* `announcement_date_epoch`:(int) Epoch time of this notice Announced. 
* `bulletin_no`:(string) The bulletinno of this software release end of life notice. 
* `description`:(string) The description of this software release end of life notice. 
* `endof_new_service_attachment_date`:(string) Date time of End of New service attachment . 
* `endof_new_service_attachment_date_epoch`:(int) Epoch time of End of New service attachment . 
* `endof_service_contract_renewal_date`:(string) Date time of End of Renewal of service contract . 
* `endof_service_contract_renewal_date_epoch`:(int) Epoch time of End of Renewal of service contract . 
* `endof_sw_maintenance_date`:(string) Date time of End of Maintenance. 
* `endof_sw_maintenance_date_epoch`:(int) Epoch time of End of Maintenance. 
* `headline`:(string) The title of this software release end of life notice. 
* `last_dateof_support`:(string) Date time of Last day of Support . 
* `last_dateof_support_epoch`:(int) Epoch time of Last day of Support . 
* `last_ship_date`:(string) Date time of Lastship Date. 
* `last_ship_date_epoch`:(int) Epoch time of Lastship Date. 
* `migration_url`:(string) The URL of this migration notice. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `software_eol_url`:(string) Software end of life notice URL link to the notice webpage. 
 