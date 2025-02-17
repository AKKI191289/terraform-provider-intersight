---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_switch_upgrade"
description: |-
  Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
---

# Resource: intersight_firmware_switch_upgrade
Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
## Argument Reference
The following arguments are supported:
* `device`:(HashMap) -(Computed) A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `direct_download`:(HashMap) - Direct download options in case the upgradeType is direct download based upgrade. 
This complex property has following sub-properties:
  + `http_server`:(HashMap) - HTTP Server option when the image source is a local HTTPS server. 
This complex property has following sub-properties:
    + `location_link`:(string) HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. 
    + `mount_options`:(string) Mount option as configured on the HTTP server. Empty if nothing is configured. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `image_source`:(string) Source type referring the image to be downloaded from CCO or from a local HTTPS server.* `cisco` - Image to be downloaded from Cisco software repository.* `localHttp` - Image to be downloaded from a https server. 
  + `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `password`:(string) Password as configured on the local https server. 
  + `upgradeoption`:(string) Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot.* `sd_upgrade_mount_only` - Direct upgrade SD upgrade mount only.* `sd_download_only` - Direct upgrade SD download only.* `sd_upgrade_only` - Direct upgrade SD upgrade only.* `sd_upgrade_full` - Direct upgrade SD upgrade full.* `upgrade_full` - The upgrade downloads or mounts the image, and reboots immediately for an upgrade.* `upgrade_mount_only` - The upgrade downloads or mounts the image. The upgrade happens in next reboot.* `chassis_upgrade_full` - Direct upgrade chassis upgrade full. 
  + `username`:(string) Username as configured on the local https server. 
* `distributable`:(HashMap) - A reference to a firmwareDistributable resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `enable_fabric_evacuation`:(bool) The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen. 
* `file_server`:(HashMap) - Location of the image in user software repository. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `network_elements`:(Array) An array of relationships to networkElement resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `network_share`:(HashMap) - Deprecated (Use 'fileServer' property). Network share options in case of the upgradeType is network share based upgrade. 
This complex property has following sub-properties:
  + `cifs_server`:(HashMap) - CIFS file server option for network share upgrade. 
This complex property has following sub-properties:
    + `file_location`:(string) The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. 
    + `mount_options`:(string) Mount option (Authentication Protocol) as configured on the CIFS Server. Example:ntlmv2.* `none` - The default authentication protocol is decided by the endpoint.* `ntlm` - The external CIFS server is configured with ntlm as the authentication protocol.* `ntlmi` - Mount options of CIFS file server is ntlmi.* `ntlmv2` - Mount options of CIFS file server is ntlmv2.* `ntlmv2i` - Mount options of CIFS file server is ntlmv2i.* `ntlmssp` - Mount options of CIFS file server is ntlmssp.* `ntlmsspi` - Mount options of CIFS file server is ntlmsspi. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `remote_file`:(string)(Computed) Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso. 
    + `remote_ip`:(string)(Computed) CIFS Server Hostname or IP Address. For example:CIFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. 
    + `remote_share`:(string)(Computed) Directory where the image is stored. Example:share/subfolder. 
  + `http_server`:(HashMap) - HTTP (for WWW) file server option for network share upgrade. 
This complex property has following sub-properties:
    + `location_link`:(string) HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. 
    + `mount_options`:(string) Mount option as configured on the HTTP server. Empty if nothing is configured. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
  + `map_type`:(string) File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image.* `nfs` - File server protocol used is NFS.* `cifs` - File server protocol used is CIFS.* `www` - File server protocol used is WWW. 
  + `nfs_server`:(HashMap) - NFS file server option for network share upgrade. 
This complex property has following sub-properties:
    + `file_location`:(string) The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. 
    + `mount_options`:(string) Mount option as configured on the NFS Server. For example:nolock. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `remote_file`:(string)(Computed) Filename of the image in the remote share location. For example:ucs-c220m5-huu-3.1.2c.iso. 
    + `remote_ip`:(string)(Computed) NFS Server Hostname or IP Address. For example:NFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. 
    + `remote_share`:(string)(Computed) Directory where the image is stored. For example:/share/subfolder. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `password`:(string) Password as configured on the file server. 
  + `upgradeoption`:(string) Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade.* `nw_upgrade_full` - Network upgrade option for full upgrade.* `nw_upgrade_mount_only` - Network upgrade mount only. 
  + `username`:(string) Username as configured on the file server. 
* `release`:(HashMap) - A reference to a softwarerepositoryRelease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `skip_estimate_impact`:(bool) User has the option to skip the estimate impact calculation. 
* `status`:(string) Status of the upgrade operation.* `NONE` - Upgrade status is not populated.* `IN_PROGRESS` - The upgrade is in progress.* `SUCCESSFUL` - The upgrade successfully completed.* `FAILED` - The upgrade shows failed status.* `TERMINATED` - The upgrade has been terminated. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `upgrade_impact`:(HashMap) -(Computed) A reference to a firmwareUpgradeImpactStatus resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `upgrade_status`:(HashMap) -(Computed) A reference to a firmwareUpgradeStatus resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `upgrade_type`:(string) Desired upgrade mode to choose either direct download based upgrade or network share upgrade.* `direct_upgrade` - Upgrade mode is direct download.* `network_upgrade` - Upgrade mode is network upgrade. 


## Import
`intersight_firmware_switch_upgrade` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_firmware_switch_upgrade.example 1234567890987654321abcde
``` 