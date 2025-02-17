---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_capability"
description: |-
  Type to represent additional switch specific capabilities.
---

# Resource: intersight_capability_switch_capability
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments are supported:
* `default_fcoe_vlan`:(int) Default Fcoe VLAN associated with this switch. 
* `dynamic_vifs_supported`:(bool) Dynamic VIFs support on this switch. 
* `fan_modules_supported`:(bool) Fan Modules support on this switch. 
* `fc_end_host_mode_reserved_vsans`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `fc_uplink_ports_auto_negotiation_supported`:(bool) Fc Uplink ports auto negotiation speed support on this switch. 
* `locator_beacon_supported`:(bool) Locator Beacon LED support on this switch. 
* `max_ports`:(int) Maximum allowed physical ports on this switch. 
* `max_slots`:(int) Maximum allowed physical slots on this switch. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `network_limits`:(HashMap) - List of network limitations for this switch. 
This complex property has following sub-properties:
  + `max_compressed_port_vlan_count`:(int) Maximum Compressed configurable VLANs on Switch/Fabric-Interconnect. 
  + `max_uncompressed_port_vlan_count`:(int) Maximum configurable VLANs on Switch/Fabric-Interconnect. 
  + `maximum_active_traffic_monitoring_sessions`:(int) Maximum configured and enabled Traffic Monitoring sessions on this Switch/Fabric-Interconnect. 
  + `maximum_ethernet_port_channels`:(int) Maximum configurable Ethernet port-channels on Switch/Fabric-Interconnect. 
  + `maximum_ethernet_uplink_ports`:(int) Maximum configurable Ethernet Uplink ports on Switch/Fabric-Interconnect. 
  + `maximum_fc_port_channel_members`:(int) Maximum configurable Fibre Channel port-channel member ports on Switch/Fabric-Interconnect. 
  + `maximum_fc_port_channels`:(int) Maximum configurable Fibre Channel port-channels on Switch/Fabric-Interconnect. 
  + `maximum_igmp_groups`:(int) Maximum configurable IGMP Groups on Switch/Fabric-Interconnect. 
  + `maximum_port_channel_members`:(int) Maximum configurable ports per each port-channel on Switch/Fabric-Interconnect. 
  + `maximum_primary_vlan`:(int) Maximum configurable Primary Private VLANs on Switch/Fabric-Interconnect. 
  + `maximum_secondary_vlan`:(int) Maximum configurable Secondary Private VLANs on Switch/Fabric-Interconnect. 
  + `maximum_secondary_vlan_per_primary`:(int) Maximum configurable Secondary VLANs per each Primary VLAN on Switch/Fabric-Interconnect. 
  + `maximum_vifs`:(int) Maximum allowes VIFs on Switch/Fabric-Interconnect. 
  + `maximum_vlans`:(int) Maximum configurable VLANs on Switch/Fabric-Interconnect. 
  + `minimum_active_fans`:(int) Minimum required fans in 'active' state for this Switch/Fabric-Interconnect. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `ports_supporting100g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting10g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting1g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting25g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting40g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_breakout`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_fcoe`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_server_role`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `reserved_vsans`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `sereno_netflow_supported`:(bool) Sereno Adaptor with Netflow support on this switch. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `storage_limits`:(HashMap) - List of storage limitations for this switch. 
This complex property has following sub-properties:
  + `maximum_user_zone_count`:(int) Maximum user zones per Switch/Fabric-Interconnect. 
  + `maximum_virtual_fc_interfaces`:(int) Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect. 
  + `maximum_virtual_fc_interfaces_per_blade_server`:(int) Maximum configurable Virtual Fibre Channel interfaces per blade. 
  + `maximum_vsans`:(int) Maximum configurable VSANs on Switch/Fabric-Interconnect. 
  + `maximum_zone_count`:(int) Zone limit per Switch/Fabric-Interconnect. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `switching_mode_capabilities`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `switching_mode`:(string) Switching mode type (endhost, switch) of the switch.* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
  + `vp_compression_supported`:(bool) VP Compression support on this switch. 
* `system_limits`:(HashMap) - List of system limitations for this switch. 
This complex property has following sub-properties:
  + `maximum_chassis_count`:(int) Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect. 
  + `maximum_fex_per_domain`:(int) Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect. 
  + `maximum_servers_per_domain`:(int) Maximum UCS servers per Switch/Fabric-Interconnect. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `unified_ports`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `unified_rule`:(string) The Slider rule for Unified ports on this switch. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 


## Import
`intersight_capability_switch_capability` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_switch_capability.example 1234567890987654321abcde
``` 