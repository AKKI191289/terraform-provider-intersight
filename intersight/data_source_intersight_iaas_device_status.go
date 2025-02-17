package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIaasDeviceStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIaasDeviceStatusRead,
		Schema: map[string]*schema.Schema{
			"account_name": {
				Description: "The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"account_type": {
				Description: "The UCSD Infra Account type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Description: "The UCSD Infra Account category.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"claim_status": {
				Description: "Describes if the device is claimed in Intersight or not.\n* `Unknown` - If UCS Director managed account claim status information is unknown.\n* `Yes` - If UCS Director managed account is claimed in Intersight.\n* `No` - If UCS Director managed account is not claimed in Intersight.\n* `Not Applicable` - If UCS Director managed account is not capable of providing claim status information.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connection_status": {
				Description: "Describes about the connection status between the UCSD and the actual end device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_model": {
				Description: "Describes about the device model.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_vendor": {
				Description: "Describes about the device vendor/manufacturer of the device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_version": {
				Description: "Describes about the current firmware version running on the device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ip_address": {
				Description: "The IPAddress of the device.",
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
			"pod": {
				Description: "Describes about the pod to which this device belongs to in UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pod_type": {
				Description: "Describes about the podType of Pod to which this device belongs to in UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account_name": {
					Description: "The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD.",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
					"account_type": {
						Description: "The UCSD Infra Account type.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"category": {
						Description: "The UCSD Infra Account category.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"claim_status": {
						Description: "Describes if the device is claimed in Intersight or not.\n* `Unknown` - If UCS Director managed account claim status information is unknown.\n* `Yes` - If UCS Director managed account is claimed in Intersight.\n* `No` - If UCS Director managed account is not claimed in Intersight.\n* `Not Applicable` - If UCS Director managed account is not capable of providing claim status information.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"connection_status": {
						Description: "Describes about the connection status between the UCSD and the actual end device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_model": {
						Description: "Describes about the device model.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_vendor": {
						Description: "Describes about the device vendor/manufacturer of the device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_version": {
						Description: "Describes about the current firmware version running on the device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"guid": {
						Description: "A reference to a iaasUcsdInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"ip_address": {
						Description: "The IPAddress of the device.",
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
					"pod": {
						Description: "Describes about the pod to which this device belongs to in UCSD.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"pod_type": {
						Description: "Describes about the podType of Pod to which this device belongs to in UCSD.",
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

func dataSourceIaasDeviceStatusRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IaasDeviceStatus{}
	if v, ok := d.GetOk("account_name"); ok {
		x := (v.(string))
		o.SetAccountName(x)
	}
	if v, ok := d.GetOk("account_type"); ok {
		x := (v.(string))
		o.SetAccountType(x)
	}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("claim_status"); ok {
		x := (v.(string))
		o.SetClaimStatus(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.SetConnectionStatus(x)
	}
	if v, ok := d.GetOk("device_model"); ok {
		x := (v.(string))
		o.SetDeviceModel(x)
	}
	if v, ok := d.GetOk("device_vendor"); ok {
		x := (v.(string))
		o.SetDeviceVendor(x)
	}
	if v, ok := d.GetOk("device_version"); ok {
		x := (v.(string))
		o.SetDeviceVersion(x)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.SetIpAddress(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pod"); ok {
		x := (v.(string))
		o.SetPod(x)
	}
	if v, ok := d.GetOk("pod_type"); ok {
		x := (v.(string))
		o.SetPodType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IaasDeviceStatus object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IaasApi.GetIaasDeviceStatusList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IaasDeviceStatus: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IaasDeviceStatusList.GetCount()
	var i int32
	var iaasDeviceStatusResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IaasApi.GetIaasDeviceStatusList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IaasDeviceStatus: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IaasDeviceStatusList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IaasDeviceStatus data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_name"] = (s.GetAccountName())
				temp["account_type"] = (s.GetAccountType())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["category"] = (s.GetCategory())
				temp["claim_status"] = (s.GetClaimStatus())
				temp["class_id"] = (s.GetClassId())
				temp["connection_status"] = (s.GetConnectionStatus())
				temp["device_model"] = (s.GetDeviceModel())
				temp["device_vendor"] = (s.GetDeviceVendor())
				temp["device_version"] = (s.GetDeviceVersion())

				temp["guid"] = flattenMapIaasUcsdInfoRelationship(s.GetGuid(), d)
				temp["ip_address"] = (s.GetIpAddress())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["pod"] = (s.GetPod())
				temp["pod_type"] = (s.GetPodType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				iaasDeviceStatusResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iaasDeviceStatusResults))
	if err := d.Set("results", iaasDeviceStatusResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iaasDeviceStatusResults[0]["moid"].(string))
	return de
}
