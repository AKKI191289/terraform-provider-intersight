package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIaasServiceRequest() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIaasServiceRequestRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"duration": {
				Description: "Service request duration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"initiating_user": {
				Description: "Service Request initiating user.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_end_time": {
				Description: "Service request end time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_id": {
				Description: "Service request id of an SR.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_start_time": {
				Description: "Service request start time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_type": {
				Description: "Service request type of an SR.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "UCSD service request status.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"workflow_name": {
				Description: "Executed workflow name for an SR.",
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
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"duration": {
						Description: "Service request duration.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"initiating_user": {
						Description: "Service Request initiating user.",
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
					"request_end_time": {
						Description: "Service request end time.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"request_id": {
						Description: "Service request id of an SR.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"request_start_time": {
						Description: "Service request start time.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"request_type": {
						Description: "Service request type of an SR.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"status": {
						Description: "UCSD service request status.",
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
					"workflow_name": {
						Description: "Executed workflow name for an SR.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"workflow_steps": {
						Type:     schema.TypeList,
						Optional: true,
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
								"completed_time": {
									Description: "Completed time of the workflow step.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the workflow step.",
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
								"status": {
									Description: "Status of the workflow step.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"status_message": {
									Description: "Status message of the workflow step.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceIaasServiceRequestRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IaasServiceRequest{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("duration"); ok {
		x := (v.(string))
		o.SetDuration(x)
	}
	if v, ok := d.GetOk("initiating_user"); ok {
		x := (v.(string))
		o.SetInitiatingUser(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("request_end_time"); ok {
		x := (v.(string))
		o.SetRequestEndTime(x)
	}
	if v, ok := d.GetOk("request_id"); ok {
		x := (v.(string))
		o.SetRequestId(x)
	}
	if v, ok := d.GetOk("request_start_time"); ok {
		x := (v.(string))
		o.SetRequestStartTime(x)
	}
	if v, ok := d.GetOk("request_type"); ok {
		x := (v.(string))
		o.SetRequestType(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("workflow_name"); ok {
		x := (v.(string))
		o.SetWorkflowName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IaasServiceRequest object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IaasApi.GetIaasServiceRequestList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IaasServiceRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IaasServiceRequestList.GetCount()
	var i int32
	var iaasServiceRequestResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IaasApi.GetIaasServiceRequestList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IaasServiceRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IaasServiceRequestList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IaasServiceRequest data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["duration"] = (s.GetDuration())
				temp["initiating_user"] = (s.GetInitiatingUser())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["request_end_time"] = (s.GetRequestEndTime())
				temp["request_id"] = (s.GetRequestId())
				temp["request_start_time"] = (s.GetRequestStartTime())
				temp["request_type"] = (s.GetRequestType())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["workflow_name"] = (s.GetWorkflowName())

				temp["workflow_steps"] = flattenListIaasWorkflowSteps(s.GetWorkflowSteps(), d)
				iaasServiceRequestResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iaasServiceRequestResults))
	if err := d.Set("results", iaasServiceRequestResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iaasServiceRequestResults[0]["moid"].(string))
	return de
}
