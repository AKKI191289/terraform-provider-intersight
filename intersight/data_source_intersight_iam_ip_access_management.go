package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIamIpAccessManagement() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamIpAccessManagementRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable": {
				Description: "Flag stores the state of IP address based access management. Access management is enabled when it's true.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"last_recovery_time": {
				Description: "The access to account gets locked out if wrong IP addresses are configured. Account Administrators have privilege to unblock the account. It stores the time when the account was last recovered from lock out.",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamIpAccessManagement().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamIpAccessManagementRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamIpAccessManagement{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("enable"); ok {
		x := (v.(bool))
		o.SetEnable(x)
	}
	if v, ok := d.GetOk("last_recovery_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastRecoveryTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamIpAccessManagement object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamIpAccessManagementList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamIpAccessManagement: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamIpAccessManagementList.GetCount()
	var i int32
	var iamIpAccessManagementResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamIpAccessManagementList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamIpAccessManagement: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamIpAccessManagementList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamIpAccessManagement data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["enable"] = (s.GetEnable())

				temp["holder"] = flattenMapIamSecurityHolderRelationship(s.GetHolder(), d)

				temp["ip_addresses"] = flattenListIamIpAddressRelationship(s.GetIpAddresses(), d)

				temp["last_recovery_time"] = (s.GetLastRecoveryTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				iamIpAccessManagementResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamIpAccessManagementResults))
	if err := d.Set("results", iamIpAccessManagementResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamIpAccessManagementResults[0]["moid"].(string))
	return de
}
