---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_appliance_role"
description: |-
  Configuration object sent by user to create an appliance port.
---

# Data Source: intersight_fabric_appliance_role
Configuration object sent by user to create an appliance port.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_appliance_role.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port.* `Auto` - Admin configurable speed AUTO ( default ).* `1Gbps` - Admin configurable speed 1Gbps.* `10Gbps` - Admin configurable speed 10Gbps.* `25Gbps` - Admin configurable speed 25Gbps.* `40Gbps` - Admin configurable speed 40Gbps.* `100Gbps` - Admin configurable speed 100Gbps. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `fec`:(string) Forward error correction configuration for the port.* `Auto` - Forward error correction option 'Auto'.* `Cl91` - Forward error correction option 'cl91'.* `Cl74` - Forward error correction option 'cl74'. 
* `mode`:(string) Port mode to be set on the appliance port.* `trunk` - Trunk Mode Switch Port Type.* `access` - Access Mode Switch Port Type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `priority`:(string) The 'name' of the System QoS Class.* `Best Effort` - QoS Priority for Best-effort traffic.* `FC` - QoS Priority for FC traffic.* `Platinum` - QoS Priority for Platinum traffic.* `Gold` - QoS Priority for Gold traffic.* `Silver` - QoS Priority for Silver traffic.* `Bronze` - QoS Priority for Bronze traffic. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
 