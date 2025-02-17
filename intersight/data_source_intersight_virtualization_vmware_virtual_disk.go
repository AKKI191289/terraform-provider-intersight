package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualizationVmwareVirtualDisk() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualizationVmwareVirtualDiskRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"compatibility_mode": {
				Description: "Compatibility mode of the raw disk mapping (RDM).\n* `notApplicable` - Value specified for any disk which is not of raw device mapping type.\n* `physicalMode` - A disk device backed by a physical compatibility mode raw disk mapping cannot use disk modes, and commands are passed straight through to the LUN indicated by the raw disk mapping.\n* `virtualMode` - A disk device backed by a virtual compatibility mode raw disk mapping can use disk modes.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"controller_key": {
				Description: "Key of the controller on which the disk is created.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"device_name": {
				Description: "Host-specific device the LUN is being accessed through. If the target LUN is not available on the host then it is empty. For example, this could happen if it has accidentally been masked out.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disk_mode": {
				Description: "Persistence mode of the virtual disk. For RDM disks, it is only supported when the raw disk mapping is using virtual compatibility mode.\n* `persistent` - Changes are immediately and permanently written to the virtual disk.\n* `independent_persistent` - Changes are immediately and permanently written to the virtual disk and not affected by snapshots.\n* `independent_nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off and not affected by snapshots.\n* `nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off.\n* `undoable` - Changes to virtual disk are made to a redo log and has the option to commit or undo.\n* `append` - Changes are appended to the redo log and can be revoked by removing the undo log.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disk_type": {
				Description: "Specifies whether the virtual disk is a RDM or a flat disk.\n* `flatDisk` - Specifies that it is a flat disk.\n* `rdmDisk` - Specifies that it is a raw device mapping disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"key": {
				Description: "The internally assigned key of this disk. This entity is not manipulated by users.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"limit": {
				Description: "The utilization of a virtual machine will not exceed this limit, even if there are available resources. Used to ensure a consistent performance of virtual machines independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). The unit is number of I/O per second.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"lun_uuid": {
				Description: "Unique identifier of the LUN accessed by the raw disk mapping (RDM).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the storage disk. Name must be unique within a Datastore.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "Serial ID of the storage device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sharing": {
				Description: "Sharing mode of the virtual disk.\n* `sharingNone` - The virtual disk is not shared.\n* `sharingMultiWriter` - The virtual disk is shared between multiple virtual machines.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"size": {
				Description: "Disk size represented in bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"storage_allocation_type": {
				Description: "Allocation type for the virtual disk.\n* `notApplicable` - Value specified for a disk for which storage allocation type is not applicable.\n* `thin` - A thin provisioned disk consumes only the space that it needs for its initial operrations, and grows with time according to demand. It is the fastest method to create a virtual disk because it creates a disk with just the header information.\n* `lazyZeroedThick` - A lazy zeroed thick disk has all space allocated at the time of its creation. Data remaining on the physical device is not erased during creation, but is zeroed out on demand later on first write from the virtual machine.\n* `eagerZeroedThick` - An eager zeroed thick disk has all space allocated and wiped clean of any previous contents on the physical media at creation time. Such disks may take longer time during creation compared to other disk formats.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"unit_number": {
				Description: "Unit number of the disk on its controller.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"uuid": {
				Description: "UUID assigned by vCenter to every disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vdisk_id": {
				Description: "Identity of the virtual disk object as the first class entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vendor": {
				Description: "Vendor of the storage device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"virtual_disk_path": {
				Description: "Path of the virtual disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_identity": {
				Description: "Identity of the virtual machine where the virtual disk is created.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"compatibility_mode": {
						Description: "Compatibility mode of the raw disk mapping (RDM).\n* `notApplicable` - Value specified for any disk which is not of raw device mapping type.\n* `physicalMode` - A disk device backed by a physical compatibility mode raw disk mapping cannot use disk modes, and commands are passed straight through to the LUN indicated by the raw disk mapping.\n* `virtualMode` - A disk device backed by a virtual compatibility mode raw disk mapping can use disk modes.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"controller_key": {
						Description: "Key of the controller on which the disk is created.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"datastore": {
						Description: "A reference to a virtualizationVmwareDatastore resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"device_name": {
						Description: "Host-specific device the LUN is being accessed through. If the target LUN is not available on the host then it is empty. For example, this could happen if it has accidentally been masked out.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"disk_mode": {
						Description: "Persistence mode of the virtual disk. For RDM disks, it is only supported when the raw disk mapping is using virtual compatibility mode.\n* `persistent` - Changes are immediately and permanently written to the virtual disk.\n* `independent_persistent` - Changes are immediately and permanently written to the virtual disk and not affected by snapshots.\n* `independent_nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off and not affected by snapshots.\n* `nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off.\n* `undoable` - Changes to virtual disk are made to a redo log and has the option to commit or undo.\n* `append` - Changes are appended to the redo log and can be revoked by removing the undo log.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"disk_type": {
						Description: "Specifies whether the virtual disk is a RDM or a flat disk.\n* `flatDisk` - Specifies that it is a flat disk.\n* `rdmDisk` - Specifies that it is a raw device mapping disk.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"key": {
						Description: "The internally assigned key of this disk. This entity is not manipulated by users.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"limit": {
						Description: "The utilization of a virtual machine will not exceed this limit, even if there are available resources. Used to ensure a consistent performance of virtual machines independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). The unit is number of I/O per second.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"lun_uuid": {
						Description: "Unique identifier of the LUN accessed by the raw disk mapping (RDM).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Name of the storage disk. Name must be unique within a Datastore.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"serial": {
						Description: "Serial ID of the storage device.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"shares": {
						Description: "Disk shares used for resource scheduling.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"level": {
									Description: "The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. If the shares value does not map to a predefined size, then the level is set as custom.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"shares": {
									Description: "The number of shares allocated. It is used to determine resource allocation in case of resource contention. Set if level is custom. If level is not set to custom, this value is ignored. Therefore, only shares with custom values can be compared.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"sharing": {
						Description: "Sharing mode of the virtual disk.\n* `sharingNone` - The virtual disk is not shared.\n* `sharingMultiWriter` - The virtual disk is shared between multiple virtual machines.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"size": {
						Description: "Disk size represented in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"storage_allocation_type": {
						Description: "Allocation type for the virtual disk.\n* `notApplicable` - Value specified for a disk for which storage allocation type is not applicable.\n* `thin` - A thin provisioned disk consumes only the space that it needs for its initial operrations, and grows with time according to demand. It is the fastest method to create a virtual disk because it creates a disk with just the header information.\n* `lazyZeroedThick` - A lazy zeroed thick disk has all space allocated at the time of its creation. Data remaining on the physical device is not erased during creation, but is zeroed out on demand later on first write from the virtual machine.\n* `eagerZeroedThick` - An eager zeroed thick disk has all space allocated and wiped clean of any previous contents on the physical media at creation time. Such disks may take longer time during creation compared to other disk formats.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"unit_number": {
						Description: "Unit number of the disk on its controller.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"uuid": {
						Description: "UUID assigned by vCenter to every disk.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vdisk_id": {
						Description: "Identity of the virtual disk object as the first class entity.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vendor": {
						Description: "Vendor of the storage device.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"virtual_disk_path": {
						Description: "Path of the virtual disk.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"virtual_machine": {
						Description: "A reference to a virtualizationVmwareVirtualMachine resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"vm_identity": {
						Description: "Identity of the virtual machine where the virtual disk is created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceVirtualizationVmwareVirtualDiskRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VirtualizationVmwareVirtualDisk{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("compatibility_mode"); ok {
		x := (v.(string))
		o.SetCompatibilityMode(x)
	}
	if v, ok := d.GetOk("controller_key"); ok {
		x := int64(v.(int))
		o.SetControllerKey(x)
	}
	if v, ok := d.GetOk("device_name"); ok {
		x := (v.(string))
		o.SetDeviceName(x)
	}
	if v, ok := d.GetOk("disk_mode"); ok {
		x := (v.(string))
		o.SetDiskMode(x)
	}
	if v, ok := d.GetOk("disk_type"); ok {
		x := (v.(string))
		o.SetDiskType(x)
	}
	if v, ok := d.GetOk("key"); ok {
		x := int64(v.(int))
		o.SetKey(x)
	}
	if v, ok := d.GetOk("limit"); ok {
		x := int64(v.(int))
		o.SetLimit(x)
	}
	if v, ok := d.GetOk("lun_uuid"); ok {
		x := (v.(string))
		o.SetLunUuid(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("sharing"); ok {
		x := (v.(string))
		o.SetSharing(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("storage_allocation_type"); ok {
		x := (v.(string))
		o.SetStorageAllocationType(x)
	}
	if v, ok := d.GetOk("unit_number"); ok {
		x := int64(v.(int))
		o.SetUnitNumber(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}
	if v, ok := d.GetOk("vdisk_id"); ok {
		x := (v.(string))
		o.SetVdiskId(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("virtual_disk_path"); ok {
		x := (v.(string))
		o.SetVirtualDiskPath(x)
	}
	if v, ok := d.GetOk("vm_identity"); ok {
		x := (v.(string))
		o.SetVmIdentity(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VirtualizationVmwareVirtualDisk object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVmwareVirtualDiskList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VirtualizationVmwareVirtualDisk: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VirtualizationVmwareVirtualDiskList.GetCount()
	var i int32
	var virtualizationVmwareVirtualDiskResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVmwareVirtualDiskList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VirtualizationVmwareVirtualDisk: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VirtualizationVmwareVirtualDiskList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VirtualizationVmwareVirtualDisk data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["compatibility_mode"] = (s.GetCompatibilityMode())
				temp["controller_key"] = (s.GetControllerKey())

				temp["datastore"] = flattenMapVirtualizationVmwareDatastoreRelationship(s.GetDatastore(), d)
				temp["device_name"] = (s.GetDeviceName())
				temp["disk_mode"] = (s.GetDiskMode())
				temp["disk_type"] = (s.GetDiskType())
				temp["key"] = (s.GetKey())
				temp["limit"] = (s.GetLimit())
				temp["lun_uuid"] = (s.GetLunUuid())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["serial"] = (s.GetSerial())

				temp["shares"] = flattenMapVirtualizationVmwareSharesInfo(s.GetShares(), d)
				temp["sharing"] = (s.GetSharing())
				temp["size"] = (s.GetSize())
				temp["storage_allocation_type"] = (s.GetStorageAllocationType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["unit_number"] = (s.GetUnitNumber())
				temp["uuid"] = (s.GetUuid())
				temp["vdisk_id"] = (s.GetVdiskId())
				temp["vendor"] = (s.GetVendor())
				temp["virtual_disk_path"] = (s.GetVirtualDiskPath())

				temp["virtual_machine"] = flattenMapVirtualizationVmwareVirtualMachineRelationship(s.GetVirtualMachine(), d)
				temp["vm_identity"] = (s.GetVmIdentity())
				virtualizationVmwareVirtualDiskResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(virtualizationVmwareVirtualDiskResults))
	if err := d.Set("results", virtualizationVmwareVirtualDiskResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(virtualizationVmwareVirtualDiskResults[0]["moid"].(string))
	return de
}
