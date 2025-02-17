package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCondAlarm() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCondAlarmRead,
		Schema: map[string]*schema.Schema{
			"acknowledge": {
				Description: "Alarm acknowledgment state. Default value is None.\n* `None` - The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count.\n* `Acknowledge` - The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acknowledge_by": {
				Description: "User who acknowledged the alarm.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"acknowledge_time": {
				Description: "Time at which the alarm was acknowledged by the user.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"affected_mo_display_name": {
				Description: "Display name of the affected object on which the alarm is raised. The name uniquely identifies an alarm target such that it can be distinguished from similar objects managed by Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"affected_mo_id": {
				Description: "MoId of the affected object from the managed system's point of view.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"affected_mo_type": {
				Description: "Managed system affected object type. For example Adaptor, FI, CIMC.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ancestor_mo_id": {
				Description: "Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ancestor_mo_type": {
				Description: "Parent MO type of the fault from managed system. For example, Blade for adaptor fault.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"code": {
				Description: "A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"creation_time": {
				Description: "The time the alarm was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "A longer description of the alarm than the name. The description contains details of the component reporting the issue.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_transition_time": {
				Description: "The time the alarm last had a change in severity.",
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
			"ms_affected_object": {
				Description: "A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"orig_severity": {
				Description: "The original severity when the alarm was first created.\n* `None` - The Enum value None represents that there is no severity.\n* `Info` - The Enum value Info represents the Informational level of severity.\n* `Critical` - The Enum value Critical represents the Critical level of severity.\n* `Warning` - The Enum value Warning represents the Warning level of severity.\n* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Description: "The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared.\n* `None` - The Enum value None represents that there is no severity.\n* `Info` - The Enum value Info represents the Informational level of severity.\n* `Critical` - The Enum value Critical represents the Critical level of severity.\n* `Warning` - The Enum value Warning represents the Warning level of severity.\n* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"acknowledge": {
					Description: "Alarm acknowledgment state. Default value is None.\n* `None` - The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count.\n* `Acknowledge` - The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count.",
					Type:        schema.TypeString,
					Optional:    true,
				},
					"acknowledge_by": {
						Description: "User who acknowledged the alarm.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"acknowledge_time": {
						Description: "Time at which the alarm was acknowledged by the user.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"affected_mo": {
						Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"affected_mo_display_name": {
						Description: "Display name of the affected object on which the alarm is raised. The name uniquely identifies an alarm target such that it can be distinguished from similar objects managed by Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"affected_mo_id": {
						Description: "MoId of the affected object from the managed system's point of view.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"affected_mo_type": {
						Description: "Managed system affected object type. For example Adaptor, FI, CIMC.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ancestor_mo_id": {
						Description: "Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ancestor_mo_type": {
						Description: "Parent MO type of the fault from managed system. For example, Blade for adaptor fault.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"code": {
						Description: "A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"creation_time": {
						Description: "The time the alarm was created.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"description": {
						Description: "A longer description of the alarm than the name. The description contains details of the component reporting the issue.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"last_transition_time": {
						Description: "The time the alarm last had a change in severity.",
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
					"ms_affected_object": {
						Description: "A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"orig_severity": {
						Description: "The original severity when the alarm was first created.\n* `None` - The Enum value None represents that there is no severity.\n* `Info` - The Enum value Info represents the Informational level of severity.\n* `Critical` - The Enum value Critical represents the Critical level of severity.\n* `Warning` - The Enum value Warning represents the Warning level of severity.\n* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.",
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
					"severity": {
						Description: "The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared.\n* `None` - The Enum value None represents that there is no severity.\n* `Info` - The Enum value Info represents the Informational level of severity.\n* `Critical` - The Enum value Critical represents the Critical level of severity.\n* `Warning` - The Enum value Warning represents the Warning level of severity.\n* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.",
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

func dataSourceCondAlarmRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CondAlarm{}
	if v, ok := d.GetOk("acknowledge"); ok {
		x := (v.(string))
		o.SetAcknowledge(x)
	}
	if v, ok := d.GetOk("acknowledge_by"); ok {
		x := (v.(string))
		o.SetAcknowledgeBy(x)
	}
	if v, ok := d.GetOk("acknowledge_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetAcknowledgeTime(x)
	}
	if v, ok := d.GetOk("affected_mo_display_name"); ok {
		x := (v.(string))
		o.SetAffectedMoDisplayName(x)
	}
	if v, ok := d.GetOk("affected_mo_id"); ok {
		x := (v.(string))
		o.SetAffectedMoId(x)
	}
	if v, ok := d.GetOk("affected_mo_type"); ok {
		x := (v.(string))
		o.SetAffectedMoType(x)
	}
	if v, ok := d.GetOk("ancestor_mo_id"); ok {
		x := (v.(string))
		o.SetAncestorMoId(x)
	}
	if v, ok := d.GetOk("ancestor_mo_type"); ok {
		x := (v.(string))
		o.SetAncestorMoType(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("code"); ok {
		x := (v.(string))
		o.SetCode(x)
	}
	if v, ok := d.GetOk("creation_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreationTime(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("last_transition_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastTransitionTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("ms_affected_object"); ok {
		x := (v.(string))
		o.SetMsAffectedObject(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("orig_severity"); ok {
		x := (v.(string))
		o.SetOrigSeverity(x)
	}
	if v, ok := d.GetOk("severity"); ok {
		x := (v.(string))
		o.SetSeverity(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CondAlarm object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CondApi.GetCondAlarmList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CondAlarm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CondAlarmList.GetCount()
	var i int32
	var condAlarmResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CondApi.GetCondAlarmList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CondAlarm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CondAlarmList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CondAlarm data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["acknowledge"] = (s.GetAcknowledge())
				temp["acknowledge_by"] = (s.GetAcknowledgeBy())

				temp["acknowledge_time"] = (s.GetAcknowledgeTime()).String()
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["affected_mo"] = flattenMapMoBaseMoRelationship(s.GetAffectedMo(), d)
				temp["affected_mo_display_name"] = (s.GetAffectedMoDisplayName())
				temp["affected_mo_id"] = (s.GetAffectedMoId())
				temp["affected_mo_type"] = (s.GetAffectedMoType())
				temp["ancestor_mo_id"] = (s.GetAncestorMoId())
				temp["ancestor_mo_type"] = (s.GetAncestorMoType())
				temp["class_id"] = (s.GetClassId())
				temp["code"] = (s.GetCode())

				temp["creation_time"] = (s.GetCreationTime()).String()
				temp["description"] = (s.GetDescription())

				temp["last_transition_time"] = (s.GetLastTransitionTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["ms_affected_object"] = (s.GetMsAffectedObject())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["orig_severity"] = (s.GetOrigSeverity())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["severity"] = (s.GetSeverity())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				condAlarmResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(condAlarmResults))
	if err := d.Set("results", condAlarmResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(condAlarmResults[0]["moid"].(string))
	return de
}
