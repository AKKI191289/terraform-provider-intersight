package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEtherPortChannel() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEtherPortChannelRead,
		Schema: map[string]*schema.Schema{
			"access_vlan": {
				Description: "Access VLANs for this port-channel, on this FI.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"admin_state": {
				Description: "Administratively configured state (enabled/disabled) for this port-channel.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"allowed_vlans": {
				Description: "Allowed VLANs on this port-channel, on this FI.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_mo_id": {
				Description: "The database identifier of the registered device of an object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Description: "Operating mode of this port-channel.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"native_vlan": {
				Description: "Native VLAN for this port-channel, on this FI.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_speed": {
				Description: "Operational speed of this port-channel.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"oper_state": {
				Description: "Operational state of this port-channel.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"oper_state_qual": {
				Description: "Reason for this port-channel's Operational state.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"port_channel_id": {
				Description: "Unique identifier for this port-channel on the FI.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Description: "This port-channel's configured role (uplink, server, etc.).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"switch_id": {
				Description: "Switch Identifier that is local to a cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"access_vlan": {
					Description: "Access VLANs for this port-channel, on this FI.",
					Type:        schema.TypeString,
					Optional:    true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"admin_state": {
						Description: "Administratively configured state (enabled/disabled) for this port-channel.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"allowed_vlans": {
						Description: "Allowed VLANs on this port-channel, on this FI.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"device_mo_id": {
						Description: "The database identifier of the registered device of an object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"dn": {
						Description: "The Distinguished Name unambiguously identifies an object in the system.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"equipment_switch_card": {
						Description: "A reference to a equipmentSwitchCard resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"mode": {
						Description: "Operating mode of this port-channel.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"native_vlan": {
						Description: "Native VLAN for this port-channel, on this FI.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"oper_speed": {
						Description: "Operational speed of this port-channel.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"oper_state": {
						Description: "Operational state of this port-channel.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"oper_state_qual": {
						Description: "Reason for this port-channel's Operational state.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"port_channel_id": {
						Description: "Unique identifier for this port-channel on the FI.",
						Type:        schema.TypeInt,
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
					"rn": {
						Description: "The Relative Name uniquely identifies an object within a given context.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"role": {
						Description: "This port-channel's configured role (uplink, server, etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"switch_id": {
						Description: "Switch Identifier that is local to a cluster.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceEtherPortChannelRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.EtherPortChannel{}
	if v, ok := d.GetOk("access_vlan"); ok {
		x := (v.(string))
		o.SetAccessVlan(x)
	}
	if v, ok := d.GetOk("admin_state"); ok {
		x := (v.(string))
		o.SetAdminState(x)
	}
	if v, ok := d.GetOk("allowed_vlans"); ok {
		x := (v.(string))
		o.SetAllowedVlans(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("mode"); ok {
		x := (v.(string))
		o.SetMode(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("native_vlan"); ok {
		x := (v.(string))
		o.SetNativeVlan(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oper_speed"); ok {
		x := (v.(string))
		o.SetOperSpeed(x)
	}
	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.SetOperState(x)
	}
	if v, ok := d.GetOk("oper_state_qual"); ok {
		x := (v.(string))
		o.SetOperStateQual(x)
	}
	if v, ok := d.GetOk("port_channel_id"); ok {
		x := int64(v.(int))
		o.SetPortChannelId(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("role"); ok {
		x := (v.(string))
		o.SetRole(x)
	}
	if v, ok := d.GetOk("switch_id"); ok {
		x := (v.(string))
		o.SetSwitchId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of EtherPortChannel object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.EtherApi.GetEtherPortChannelList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of EtherPortChannel: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.EtherPortChannelList.GetCount()
	var i int32
	var etherPortChannelResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.EtherApi.GetEtherPortChannelList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching EtherPortChannel: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.EtherPortChannelList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for EtherPortChannel data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["access_vlan"] = (s.GetAccessVlan())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["admin_state"] = (s.GetAdminState())
				temp["allowed_vlans"] = (s.GetAllowedVlans())
				temp["class_id"] = (s.GetClassId())
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["dn"] = (s.GetDn())

				temp["equipment_switch_card"] = flattenMapEquipmentSwitchCardRelationship(s.GetEquipmentSwitchCard(), d)
				temp["mode"] = (s.GetMode())
				temp["moid"] = (s.GetMoid())
				temp["native_vlan"] = (s.GetNativeVlan())
				temp["object_type"] = (s.GetObjectType())
				temp["oper_speed"] = (s.GetOperSpeed())
				temp["oper_state"] = (s.GetOperState())
				temp["oper_state_qual"] = (s.GetOperStateQual())
				temp["port_channel_id"] = (s.GetPortChannelId())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["rn"] = (s.GetRn())
				temp["role"] = (s.GetRole())
				temp["switch_id"] = (s.GetSwitchId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				etherPortChannelResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(etherPortChannelResults))
	if err := d.Set("results", etherPortChannelResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(etherPortChannelResults[0]["moid"].(string))
	return de
}
