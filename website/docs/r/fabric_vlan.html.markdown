---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vlan"
description: |-
  Configuration object for Virtual LAN.
---

# Resource: intersight_fabric_vlan
Configuration object for Virtual LAN.
## Argument Reference
The following arguments are supported:
* `auto_allow_on_uplinks`:(bool) Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI. 
* `eth_network_policy`:(HashMap) - A reference to a fabricEthNetworkPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `is_native`:(bool) Used to define whether this VLAN is to be classified as 'native' for traffic in this FI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `multicast_policy`:(HashMap) - A reference to a fabricMulticastPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `name`:(string) The 'name' used to identify this VLAN. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vlan_id`:(int) The identifier for this Virtual LAN. 


## Import
`intersight_fabric_vlan` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_vlan.example 1234567890987654321abcde
``` 