package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplianceDataExportPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceDataExportPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable": {
				Description: "Status of the data collection mode. If the value is 'true', then data collection is enabled.",
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
				Description: "Name of the Data Export Policy.",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceApplianceDataExportPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceApplianceDataExportPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceDataExportPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("enable"); ok {
		x := (v.(bool))
		o.SetEnable(x)
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
		return diag.Errorf("json marshal of ApplianceDataExportPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceDataExportPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ApplianceDataExportPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ApplianceDataExportPolicyList.GetCount()
	var i int32
	var applianceDataExportPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceDataExportPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceDataExportPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ApplianceDataExportPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ApplianceDataExportPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["enable"] = (s.GetEnable())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["parent_config"] = flattenMapApplianceDataExportPolicyRelationship(s.GetParentConfig(), d)

				temp["sub_configs"] = flattenListApplianceDataExportPolicyRelationship(s.GetSubConfigs(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				applianceDataExportPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceDataExportPolicyResults))
	if err := d.Set("results", applianceDataExportPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceDataExportPolicyResults[0]["moid"].(string))
	return de
}
