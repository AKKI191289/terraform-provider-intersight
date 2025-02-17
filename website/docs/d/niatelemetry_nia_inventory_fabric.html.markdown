---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_inventory_fabric"
description: |-
  Inventory Object available for Fabric-scoped attributes.
---

# Data Source: intersight_niatelemetry_nia_inventory_fabric
Inventory Object available for Fabric-scoped attributes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nia_inventory_fabric.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `anycast_gw_mac`:(string) Returns the aycast gateway mac. 
* `bgp_established_interface_count`:(int) Counts the number of BGP interfaces that are in established state. 
* `bgw_interface_up_count`:(int) Count number of active interfaces on border gateways. 
* `border_gateway_spine_count`:(int) Count number of border gateway spines in the fabric inventory. 
* `border_leaf_count`:(int) Count number of border leafs in the fabric inventory. 
* `dci_subnet_range`:(string) Returns the dci subnet range. 
* `dci_subnet_target_mask`:(string) Returns the dci subnet target mask. 
* `dcnmtracker_enabled`:(bool) Returns the value of the dcnmtrackerEnabled field. 
* `ebgp_evpn_link_up_count`:(int) Count number of ebgp evpn active interfaces. 
* `fabric_id`:(string) Uniquely identifies a fabric. 
* `fabric_name`:(string) Returns the value of the Name of a fabric. 
* `is_bgw_present`:(bool) Checks if border gateway is present in the fabric inventory. 
* `is_ngoam_enabled`:(bool) Returns if ngoam is enabled. 
* `is_scheduled_back_up_enabled`:(bool) Returns if the scheduled backup is enabled. 
* `leaf_count`:(int) Returns total number of leafs in the fabric. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nxos_vni_bw_sites_count`:(int) Returns the count of vnis between sites. 
* `nxos_vrf_bw_sites_count`:(int) Returns the count of vrfs between sites. 
* `nxos_vrf_count`:(int) Returns the value of the nxosVrfCount field. 
* `serial`:(string) Serial number of device being inventoried. The serial number is unique per device. 
* `site_name`:(string) Name of fabric domain of the controller. 
* `spine_count`:(int) Returns total number of spines in the fabric. 
* `vlan_vni_mappings`:(string) VLAN to VNI mappings configured in the DCNM. 
* `vni_ip_count`:(int) Count number of IP addresses configured in the DCNM networks. 
 