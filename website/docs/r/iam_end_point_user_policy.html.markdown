---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_policy"
description: |-
  Enables creation of local users on endpoints.
---

# Resource: intersight_iam_end_point_user_policy
Enables creation of local users on endpoints.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `end_point_user_roles`:(Array) An array of relationships to iamEndPointUserRole resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `password_properties`:(HashMap) - Set password properties for endpoint users. 
This complex property has following sub-properties:
  + `enable_password_expiry`:(bool) Enables password expiry on the endpoint. 
  + `enforce_strong_password`:(bool) Enables a strong password policy. Strong password requirements: A. The password must have a minimum of 8 and a maximum of 20 characters. B. The password must not contain the User's Name. C. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &, *, -, _, +, =). 
  + `force_send_password`:(bool) User password will always be sent to endpoint device. If the option is not selected, then user password will be sent to endpoint device for new users and if user password is changed for existing users. 
  + `grace_period`:(int) Time period until when you can use the existing password, after it expires. 
  + `notification_period`:(int) The duration after which the password will expire. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `password_expiry_duration`:(int) Set time period for password expiration. Value should be greater than notification period and grace period. 
  + `password_history`:(int) Tracks password change history. Specifies in number of instances, that the new password was already used. 
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
resource "intersight_iam_end_point_user_policy" "user_policy1" {
  name = "user_policy1"
  description = "test policy"

  password_properties {
    enforce_strong_password = true
    enable_password_expiry = true
    password_expiry_duration = 50
    password_history = 5
    notification_period = 1
    grace_period = 2
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}
```

## Import
`intersight_iam_end_point_user_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_iam_end_point_user_policy.example 1234567890987654321abcde
``` 