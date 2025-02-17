package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryNiaInventoryDcnm() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNiaInventoryDcnmRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dev": {
				Description: "Returns the value of the dev Field.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"epld_image_count": {
				Description: "Number of EPLD images uploaded to DCNM.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ha_enabled": {
				Description: "Returns the value of the haEnabled field.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"ha_replication_status": {
				Description: "Returns the value of the haReplicationStatus field.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"install": {
				Description: "Returns the value of the install field.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_isn_configured": {
				Description: "Returns true if ISN is configured.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_media_controller": {
				Description: "Returns the value of the isMediaController field.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_smart_license_enabled": {
				Description: "Returns true if the Smart license is enabled and is in use.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_fabrics": {
				Description: "Returns total number of fabrics in DCNM set-up.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_fabrics_in_msd": {
				Description: "Returns the number of fabrics in msd.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_ingress_replication_fabrics": {
				Description: "Returns the number of fabrics that have ingress replication type.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_local_users": {
				Description: "Returns the number of local users other than admin user.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_msd": {
				Description: "Returns the number of MSD fabrics.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_svi_vrf_count": {
				Description: "Returns the number of svi interfaces configured for VRF vlans.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_trm_enabled_count": {
				Description: "Returns the number of links where TRM is enabled.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_upg_users": {
				Description: "Number of users who have upgrade privileges excluding the admin.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"nxos_image_count": {
				Description: "Number of NXOS images uploaded to DCNM.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "Serial number of device being inventoried. The serial number is unique per device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_name": {
				Description: "Name of fabric domain of the controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"underlay_peering_active_links_count": {
				Description: "Returns the number of underlay peering active links.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"upg_job_count": {
				Description: "Number of upgrade jobs configured on DCNM.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"upg_status": {
				Description: "Upgrade status of jobs created on DCNM.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "Returns the value of the version field.",
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
					"dev": {
						Description: "Returns the value of the dev Field.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"epld_image_count": {
						Description: "Number of EPLD images uploaded to DCNM.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ha_enabled": {
						Description: "Returns the value of the haEnabled field.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"ha_replication_status": {
						Description: "Returns the value of the haReplicationStatus field.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"install": {
						Description: "Returns the value of the install field.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"is_isn_configured": {
						Description: "Returns true if ISN is configured.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"is_media_controller": {
						Description: "Returns the value of the isMediaController field.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"is_smart_license_enabled": {
						Description: "Returns true if the Smart license is enabled and is in use.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"num_fabrics": {
						Description: "Returns total number of fabrics in DCNM set-up.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_fabrics_in_msd": {
						Description: "Returns the number of fabrics in msd.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_ingress_replication_fabrics": {
						Description: "Returns the number of fabrics that have ingress replication type.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_local_users": {
						Description: "Returns the number of local users other than admin user.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_msd": {
						Description: "Returns the number of MSD fabrics.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_svi_vrf_count": {
						Description: "Returns the number of svi interfaces configured for VRF vlans.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_trm_enabled_count": {
						Description: "Returns the number of links where TRM is enabled.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_upg_users": {
						Description: "Number of users who have upgrade privileges excluding the admin.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"nxos_image_count": {
						Description: "Number of NXOS images uploaded to DCNM.",
						Type:        schema.TypeInt,
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
					"serial": {
						Description: "Serial number of device being inventoried. The serial number is unique per device.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"site_name": {
						Description: "Name of fabric domain of the controller.",
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
					"underlay_peering_active_links_count": {
						Description: "Returns the number of underlay peering active links.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"upg_job_count": {
						Description: "Number of upgrade jobs configured on DCNM.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"upg_status": {
						Description: "Upgrade status of jobs created on DCNM.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Returns the value of the version field.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiatelemetryNiaInventoryDcnmRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNiaInventoryDcnm{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("dev"); ok {
		x := (v.(bool))
		o.SetDev(x)
	}
	if v, ok := d.GetOk("epld_image_count"); ok {
		x := int64(v.(int))
		o.SetEpldImageCount(x)
	}
	if v, ok := d.GetOk("ha_enabled"); ok {
		x := (v.(bool))
		o.SetHaEnabled(x)
	}
	if v, ok := d.GetOk("ha_replication_status"); ok {
		x := (v.(string))
		o.SetHaReplicationStatus(x)
	}
	if v, ok := d.GetOk("install"); ok {
		x := (v.(string))
		o.SetInstall(x)
	}
	if v, ok := d.GetOk("is_isn_configured"); ok {
		x := (v.(bool))
		o.SetIsIsnConfigured(x)
	}
	if v, ok := d.GetOk("is_media_controller"); ok {
		x := (v.(bool))
		o.SetIsMediaController(x)
	}
	if v, ok := d.GetOk("is_smart_license_enabled"); ok {
		x := (v.(bool))
		o.SetIsSmartLicenseEnabled(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("num_fabrics"); ok {
		x := int64(v.(int))
		o.SetNumFabrics(x)
	}
	if v, ok := d.GetOk("num_fabrics_in_msd"); ok {
		x := int64(v.(int))
		o.SetNumFabricsInMsd(x)
	}
	if v, ok := d.GetOk("num_ingress_replication_fabrics"); ok {
		x := int64(v.(int))
		o.SetNumIngressReplicationFabrics(x)
	}
	if v, ok := d.GetOk("num_local_users"); ok {
		x := int64(v.(int))
		o.SetNumLocalUsers(x)
	}
	if v, ok := d.GetOk("num_msd"); ok {
		x := int64(v.(int))
		o.SetNumMsd(x)
	}
	if v, ok := d.GetOk("num_svi_vrf_count"); ok {
		x := int64(v.(int))
		o.SetNumSviVrfCount(x)
	}
	if v, ok := d.GetOk("num_trm_enabled_count"); ok {
		x := int64(v.(int))
		o.SetNumTrmEnabledCount(x)
	}
	if v, ok := d.GetOk("num_upg_users"); ok {
		x := int64(v.(int))
		o.SetNumUpgUsers(x)
	}
	if v, ok := d.GetOk("nxos_image_count"); ok {
		x := int64(v.(int))
		o.SetNxosImageCount(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("underlay_peering_active_links_count"); ok {
		x := int64(v.(int))
		o.SetUnderlayPeeringActiveLinksCount(x)
	}
	if v, ok := d.GetOk("upg_job_count"); ok {
		x := int64(v.(int))
		o.SetUpgJobCount(x)
	}
	if v, ok := d.GetOk("upg_status"); ok {
		x := (v.(string))
		o.SetUpgStatus(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaInventoryDcnm object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryDcnmList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventoryDcnm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryNiaInventoryDcnmList.GetCount()
	var i int32
	var niatelemetryNiaInventoryDcnmResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryDcnmList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryNiaInventoryDcnm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryNiaInventoryDcnmList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryNiaInventoryDcnm data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["dev"] = (s.GetDev())
				temp["epld_image_count"] = (s.GetEpldImageCount())
				temp["ha_enabled"] = (s.GetHaEnabled())
				temp["ha_replication_status"] = (s.GetHaReplicationStatus())
				temp["install"] = (s.GetInstall())
				temp["is_isn_configured"] = (s.GetIsIsnConfigured())
				temp["is_media_controller"] = (s.GetIsMediaController())
				temp["is_smart_license_enabled"] = (s.GetIsSmartLicenseEnabled())
				temp["moid"] = (s.GetMoid())
				temp["num_fabrics"] = (s.GetNumFabrics())
				temp["num_fabrics_in_msd"] = (s.GetNumFabricsInMsd())
				temp["num_ingress_replication_fabrics"] = (s.GetNumIngressReplicationFabrics())
				temp["num_local_users"] = (s.GetNumLocalUsers())
				temp["num_msd"] = (s.GetNumMsd())
				temp["num_svi_vrf_count"] = (s.GetNumSviVrfCount())
				temp["num_trm_enabled_count"] = (s.GetNumTrmEnabledCount())
				temp["num_upg_users"] = (s.GetNumUpgUsers())
				temp["nxos_image_count"] = (s.GetNxosImageCount())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["serial"] = (s.GetSerial())
				temp["site_name"] = (s.GetSiteName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["underlay_peering_active_links_count"] = (s.GetUnderlayPeeringActiveLinksCount())
				temp["upg_job_count"] = (s.GetUpgJobCount())
				temp["upg_status"] = (s.GetUpgStatus())
				temp["nr_version"] = (s.GetVersion())
				niatelemetryNiaInventoryDcnmResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryNiaInventoryDcnmResults))
	if err := d.Set("results", niatelemetryNiaInventoryDcnmResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryNiaInventoryDcnmResults[0]["moid"].(string))
	return de
}
