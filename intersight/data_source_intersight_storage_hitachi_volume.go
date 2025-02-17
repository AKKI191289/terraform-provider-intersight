package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageHitachiVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageHitachiVolumeRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"clpr_id": {
				Description: "CLPR (Cache Logical Partition) number of this volume.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"data_reduction_mode": {
				Description: "Setting of the capacity saving function (dedupe and compression).\n* `N/A` - The capacity saving function is not available.\n* `Compression` - The capacity saving function (compression) is enabled.\n* `Compression Deduplication` - The capacity saving function (compression and deduplication) is enabled.\n* `Disabled` - The capacity saving function (compression and deduplication) is disabled.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"data_reduction_status": {
				Description: "Status of the capacity saving function.\n* `N/A` - The capacity saving function is not available.\n* `Enabled` - The capacity saving function is enabled.\n* `Disabled` - The capacity saving function is disabled.\n* `Enabling` - The capacity saving function is being enabled.\n* `Rehydrating` - The capacity saving function is being disabled.\n* `Deleting` - The volumes for which the capacity saving function is enabled are being deleted.\n* `Failed` - An attempt to enable the capacity saving function failed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Short description about the volume.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drive_type": {
				Description: "Code indicating the drive type of the drive belonging to the volume.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"emulation_type": {
				Description: "The volume emulation type or the volume status information.\n* `N/A` - Not available.\n* `NOT DEFINED` - The volume is not implemented.\n* `DEFINING` - The volume is being created.\n* `REMOVING` - The volume is being removed.\n* `OPEN-V` - To be provided by Hitachi.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_full_allocation_enabled": {
				Description: "Whether pages are reserved by the FullAllocation functionality.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"label": {
				Description: "Label of the volume, as configured in the storage array.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"naa_id": {
				Description: "NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Named entity of the volume.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_of_paths": {
				Description: "Number of paths set for the volume.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pool_id": {
				Description: "ID of the pool with which the volume is associated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"raid_level": {
				Description: "RAID level for the volume.\n* `N/A` - RAID level is unknown or multiple RAID levels are being used.\n* `RAID1` - RAID1.\n* `RAID5` - RAID5.\n* `RAID6` - RAID6.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"raid_type": {
				Description: "RAID type drive configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"size": {
				Description: "User provisioned volume size. It is the size exposed to host.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status information of the volume.\n* `N/A` - The volume status is not available.\n* `NML` - The volume is in normal status.\n* `BLK` - The volume is in blocked state.\n* `BSY` - The volume status is being changed.\n* `Unknown` - The volume status is unknown (not supported).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"array": {
						Description: "A reference to a storageHitachiArray resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"attributes": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"clpr_id": {
						Description: "CLPR (Cache Logical Partition) number of this volume.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"data_reduction_mode": {
						Description: "Setting of the capacity saving function (dedupe and compression).\n* `N/A` - The capacity saving function is not available.\n* `Compression` - The capacity saving function (compression) is enabled.\n* `Compression Deduplication` - The capacity saving function (compression and deduplication) is enabled.\n* `Disabled` - The capacity saving function (compression and deduplication) is disabled.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"data_reduction_status": {
						Description: "Status of the capacity saving function.\n* `N/A` - The capacity saving function is not available.\n* `Enabled` - The capacity saving function is enabled.\n* `Disabled` - The capacity saving function is disabled.\n* `Enabling` - The capacity saving function is being enabled.\n* `Rehydrating` - The capacity saving function is being disabled.\n* `Deleting` - The volumes for which the capacity saving function is enabled are being deleted.\n* `Failed` - An attempt to enable the capacity saving function failed.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"description": {
						Description: "Short description about the volume.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"drive_type": {
						Description: "Code indicating the drive type of the drive belonging to the volume.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"emulation_type": {
						Description: "The volume emulation type or the volume status information.\n* `N/A` - Not available.\n* `NOT DEFINED` - The volume is not implemented.\n* `DEFINING` - The volume is being created.\n* `REMOVING` - The volume is being removed.\n* `OPEN-V` - To be provided by Hitachi.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"is_full_allocation_enabled": {
						Description: "Whether pages are reserved by the FullAllocation functionality.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"label": {
						Description: "Label of the volume, as configured in the storage array.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"naa_id": {
						Description: "NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Named entity of the volume.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"num_of_paths": {
						Description: "Number of paths set for the volume.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"parity_group_ids": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"parity_groups": {
						Description: "An array of relationships to storageHitachiParityGroup resources.",
						Type:        schema.TypeList,
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
					"pool": {
						Description: "A reference to a storageHitachiPool resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"pool_id": {
						Description: "ID of the pool with which the volume is associated.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"raid_level": {
						Description: "RAID level for the volume.\n* `N/A` - RAID level is unknown or multiple RAID levels are being used.\n* `RAID1` - RAID1.\n* `RAID5` - RAID5.\n* `RAID6` - RAID6.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"raid_type": {
						Description: "RAID type drive configuration.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"size": {
						Description: "User provisioned volume size. It is the size exposed to host.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"status": {
						Description: "Status information of the volume.\n* `N/A` - The volume status is not available.\n* `NML` - The volume is in normal status.\n* `BLK` - The volume is in blocked state.\n* `BSY` - The volume status is being changed.\n* `Unknown` - The volume status is unknown (not supported).",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"storage_utilization": {
						Description: "Storage utilization of volume entity in storage array.",
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
								"available": {
									Description: "Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"capacity_utilization": {
									Description: "Percentage of used capacity.",
									Type:        schema.TypeFloat,
									Optional:    true,
									Computed:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"free": {
									Description: "Unused space available for applications to consume, represented in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"total": {
									Description: "Total storage capacity, represented in bytes. It is set by the component manufacturer.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"used": {
									Description: "Used or consumed storage capacity, represented in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
							},
						},
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageHitachiVolumeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageHitachiVolume{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("clpr_id"); ok {
		x := int64(v.(int))
		o.SetClprId(x)
	}
	if v, ok := d.GetOk("data_reduction_mode"); ok {
		x := (v.(string))
		o.SetDataReductionMode(x)
	}
	if v, ok := d.GetOk("data_reduction_status"); ok {
		x := (v.(string))
		o.SetDataReductionStatus(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("drive_type"); ok {
		x := (v.(string))
		o.SetDriveType(x)
	}
	if v, ok := d.GetOk("emulation_type"); ok {
		x := (v.(string))
		o.SetEmulationType(x)
	}
	if v, ok := d.GetOk("is_full_allocation_enabled"); ok {
		x := (v.(bool))
		o.SetIsFullAllocationEnabled(x)
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("naa_id"); ok {
		x := (v.(string))
		o.SetNaaId(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("num_of_paths"); ok {
		x := int64(v.(int))
		o.SetNumOfPaths(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pool_id"); ok {
		x := (v.(string))
		o.SetPoolId(x)
	}
	if v, ok := d.GetOk("raid_level"); ok {
		x := (v.(string))
		o.SetRaidLevel(x)
	}
	if v, ok := d.GetOk("raid_type"); ok {
		x := (v.(string))
		o.SetRaidType(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageHitachiVolume object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiVolumeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageHitachiVolume: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageHitachiVolumeList.GetCount()
	var i int32
	var storageHitachiVolumeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiVolumeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageHitachiVolume: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageHitachiVolumeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageHitachiVolume data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["array"] = flattenMapStorageHitachiArrayRelationship(s.GetArray(), d)
				temp["attributes"] = (s.GetAttributes())
				temp["class_id"] = (s.GetClassId())
				temp["clpr_id"] = (s.GetClprId())
				temp["data_reduction_mode"] = (s.GetDataReductionMode())
				temp["data_reduction_status"] = (s.GetDataReductionStatus())
				temp["description"] = (s.GetDescription())
				temp["drive_type"] = (s.GetDriveType())
				temp["emulation_type"] = (s.GetEmulationType())
				temp["is_full_allocation_enabled"] = (s.GetIsFullAllocationEnabled())
				temp["label"] = (s.GetLabel())
				temp["moid"] = (s.GetMoid())
				temp["naa_id"] = (s.GetNaaId())
				temp["name"] = (s.GetName())
				temp["num_of_paths"] = (s.GetNumOfPaths())
				temp["object_type"] = (s.GetObjectType())
				temp["parity_group_ids"] = (s.GetParityGroupIds())

				temp["parity_groups"] = flattenListStorageHitachiParityGroupRelationship(s.GetParityGroups(), d)

				temp["pool"] = flattenMapStorageHitachiPoolRelationship(s.GetPool(), d)
				temp["pool_id"] = (s.GetPoolId())
				temp["raid_level"] = (s.GetRaidLevel())
				temp["raid_type"] = (s.GetRaidType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["size"] = (s.GetSize())
				temp["status"] = (s.GetStatus())

				temp["storage_utilization"] = flattenMapStorageBaseCapacity(s.GetStorageUtilization(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				storageHitachiVolumeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageHitachiVolumeResults))
	if err := d.Set("results", storageHitachiVolumeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageHitachiVolumeResults[0]["moid"].(string))
	return de
}
