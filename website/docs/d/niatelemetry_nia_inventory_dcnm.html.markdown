---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_inventory_dcnm"
description: |-
  Inventory Object available for DCNM-scoped attributes.
---

# Data Source: intersight_niatelemetry_nia_inventory_dcnm
Inventory Object available for DCNM-scoped attributes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nia_inventory_dcnm.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dev`:(bool) Returns the value of the dev Field. 
* `epld_image_count`:(int) Number of EPLD images uploaded to DCNM. 
* `ha_enabled`:(bool) Returns the value of the haEnabled field. 
* `ha_replication_status`:(string) Returns the value of the haReplicationStatus field. 
* `install`:(string) Returns the value of the install field. 
* `is_isn_configured`:(bool) Returns true if ISN is configured. 
* `is_media_controller`:(bool) Returns the value of the isMediaController field. 
* `is_smart_license_enabled`:(bool) Returns true if the Smart license is enabled and is in use. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_fabrics`:(int) Returns total number of fabrics in DCNM set-up. 
* `num_fabrics_in_msd`:(int) Returns the number of fabrics in msd. 
* `num_ingress_replication_fabrics`:(int) Returns the number of fabrics that have ingress replication type. 
* `num_local_users`:(int) Returns the number of local users other than admin user. 
* `num_msd`:(int) Returns the number of MSD fabrics. 
* `num_svi_vrf_count`:(int) Returns the number of svi interfaces configured for VRF vlans. 
* `num_trm_enabled_count`:(int) Returns the number of links where TRM is enabled. 
* `num_upg_users`:(int) Number of users who have upgrade privileges excluding the admin. 
* `nxos_image_count`:(int) Number of NXOS images uploaded to DCNM. 
* `serial`:(string) Serial number of device being inventoried. The serial number is unique per device. 
* `site_name`:(string) Name of fabric domain of the controller. 
* `underlay_peering_active_links_count`:(int) Returns the number of underlay peering active links. 
* `upg_job_count`:(int) Number of upgrade jobs configured on DCNM. 
* `upg_status`:(string) Upgrade status of jobs created on DCNM. 
* `nr_version`:(string) Returns the value of the version field. 
 