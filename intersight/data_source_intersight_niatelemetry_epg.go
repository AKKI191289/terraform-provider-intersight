package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryEpg() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryEpgRead,
		Schema: map[string]*schema.Schema{
			"azure_pack_count": {
				Description: "Azure Pack NAT with ASA feature usage.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dn": {
				Description: "Dn value for the End Point Groups present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"epg_delimiter_count": {
				Description: "Number of  objects with delimiter value present in EPG Delimiter attribute.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fc_npv_count": {
				Description: "Number of ports with FC path attribute of type FC.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fcoe_count": {
				Description: "Number of FCoE per End Point Group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fv_rs_dom_att_count": {
				Description: "Number of FvRsDomAtt objects per End Point Group with VMware configuration.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"intra_epg_dvs_bm_count": {
				Description: "Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"intra_epg_hyperv": {
				Description: "Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_attr_based": {
				Description: "Gets the state of End Point Groups with isAttrBasedEPg value as configured.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"microsegmentation": {
				Description: "Gets the state of End Point Groups where microsegmentation is present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"microsoft_useg_count": {
				Description: "Number of FvRsDomAtt objects per End Point Group with Microsoft configuration.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name value for the End Point Groups present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"orchsl_dev_vip_cfg_count": {
				Description: "Number of objects with Simplified Service Graph Integration with Windows Azure Pack.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"record_type": {
				Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"record_version": {
				Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_name": {
				Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"useg_hyperv_count": {
				Description: "Logical Operators for attribute based microsegmentation in a hypervisor.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"azure_pack_count": {
						Description: "Azure Pack NAT with ASA feature usage.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"dn": {
						Description: "Dn value for the End Point Groups present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"epg_delimiter_count": {
						Description: "Number of  objects with delimiter value present in EPG Delimiter attribute.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"fc_npv_count": {
						Description: "Number of ports with FC path attribute of type FC.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"fcoe_count": {
						Description: "Number of FCoE per End Point Group.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"fv_rs_dom_att_count": {
						Description: "Number of FvRsDomAtt objects per End Point Group with VMware configuration.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"intra_epg_dvs_bm_count": {
						Description: "Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"intra_epg_hyperv": {
						Description: "Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"is_attr_based": {
						Description: "Gets the state of End Point Groups with isAttrBasedEPg value as configured.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"microsegmentation": {
						Description: "Gets the state of End Point Groups where microsegmentation is present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"microsoft_useg_count": {
						Description: "Number of FvRsDomAtt objects per End Point Group with Microsoft configuration.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Name value for the End Point Groups present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"orchsl_dev_vip_cfg_count": {
						Description: "Number of objects with Simplified Service Graph Integration with Windows Azure Pack.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"record_type": {
						Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"record_version": {
						Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
						Type:        schema.TypeString,
						Optional:    true,
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
					"site_name": {
						Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.",
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
					"useg_hyperv_count": {
						Description: "Logical Operators for attribute based microsegmentation in a hypervisor.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiatelemetryEpgRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryEpg{}
	if v, ok := d.GetOk("azure_pack_count"); ok {
		x := int64(v.(int))
		o.SetAzurePackCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("epg_delimiter_count"); ok {
		x := int64(v.(int))
		o.SetEpgDelimiterCount(x)
	}
	if v, ok := d.GetOk("fc_npv_count"); ok {
		x := int64(v.(int))
		o.SetFcNpvCount(x)
	}
	if v, ok := d.GetOk("fcoe_count"); ok {
		x := (v.(string))
		o.SetFcoeCount(x)
	}
	if v, ok := d.GetOk("fv_rs_dom_att_count"); ok {
		x := int64(v.(int))
		o.SetFvRsDomAttCount(x)
	}
	if v, ok := d.GetOk("intra_epg_dvs_bm_count"); ok {
		x := int64(v.(int))
		o.SetIntraEpgDvsBmCount(x)
	}
	if v, ok := d.GetOk("intra_epg_hyperv"); ok {
		x := (v.(string))
		o.SetIntraEpgHyperv(x)
	}
	if v, ok := d.GetOk("is_attr_based"); ok {
		x := (v.(string))
		o.SetIsAttrBased(x)
	}
	if v, ok := d.GetOk("microsegmentation"); ok {
		x := (v.(string))
		o.SetMicrosegmentation(x)
	}
	if v, ok := d.GetOk("microsoft_useg_count"); ok {
		x := int64(v.(int))
		o.SetMicrosoftUsegCount(x)
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
	if v, ok := d.GetOk("orchsl_dev_vip_cfg_count"); ok {
		x := int64(v.(int))
		o.SetOrchslDevVipCfgCount(x)
	}
	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}
	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("useg_hyperv_count"); ok {
		x := int64(v.(int))
		o.SetUsegHypervCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryEpg object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryEpgList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryEpg: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryEpgList.GetCount()
	var i int32
	var niatelemetryEpgResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryEpgList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryEpg: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryEpgList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryEpg data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["azure_pack_count"] = (s.GetAzurePackCount())
				temp["class_id"] = (s.GetClassId())
				temp["dn"] = (s.GetDn())
				temp["epg_delimiter_count"] = (s.GetEpgDelimiterCount())
				temp["fc_npv_count"] = (s.GetFcNpvCount())
				temp["fcoe_count"] = (s.GetFcoeCount())
				temp["fv_rs_dom_att_count"] = (s.GetFvRsDomAttCount())
				temp["intra_epg_dvs_bm_count"] = (s.GetIntraEpgDvsBmCount())
				temp["intra_epg_hyperv"] = (s.GetIntraEpgHyperv())
				temp["is_attr_based"] = (s.GetIsAttrBased())
				temp["microsegmentation"] = (s.GetMicrosegmentation())
				temp["microsoft_useg_count"] = (s.GetMicrosoftUsegCount())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["orchsl_dev_vip_cfg_count"] = (s.GetOrchslDevVipCfgCount())
				temp["record_type"] = (s.GetRecordType())
				temp["record_version"] = (s.GetRecordVersion())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["site_name"] = (s.GetSiteName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["useg_hyperv_count"] = (s.GetUsegHypervCount())
				niatelemetryEpgResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryEpgResults))
	if err := d.Set("results", niatelemetryEpgResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryEpgResults[0]["moid"].(string))
	return de
}
