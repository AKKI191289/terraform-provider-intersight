package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStoragePureSnapshotSchedule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStoragePureSnapshotScheduleRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"daily_limit": {
				Description: "Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"frequency": {
				Description: "Snapshot frequency. It is an interval at which snapshot is set to trigger on source array.\nExamples:\n    PT2H Snapshot is generated every 2 hours.\n    P4D, Snapshot is scheduled for every 4 days.\n    PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.",
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
			"name": {
				Description: "Name of the snapshot schedule.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"retention_time": {
				Description: "Duration to keep the snapshots on the source array.\nOnce this period expires, system deletes the snapshot automatically from the source array.\nExamples:\nP200D,  200 days.\nPT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"snapshot_expiry_time": {
				Description: "Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"snapshot_time": {
				Description: "Preferred time of the day to capture the snapshot. It is applicable only if the frequency is set for a day or more.\nFormat: hh:mm:ss\nExample: 08:30:00, Snapshot is set for 08:30 AM.",
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
						Description: "A reference to a storagePureArray resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"daily_limit": {
						Description: "Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"frequency": {
						Description: "Snapshot frequency. It is an interval at which snapshot is set to trigger on source array.\nExamples:\n    PT2H Snapshot is generated every 2 hours.\n    P4D, Snapshot is scheduled for every 4 days.\n    PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.",
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
					"name": {
						Description: "Name of the snapshot schedule.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"protection_group": {
						Description: "A reference to a storagePureProtectionGroup resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"retention_time": {
						Description: "Duration to keep the snapshots on the source array.\nOnce this period expires, system deletes the snapshot automatically from the source array.\nExamples:\nP200D,  200 days.\nPT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"snapshot_expiry_time": {
						Description: "Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"snapshot_time": {
						Description: "Preferred time of the day to capture the snapshot. It is applicable only if the frequency is set for a day or more.\nFormat: hh:mm:ss\nExample: 08:30:00, Snapshot is set for 08:30 AM.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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

func dataSourceStoragePureSnapshotScheduleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StoragePureSnapshotSchedule{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("daily_limit"); ok {
		x := int64(v.(int))
		o.SetDailyLimit(x)
	}
	if v, ok := d.GetOk("frequency"); ok {
		x := (v.(string))
		o.SetFrequency(x)
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
	if v, ok := d.GetOk("retention_time"); ok {
		x := (v.(string))
		o.SetRetentionTime(x)
	}
	if v, ok := d.GetOk("snapshot_expiry_time"); ok {
		x := (v.(string))
		o.SetSnapshotExpiryTime(x)
	}
	if v, ok := d.GetOk("snapshot_time"); ok {
		x := (v.(string))
		o.SetSnapshotTime(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StoragePureSnapshotSchedule object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStoragePureSnapshotScheduleList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StoragePureSnapshotSchedule: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StoragePureSnapshotScheduleList.GetCount()
	var i int32
	var storagePureSnapshotScheduleResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStoragePureSnapshotScheduleList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StoragePureSnapshotSchedule: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StoragePureSnapshotScheduleList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StoragePureSnapshotSchedule data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["array"] = flattenMapStoragePureArrayRelationship(s.GetArray(), d)
				temp["class_id"] = (s.GetClassId())
				temp["daily_limit"] = (s.GetDailyLimit())
				temp["frequency"] = (s.GetFrequency())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["protection_group"] = flattenMapStoragePureProtectionGroupRelationship(s.GetProtectionGroup(), d)

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["retention_time"] = (s.GetRetentionTime())
				temp["snapshot_expiry_time"] = (s.GetSnapshotExpiryTime())
				temp["snapshot_time"] = (s.GetSnapshotTime())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				storagePureSnapshotScheduleResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storagePureSnapshotScheduleResults))
	if err := d.Set("results", storagePureSnapshotScheduleResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storagePureSnapshotScheduleResults[0]["moid"].(string))
	return de
}
