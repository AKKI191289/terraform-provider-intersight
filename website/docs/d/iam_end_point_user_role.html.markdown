---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_role"
description: |-
  Mapping of endpoint user to endpoint roles.
---

# Data Source: intersight_iam_end_point_user_role
Mapping of endpoint user to endpoint roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_end_point_user_role.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `change_password`:(bool) Denotes whether password has changed. 
* `enabled`:(bool) Enables the user account on the endpoint. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Valid login password of the user. 
 