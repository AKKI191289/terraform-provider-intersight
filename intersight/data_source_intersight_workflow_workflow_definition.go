package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWorkflowWorkflowDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowWorkflowDefinitionRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_version": {
				Description: "When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"description": {
				Description: "The description for this workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"label": {
				Description: "A user friendly short name to identify the workflow. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"license_entitlement": {
				Description: "License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks.\n* `Base` - Base as a License type. It is default license type.\n* `Essential` - Essential as a License type.\n* `Standard` - Standard as a License type.\n* `Advantage` - Advantage as a License type.\n* `Premier` - Premier as a License type.\n* `IWO-Essential` - IWO-Essential as a License type.\n* `IWO-Advantage` - IWO-Advantage as a License type.\n* `IWO-Premier` - IWO-Premier as a License type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"max_task_count": {
				Description: "The maximum number of tasks that can be executed on this workflow.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"max_worker_task_count": {
				Description: "The maximum number of external (worker) tasks that can be executed on this workflow.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nr_version": {
				Description: "The version of the workflow to support multiple versions.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceWorkflowWorkflowDefinition().Schema},
				Computed: true,
			}},
	}
}

func dataSourceWorkflowWorkflowDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowWorkflowDefinition{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_version"); ok {
		x := (v.(bool))
		o.SetDefaultVersion(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}
	if v, ok := d.GetOk("license_entitlement"); ok {
		x := (v.(string))
		o.SetLicenseEntitlement(x)
	}
	if v, ok := d.GetOk("max_task_count"); ok {
		x := int64(v.(int))
		o.SetMaxTaskCount(x)
	}
	if v, ok := d.GetOk("max_worker_task_count"); ok {
		x := int64(v.(int))
		o.SetMaxWorkerTaskCount(x)
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
	if v, ok := d.GetOk("nr_version"); ok {
		x := int64(v.(int))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowWorkflowDefinition object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of WorkflowWorkflowDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.WorkflowWorkflowDefinitionList.GetCount()
	var i int32
	var workflowWorkflowDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching WorkflowWorkflowDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.WorkflowWorkflowDefinitionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for WorkflowWorkflowDefinition data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["catalog"] = flattenMapWorkflowCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())
				temp["default_version"] = (s.GetDefaultVersion())
				temp["description"] = (s.GetDescription())

				temp["input_definition"] = flattenListWorkflowBaseDataType(s.GetInputDefinition(), d)

				temp["input_parameter_set"] = flattenListWorkflowParameterSet(s.GetInputParameterSet(), d)
				temp["label"] = (s.GetLabel())
				temp["license_entitlement"] = (s.GetLicenseEntitlement())
				temp["max_task_count"] = (s.GetMaxTaskCount())
				temp["max_worker_task_count"] = (s.GetMaxWorkerTaskCount())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["output_definition"] = flattenListWorkflowBaseDataType(s.GetOutputDefinition(), d)

				temp["properties"] = flattenMapWorkflowWorkflowProperties(s.GetProperties(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["tasks"] = flattenListWorkflowWorkflowTask(s.GetTasks(), d)

				temp["ui_input_filters"] = flattenListWorkflowUiInputFilter(s.GetUiInputFilters(), d)

				temp["validation_information"] = flattenMapWorkflowValidationInformation(s.GetValidationInformation(), d)
				temp["nr_version"] = (s.GetVersion())

				temp["workflow_metadata"] = flattenMapWorkflowWorkflowMetadataRelationship(s.GetWorkflowMetadata(), d)
				workflowWorkflowDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(workflowWorkflowDefinitionResults))
	if err := d.Set("results", workflowWorkflowDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(workflowWorkflowDefinitionResults[0]["moid"].(string))
	return de
}
