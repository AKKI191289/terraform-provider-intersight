---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_capability_info"
description: |-
  A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, minUcsVersion for HyperFlex version 4.0.1a corresponds to 3.2.3 or minHxdpVersion for HyperFlex Upgrade operation is 4.0.1a etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
---

# Resource: intersight_hyperflex_capability_info
A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, "minUcsVersion" for HyperFlex version "4.0.1a" corresponds to "3.2.3" or "minHxdpVersion" for HyperFlex Upgrade operation is "4.0.1a" etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
## Argument Reference
The following arguments are supported:
* `app_catalog`:(HashMap) - A reference to a hyperflexAppCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `capability_constraints`:(Array)
This complex property has following sub-properties:
  + `constraint_name`:(string) Name or key of the applicable compatibility constraint. 
  + `constraint_value`:(string) Value of the applicable compatibility constraint. Could either be a string value or a regex. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the capability or feature set consisting of a collection of constraint rules and value. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `value`:(string)(Computed) Capability Value which is valid only iff all specified constraints match. 


## Import
`intersight_hyperflex_capability_info` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_capability_info.example 1234567890987654321abcde
``` 