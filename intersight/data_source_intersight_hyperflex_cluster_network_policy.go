package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexClusterNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterNetworkPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"jumbo_frame": {
				Description: "Enable or disable jumbo frames.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uplink_speed": {
				Description: "Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only.\n* `default` - Current default value set on the hardware platform.\n* `1G` - A link speed of 1 gigabit per second.\n* `10G` - A link speed of 10 gigabits per second or above.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexClusterNetworkPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexClusterNetworkPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexClusterNetworkPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("jumbo_frame"); ok {
		x := (v.(bool))
		o.SetJumboFrame(x)
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
	if v, ok := d.GetOk("uplink_speed"); ok {
		x := (v.(string))
		o.SetUplinkSpeed(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexClusterNetworkPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterNetworkPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexClusterNetworkPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexClusterNetworkPolicyList.GetCount()
	var i int32
	var hyperflexClusterNetworkPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterNetworkPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexClusterNetworkPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexClusterNetworkPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexClusterNetworkPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_profiles"] = flattenListHyperflexClusterProfileRelationship(s.GetClusterProfiles(), d)
				temp["description"] = (s.GetDescription())
				temp["jumbo_frame"] = (s.GetJumboFrame())

				temp["kvm_ip_range"] = flattenMapHyperflexIpAddrRange(s.GetKvmIpRange(), d)

				temp["mac_prefix_range"] = flattenMapHyperflexMacAddrPrefixRange(s.GetMacPrefixRange(), d)

				temp["mgmt_vlan"] = flattenMapHyperflexNamedVlan(s.GetMgmtVlan(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["uplink_speed"] = (s.GetUplinkSpeed())

				temp["vm_migration_vlan"] = flattenMapHyperflexNamedVlan(s.GetVmMigrationVlan(), d)

				temp["vm_network_vlans"] = flattenListHyperflexNamedVlan(s.GetVmNetworkVlans(), d)
				hyperflexClusterNetworkPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexClusterNetworkPolicyResults))
	if err := d.Set("results", hyperflexClusterNetworkPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexClusterNetworkPolicyResults[0]["moid"].(string))
	return de
}
