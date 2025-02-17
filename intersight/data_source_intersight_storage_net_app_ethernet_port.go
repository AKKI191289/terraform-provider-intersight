package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageNetAppEthernetPort() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageNetAppEthernetPortRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "Status of Port to determine if its enabled or not.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mac_address": {
				Description: "Macaddress  of the port available in storage array.",
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
			"mtu": {
				Description: "Maximum transmission unit of the physical port available in storage array.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the port available in storage array.",
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
			"speed": {
				Description: "Operational speed of port measured.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Description: "State of the port available in storage array.\n* `down` - An inactive port is listed as Down.\n* `up` - An active port is listed as Up.\n* `present` - An active port is listed as present.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Type of the port available in storage array.\n* `LAG` - Storage port of type lag.\n* `physical` - LIFs can be configured directly on physical ports.\n* `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "UUID of physical port.",
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
					"array_controller": {
						Description: "A reference to a storageNetAppNode resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"enabled": {
						Description: "Status of Port to determine if its enabled or not.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"mac_address": {
						Description: "Macaddress  of the port available in storage array.",
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
					"mtu": {
						Description: "Maximum transmission unit of the physical port available in storage array.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Name of the port available in storage array.",
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
					"speed": {
						Description: "Operational speed of port measured.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"state": {
						Description: "State of the port available in storage array.\n* `down` - An inactive port is listed as Down.\n* `up` - An active port is listed as Up.\n* `present` - An active port is listed as present.",
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
					"type": {
						Description: "Type of the port available in storage array.\n* `LAG` - Storage port of type lag.\n* `physical` - LIFs can be configured directly on physical ports.\n* `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"uuid": {
						Description: "UUID of physical port.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageNetAppEthernetPortRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageNetAppEthernetPort{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(string))
		o.SetEnabled(x)
	}
	if v, ok := d.GetOk("mac_address"); ok {
		x := (v.(string))
		o.SetMacAddress(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("mtu"); ok {
		x := (v.(string))
		o.SetMtu(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("speed"); ok {
		x := int64(v.(int))
		o.SetSpeed(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageNetAppEthernetPort object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppEthernetPortList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageNetAppEthernetPort: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageNetAppEthernetPortList.GetCount()
	var i int32
	var storageNetAppEthernetPortResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppEthernetPortList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageNetAppEthernetPort: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageNetAppEthernetPortList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageNetAppEthernetPort data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["array_controller"] = flattenMapStorageNetAppNodeRelationship(s.GetArrayController(), d)
				temp["class_id"] = (s.GetClassId())
				temp["enabled"] = (s.GetEnabled())
				temp["mac_address"] = (s.GetMacAddress())
				temp["moid"] = (s.GetMoid())
				temp["mtu"] = (s.GetMtu())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["speed"] = (s.GetSpeed())
				temp["state"] = (s.GetState())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["uuid"] = (s.GetUuid())
				storageNetAppEthernetPortResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageNetAppEthernetPortResults))
	if err := d.Set("results", storageNetAppEthernetPortResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageNetAppEthernetPortResults[0]["moid"].(string))
	return de
}
