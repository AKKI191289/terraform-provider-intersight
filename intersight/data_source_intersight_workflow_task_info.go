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

func dataSourceWorkflowTaskInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowTaskInfoRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The task description and this is the description that was added when the task was included into the workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "The time stamp when the task reached a final state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"failure_reason": {
				Description: "Description of the reason why the task failed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inst_id": {
				Description: "The instance ID of the task running in the workflow engine.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"internal": {
				Description: "Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"label": {
				Description: "User friendly short label to describe this task instance in the workflow.",
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
			"name": {
				Description: "Task definition name which specifies the task type.",
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
			"ref_name": {
				Description: "The task reference name to ensure its unique inside a workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"retry_count": {
				Description: "A counter for number of retries.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"rollback_disabled": {
				Description: "The task is disabled/enabled for rollback operation in this workflow if the task has rollback support.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"running_inst_id": {
				Description: "The instance ID of the task that is currently being executed. When retrying a workflow with failed tasks, the task in workflow engine will have a new instance ID, but the task may still be in-progress. In this case, the task instId reflects the instance ID in the workflow engine, while runningInstId reflects the instance ID of the instance that is currently being executed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"start_time": {
				Description: "The time stamp when the task started execution.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "The status of the task and this will specify if the task is running or has reached a final state.",
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
					"description": {
						Description: "The task description and this is the description that was added when the task was included into the workflow.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"end_time": {
						Description: "The time stamp when the task reached a final state.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"failure_reason": {
						Description: "Description of the reason why the task failed.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"input": {
						Description: "The input data that was sent to the task at the start of execution.",
						Type:        schema.TypeMap,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						}, Optional: true,
						Computed: true,
					},
					"inst_id": {
						Description: "The instance ID of the task running in the workflow engine.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"internal": {
						Description: "Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"label": {
						Description: "User friendly short label to describe this task instance in the workflow.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"message": {
						Type:     schema.TypeList,
						Optional: true,
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
								"message": {
									Description: "An i18n message that can be translated in multiple languages to support internationalization.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"severity": {
									Description: "The severity of the Task or Workflow message warning/error/info etc.\n* `Info` - The enum represents the log level to be used to convey info message.\n* `Warning` - The enum represents the log level to be used to convey warning message.\n* `Debug` - The enum represents the log level to be used to convey debug message.\n* `Error` - The enum represents the log level to be used to convey error message.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Task definition name which specifies the task type.",
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
					"output": {
						Description: "The output data that was generated by this task.",
						Type:        schema.TypeMap,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						}, Optional: true,
						Computed: true,
					},
					"ref_name": {
						Description: "The task reference name to ensure its unique inside a workflow.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"retry_count": {
						Description: "A counter for number of retries.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"rollback_disabled": {
						Description: "The task is disabled/enabled for rollback operation in this workflow if the task has rollback support.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"running_inst_id": {
						Description: "The instance ID of the task that is currently being executed. When retrying a workflow with failed tasks, the task in workflow engine will have a new instance ID, but the task may still be in-progress. In this case, the task instId reflects the instance ID in the workflow engine, while runningInstId reflects the instance ID of the instance that is currently being executed.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"start_time": {
						Description: "The time stamp when the task started execution.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"status": {
						Description: "The status of the task and this will specify if the task is running or has reached a final state.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"sub_workflow_info": {
						Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"task_definition": {
						Description: "A reference to a workflowTaskDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"task_inst_id_list": {
						Type:     schema.TypeList,
						Optional: true,
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
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"status": {
									Description: "Status of the retried task.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"task_inst_id": {
									Description: "Retry instance will get a unique instance id.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"workflow_info": {
						Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceWorkflowTaskInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowTaskInfo{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}
	if v, ok := d.GetOk("failure_reason"); ok {
		x := (v.(string))
		o.SetFailureReason(x)
	}
	if v, ok := d.GetOk("inst_id"); ok {
		x := (v.(string))
		o.SetInstId(x)
	}
	if v, ok := d.GetOk("internal"); ok {
		x := (v.(bool))
		o.SetInternal(x)
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
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
	if v, ok := d.GetOk("ref_name"); ok {
		x := (v.(string))
		o.SetRefName(x)
	}
	if v, ok := d.GetOk("retry_count"); ok {
		x := int64(v.(int))
		o.SetRetryCount(x)
	}
	if v, ok := d.GetOk("rollback_disabled"); ok {
		x := (v.(bool))
		o.SetRollbackDisabled(x)
	}
	if v, ok := d.GetOk("running_inst_id"); ok {
		x := (v.(string))
		o.SetRunningInstId(x)
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowTaskInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowTaskInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of WorkflowTaskInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.WorkflowTaskInfoList.GetCount()
	var i int32
	var workflowTaskInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowTaskInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching WorkflowTaskInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.WorkflowTaskInfoList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for WorkflowTaskInfo data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["end_time"] = (s.GetEndTime()).String()
				temp["failure_reason"] = (s.GetFailureReason())
				temp["inst_id"] = (s.GetInstId())
				temp["internal"] = (s.GetInternal())
				temp["label"] = (s.GetLabel())

				temp["message"] = flattenListWorkflowMessage(s.GetMessage(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["ref_name"] = (s.GetRefName())
				temp["retry_count"] = (s.GetRetryCount())
				temp["rollback_disabled"] = (s.GetRollbackDisabled())
				temp["running_inst_id"] = (s.GetRunningInstId())

				temp["start_time"] = (s.GetStartTime()).String()
				temp["status"] = (s.GetStatus())

				temp["sub_workflow_info"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetSubWorkflowInfo(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["task_definition"] = flattenMapWorkflowTaskDefinitionRelationship(s.GetTaskDefinition(), d)

				temp["task_inst_id_list"] = flattenListWorkflowTaskRetryInfo(s.GetTaskInstIdList(), d)

				temp["workflow_info"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)
				workflowTaskInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(workflowTaskInfoResults))
	if err := d.Set("results", workflowTaskInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(workflowTaskInfoResults[0]["moid"].(string))
	return de
}
