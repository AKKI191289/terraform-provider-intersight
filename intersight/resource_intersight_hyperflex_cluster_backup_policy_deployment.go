package intersight

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHyperflexClusterBackupPolicyDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHyperflexClusterBackupPolicyDeploymentCreate,
		ReadContext:   resourceHyperflexClusterBackupPolicyDeploymentRead,
		UpdateContext: resourceHyperflexClusterBackupPolicyDeploymentUpdate,
		DeleteContext: resourceHyperflexClusterBackupPolicyDeploymentDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"backup_data_store_name": {
				Description: "Backup data store name used during the auto creation of the datastore. All VMs created in this data store will be automatically backed up.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_data_store_size": {
				Description: "Replication data store size in backupDataStoreSizeUnit.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"backup_data_store_size_unit": {
				Description: "Replication data store size.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_target": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description from corresponding ClusterBackupPolicy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"discovered": {
				Description: "True if record created by discovery on HyperFlex cluster.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name from corresponding ClusterBackupPolicy.",
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
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"policy_moid": {
				Description: "Deployed cluster policy moid.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"profile_moid": {
				Description: "Deployed cluster profile moid.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"replication_pair_name_prefix": {
				Description: "Replication cluster pairing name prefix.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"replication_schedule": {
				Description: "Backup policy replication schedule.",
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
						"backup_interval": {
							Description: "Time interval between two copies in minutes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     240,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"snapshot_retention_count": {
				Description: "Number of snapshots that will be retained as part of the Multi Point in Time support.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"source_cluster": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"source_detached": {
				Description: "True if policy was detached from source Hyperflex Cluster.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"source_request_id": {
				Description: "Unique source cluster request ID allowing retry of the same logical request following a transient communication failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"source_uuid": {
				Description: "Uuid of the source Hyperflex Cluster.",
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
			"target_detached": {
				Description: "True if policy was detached from target Hyperflex Cluster.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"target_request_id": {
				Description: "Unique target cluster request ID allowing retry of the same logical request following a transient communication failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"target_uuid": {
				Description: "Uuid of the target Hyperflex Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceHyperflexClusterBackupPolicyDeploymentCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewHyperflexClusterBackupPolicyDeploymentWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("backup_data_store_name"); ok {
		x := (v.(string))
		o.SetBackupDataStoreName(x)
	}

	if v, ok := d.GetOk("backup_data_store_size"); ok {
		x := int64(v.(int))
		o.SetBackupDataStoreSize(x)
	}

	if v, ok := d.GetOk("backup_data_store_size_unit"); ok {
		x := (v.(string))
		o.SetBackupDataStoreSizeUnit(x)
	}

	if v, ok := d.GetOk("backup_target"); ok {
		p := make([]models.HyperflexClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetBackupTarget(x)
		}
	}

	o.SetClassId("hyperflex.ClusterBackupPolicyDeployment")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOkExists("discovered"); ok {
		x := v.(bool)
		o.SetDiscovered(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("hyperflex.ClusterBackupPolicyDeployment")

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if v, ok := d.GetOk("policy_moid"); ok {
		x := (v.(string))
		o.SetPolicyMoid(x)
	}

	if v, ok := d.GetOk("profile_moid"); ok {
		x := (v.(string))
		o.SetProfileMoid(x)
	}

	if v, ok := d.GetOk("replication_pair_name_prefix"); ok {
		x := (v.(string))
		o.SetReplicationPairNamePrefix(x)
	}

	if v, ok := d.GetOk("replication_schedule"); ok {
		p := make([]models.HyperflexReplicationSchedule, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewHyperflexReplicationScheduleWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["backup_interval"]; ok {
				{
					x := int64(v.(int))
					o.SetBackupInterval(x)
				}
			}
			o.SetClassId("hyperflex.ReplicationSchedule")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetReplicationSchedule(x)
		}
	}

	if v, ok := d.GetOk("snapshot_retention_count"); ok {
		x := int64(v.(int))
		o.SetSnapshotRetentionCount(x)
	}

	if v, ok := d.GetOk("source_cluster"); ok {
		p := make([]models.HyperflexClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSourceCluster(x)
		}
	}

	if v, ok := d.GetOkExists("source_detached"); ok {
		x := v.(bool)
		o.SetSourceDetached(x)
	}

	if v, ok := d.GetOk("source_request_id"); ok {
		x := (v.(string))
		o.SetSourceRequestId(x)
	}

	if v, ok := d.GetOk("source_uuid"); ok {
		x := (v.(string))
		o.SetSourceUuid(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOkExists("target_detached"); ok {
		x := v.(bool)
		o.SetTargetDetached(x)
	}

	if v, ok := d.GetOk("target_request_id"); ok {
		x := (v.(string))
		o.SetTargetRequestId(x)
	}

	if v, ok := d.GetOk("target_uuid"); ok {
		x := (v.(string))
		o.SetTargetUuid(x)
	}

	r := conn.ApiClient.HyperflexApi.CreateHyperflexClusterBackupPolicyDeployment(conn.ctx).HyperflexClusterBackupPolicyDeployment(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating HyperflexClusterBackupPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceHyperflexClusterBackupPolicyDeploymentRead(c, d, meta)...)
}

func resourceHyperflexClusterBackupPolicyDeploymentRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.HyperflexApi.GetHyperflexClusterBackupPolicyDeploymentByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "HyperflexClusterBackupPolicyDeployment object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching HyperflexClusterBackupPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("backup_data_store_name", (s.GetBackupDataStoreName())); err != nil {
		return diag.Errorf("error occurred while setting property BackupDataStoreName in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("backup_data_store_size", (s.GetBackupDataStoreSize())); err != nil {
		return diag.Errorf("error occurred while setting property BackupDataStoreSize in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("backup_data_store_size_unit", (s.GetBackupDataStoreSizeUnit())); err != nil {
		return diag.Errorf("error occurred while setting property BackupDataStoreSizeUnit in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("backup_target", flattenMapHyperflexClusterRelationship(s.GetBackupTarget(), d)); err != nil {
		return diag.Errorf("error occurred while setting property BackupTarget in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("discovered", (s.GetDiscovered())); err != nil {
		return diag.Errorf("error occurred while setting property Discovered in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Organization in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("policy_moid", (s.GetPolicyMoid())); err != nil {
		return diag.Errorf("error occurred while setting property PolicyMoid in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("profile_moid", (s.GetProfileMoid())); err != nil {
		return diag.Errorf("error occurred while setting property ProfileMoid in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("replication_pair_name_prefix", (s.GetReplicationPairNamePrefix())); err != nil {
		return diag.Errorf("error occurred while setting property ReplicationPairNamePrefix in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("replication_schedule", flattenMapHyperflexReplicationSchedule(s.GetReplicationSchedule(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ReplicationSchedule in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("snapshot_retention_count", (s.GetSnapshotRetentionCount())); err != nil {
		return diag.Errorf("error occurred while setting property SnapshotRetentionCount in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("source_cluster", flattenMapHyperflexClusterRelationship(s.GetSourceCluster(), d)); err != nil {
		return diag.Errorf("error occurred while setting property SourceCluster in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("source_detached", (s.GetSourceDetached())); err != nil {
		return diag.Errorf("error occurred while setting property SourceDetached in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("source_request_id", (s.GetSourceRequestId())); err != nil {
		return diag.Errorf("error occurred while setting property SourceRequestId in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("source_uuid", (s.GetSourceUuid())); err != nil {
		return diag.Errorf("error occurred while setting property SourceUuid in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("target_detached", (s.GetTargetDetached())); err != nil {
		return diag.Errorf("error occurred while setting property TargetDetached in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("target_request_id", (s.GetTargetRequestId())); err != nil {
		return diag.Errorf("error occurred while setting property TargetRequestId in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	if err := d.Set("target_uuid", (s.GetTargetUuid())); err != nil {
		return diag.Errorf("error occurred while setting property TargetUuid in HyperflexClusterBackupPolicyDeployment object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceHyperflexClusterBackupPolicyDeploymentUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexClusterBackupPolicyDeployment{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("backup_data_store_name") {
		v := d.Get("backup_data_store_name")
		x := (v.(string))
		o.SetBackupDataStoreName(x)
	}

	if d.HasChange("backup_data_store_size") {
		v := d.Get("backup_data_store_size")
		x := int64(v.(int))
		o.SetBackupDataStoreSize(x)
	}

	if d.HasChange("backup_data_store_size_unit") {
		v := d.Get("backup_data_store_size_unit")
		x := (v.(string))
		o.SetBackupDataStoreSizeUnit(x)
	}

	if d.HasChange("backup_target") {
		v := d.Get("backup_target")
		p := make([]models.HyperflexClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetBackupTarget(x)
		}
	}

	o.SetClassId("hyperflex.ClusterBackupPolicyDeployment")

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("discovered") {
		v := d.Get("discovered")
		x := (v.(bool))
		o.SetDiscovered(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("hyperflex.ClusterBackupPolicyDeployment")

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if d.HasChange("policy_moid") {
		v := d.Get("policy_moid")
		x := (v.(string))
		o.SetPolicyMoid(x)
	}

	if d.HasChange("profile_moid") {
		v := d.Get("profile_moid")
		x := (v.(string))
		o.SetProfileMoid(x)
	}

	if d.HasChange("replication_pair_name_prefix") {
		v := d.Get("replication_pair_name_prefix")
		x := (v.(string))
		o.SetReplicationPairNamePrefix(x)
	}

	if d.HasChange("replication_schedule") {
		v := d.Get("replication_schedule")
		p := make([]models.HyperflexReplicationSchedule, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HyperflexReplicationSchedule{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["backup_interval"]; ok {
				{
					x := int64(v.(int))
					o.SetBackupInterval(x)
				}
			}
			o.SetClassId("hyperflex.ReplicationSchedule")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetReplicationSchedule(x)
		}
	}

	if d.HasChange("snapshot_retention_count") {
		v := d.Get("snapshot_retention_count")
		x := int64(v.(int))
		o.SetSnapshotRetentionCount(x)
	}

	if d.HasChange("source_cluster") {
		v := d.Get("source_cluster")
		p := make([]models.HyperflexClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSourceCluster(x)
		}
	}

	if d.HasChange("source_detached") {
		v := d.Get("source_detached")
		x := (v.(bool))
		o.SetSourceDetached(x)
	}

	if d.HasChange("source_request_id") {
		v := d.Get("source_request_id")
		x := (v.(string))
		o.SetSourceRequestId(x)
	}

	if d.HasChange("source_uuid") {
		v := d.Get("source_uuid")
		x := (v.(string))
		o.SetSourceUuid(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if d.HasChange("target_detached") {
		v := d.Get("target_detached")
		x := (v.(bool))
		o.SetTargetDetached(x)
	}

	if d.HasChange("target_request_id") {
		v := d.Get("target_request_id")
		x := (v.(string))
		o.SetTargetRequestId(x)
	}

	if d.HasChange("target_uuid") {
		v := d.Get("target_uuid")
		x := (v.(string))
		o.SetTargetUuid(x)
	}

	r := conn.ApiClient.HyperflexApi.UpdateHyperflexClusterBackupPolicyDeployment(conn.ctx, d.Id()).HyperflexClusterBackupPolicyDeployment(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating HyperflexClusterBackupPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceHyperflexClusterBackupPolicyDeploymentRead(c, d, meta)...)
}

func resourceHyperflexClusterBackupPolicyDeploymentDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.HyperflexApi.DeleteHyperflexClusterBackupPolicyDeployment(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting HyperflexClusterBackupPolicyDeployment object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
