---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_unit_descriptor"
description: |-
  Descriptor that uniquely identifies an adaptor.
---

# Data Source: intersight_capability_adapter_unit_descriptor
Descriptor that uniquely identifies an adaptor.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_adapter_unit_descriptor.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `connectivity_order`:(string) Order in which the ports are connected; sequential or cyclic. 
* `description`:(string) Detailed information about the endpoint. 
* `ethernet_port_speed`:(int) The port speed for ethernet ports in Mbps. 
* `fibre_channel_port_speed`:(int) The port speed for fibre channel ports in Mbps. 
* `fibre_channel_scsi_ioq_limit`:(int) The number of SCSI I/O Queue resources to allocate. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_dce_ports`:(int) Number of Dce Ports for the adaptor. 
* `prom_card_type`:(string) Prom card type for the adaptor. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 