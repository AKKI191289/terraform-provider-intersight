---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_equipment_physical_def"
description: |-
  Type to represent additional switch specific capabilities.
---

# Resource: intersight_capability_equipment_physical_def
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments are supported:
* `depth`:(float) Depth information for a Switch/Fabric-Interconnect. 
* `height`:(float) Height information for a Switch/Fabric-Interconnect. 
* `max_power`:(int) Max Power information for a Switch/Fabric-Interconnect. 
* `min_power`:(int) Min Power information for a Switch/Fabric-Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `nominal_power`:(int) Nominal Power information for a Switch/Fabric-Interconnect. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `weight`:(float) Weight information for a Switch/Fabric-Interconnect. 
* `width`:(float) Width information for a Switch/Fabric-Interconnect. 


## Import
`intersight_capability_equipment_physical_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_equipment_physical_def.example 1234567890987654321abcde
``` 