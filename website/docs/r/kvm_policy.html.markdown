---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_policy"
description: |-
  Policy to configure KVM Launch settings.
---

# Resource: intersight_kvm_policy
Policy to configure KVM Launch settings.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `enable_local_server_video`:(bool) If enabled, displays KVM session on any monitor attached to the server. 
* `enable_video_encryption`:(bool) If enabled, encrypts all video information sent through KVM. 
* `enabled`:(bool) State of the vKVM service on the endpoint. 
* `maximum_sessions`:(int) The maximum number of concurrent KVM sessions allowed. 
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
* `remote_port`:(int) The port used for KVM communication. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 

## Usage Example
### Resource Creation
```hcl
resource "intersight_kvm_policy" "kvm1" {
  name                      = "kvm1"
  description               = "demo kvm policy"
  enabled                   = true
  maximum_sessions          = 3
  remote_port               = 2069
  enable_video_encryption   = true
  enable_local_server_video = true
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}
```

## Import
`intersight_kvm_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kvm_policy.example 1234567890987654321abcde
``` 