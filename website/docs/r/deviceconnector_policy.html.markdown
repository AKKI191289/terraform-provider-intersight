---
subcategory: "deviceconnector"
layout: "intersight"
page_title: "Intersight: intersight_deviceconnector_policy"
description: |-
  Policy to control configuration changes allowed from Cisco IMC.
---

# Resource: intersight_deviceconnector_policy
Policy to control configuration changes allowed from Cisco IMC.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `lockout_enabled`:(bool) Enables configuration lockout on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 

## Usage Example
### Resource Creation
```hcl
resource "intersight_deviceconnector_policy" "dcp1" {
  name            = "device_con1"
  description     = "test policy"
  lockout_enabled = true
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}
```

## Import
`intersight_deviceconnector_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_deviceconnector_policy.example 1234567890987654321abcde
``` 