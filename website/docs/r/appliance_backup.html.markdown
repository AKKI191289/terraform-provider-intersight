---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup"
description: |-
  Backup tracks all backup requests to create a full system backup of the Intersight
Appliance. There will be only one Backup managed object with a 'Started' state at
any time. All other Backup managed objects will be in terminal states.
---

# Resource: intersight_appliance_backup
Backup tracks all backup requests to create a full system backup of the Intersight
Appliance. There will be only one Backup managed object with a 'Started' state at
any time. All other Backup managed objects will be in terminal states.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `elapsed_time`:(int)(Computed) Elapsed time in seconds since the backup process has started. 
* `end_time`:(string)(Computed) End date and time of the backup process. 
* `filename`:(string) Backup filename to backup or restore. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `messages`:
                (Array of schema.TypeString) -
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password to authenticate the fileserver. 
* `protocol`:(string) Communication protocol used by the file server (e.g. scp or sftp).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server. 
* `remote_host`:(string) Hostname of the remote file server. 
* `remote_path`:(string) File server directory to copy the file. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). 
* `start_time`:(string)(Computed) Start date and time of the backup process. 
* `status`:(string)(Computed) Status of the backup managed object.* `Started` - Backup or restore process has started.* `Created` - Backup or restore is in created state.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `Copied` - Backup file has been copied. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `username`:(string) Username to authenticate the fileserver. 


## Import
`intersight_appliance_backup` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_appliance_backup.example 1234567890987654321abcde
``` 