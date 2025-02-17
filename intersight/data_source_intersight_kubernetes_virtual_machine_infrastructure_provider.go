package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesVirtualMachineInfrastructureProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesVirtualMachineInfrastructureProviderRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description for the infrastructure provider.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of an infrastructure provider.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesVirtualMachineInfrastructureProvider().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesVirtualMachineInfrastructureProviderRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesVirtualMachineInfrastructureProvider{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesVirtualMachineInfrastructureProvider object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesVirtualMachineInfrastructureProviderList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesVirtualMachineInfrastructureProvider: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesVirtualMachineInfrastructureProviderList.GetCount()
	var i int32
	var kubernetesVirtualMachineInfrastructureProviderResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesVirtualMachineInfrastructureProviderList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesVirtualMachineInfrastructureProvider: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesVirtualMachineInfrastructureProviderList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesVirtualMachineInfrastructureProvider data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["infra_config"] = flattenMapKubernetesBaseVirtualMachineInfraConfig(s.GetInfraConfig(), d)

				temp["infra_config_policy"] = flattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(s.GetInfraConfigPolicy(), d)

				temp["instance_type"] = flattenMapKubernetesVirtualMachineInstanceTypeRelationship(s.GetInstanceType(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["node_group"] = flattenMapKubernetesNodeGroupProfileRelationship(s.GetNodeGroup(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["target"] = flattenMapAssetDeviceRegistrationRelationship(s.GetTarget(), d)
				kubernetesVirtualMachineInfrastructureProviderResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesVirtualMachineInfrastructureProviderResults))
	if err := d.Set("results", kubernetesVirtualMachineInfrastructureProviderResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesVirtualMachineInfrastructureProviderResults[0]["moid"].(string))
	return de
}
