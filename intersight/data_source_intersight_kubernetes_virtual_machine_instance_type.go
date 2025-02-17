package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesVirtualMachineInstanceType() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesVirtualMachineInstanceTypeRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu": {
				Description: "Number of CPUs allocated to virtual machine.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disk_size": {
				Description: "Ephemeral disk capacity to be provided with units example - 10Gi.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"memory": {
				Description: "Virtual machine memory defined in mebibytes (MiB).",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesVirtualMachineInstanceType().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesVirtualMachineInstanceTypeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesVirtualMachineInstanceType{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cpu"); ok {
		x := int64(v.(int))
		o.SetCpu(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("disk_size"); ok {
		x := int64(v.(int))
		o.SetDiskSize(x)
	}
	if v, ok := d.GetOk("memory"); ok {
		x := int64(v.(int))
		o.SetMemory(x)
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
		return diag.Errorf("json marshal of KubernetesVirtualMachineInstanceType object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesVirtualMachineInstanceTypeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesVirtualMachineInstanceType: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesVirtualMachineInstanceTypeList.GetCount()
	var i int32
	var kubernetesVirtualMachineInstanceTypeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesVirtualMachineInstanceTypeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesVirtualMachineInstanceType: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesVirtualMachineInstanceTypeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesVirtualMachineInstanceType data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["cpu"] = (s.GetCpu())
				temp["description"] = (s.GetDescription())
				temp["disk_size"] = (s.GetDiskSize())
				temp["memory"] = (s.GetMemory())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListKubernetesVirtualMachineInfrastructureProviderRelationship(s.GetProfiles(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				kubernetesVirtualMachineInstanceTypeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesVirtualMachineInstanceTypeResults))
	if err := d.Set("results", kubernetesVirtualMachineInstanceTypeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesVirtualMachineInstanceTypeResults[0]["moid"].(string))
	return de
}
