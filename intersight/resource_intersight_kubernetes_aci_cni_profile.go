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

func resourceKubernetesAciCniProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKubernetesAciCniProfileCreate,
		ReadContext:   resourceKubernetesAciCniProfileRead,
		UpdateContext: resourceKubernetesAciCniProfileUpdate,
		DeleteContext: resourceKubernetesAciCniProfileDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"aaep_name": {
				Description: "Name of ACI AAEP (Attachable Access Entity Profile) to be used for all Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
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
			"cluster_aci_allocations": {
				Description: "An array of relationships to kubernetesAciCniTenantClusterAllocation resources.",
				Type:        schema.TypeList,
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
			},
			"cluster_profiles": {
				Description: "An array of relationships to kubernetesClusterProfile resources.",
				Type:        schema.TypeList,
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
			},
			"description": {
				Description: "Description of the profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ext_svc_dyn_subnet_start": {
				Description: "Start of range of IP subnets for external services with dynamic IP allocation for use by Kubernetes clusters using this ACI CNI policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ext_svc_static_subnet_start": {
				Description: "Start of range of IP subnets for external services with static IP allocation for use by Kubernetes clusters using this ACI CNI policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"infra_vlan_id": {
				Description: "Value of ACI infrastructuere VLAN ID for the ACI fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"l3_out_network_name": {
				Description: "Name of ACI L3Out network to be used for all Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"l3_out_policy_name": {
				Description: "Name of ACI L3Out policy to be used for all Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"l3_out_tenant": {
				Description: "Tenant in ACI used by this L3Out and Common VRF.",
				Type:        schema.TypeString,
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
				Description: "Name of the concrete profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nested_vmm_domain": {
				Description: "VMM domain within which Kubernetes clusters using this policy are nested.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"node_svc_subnet_start": {
				Description: "Start of range of ACI Node Service IP subnets to use by Kubernetes clusters using this ACI CNI policy This is used for the service graph which is used for ACI PBR based load balancing.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"node_vlan_range_end": {
				Description: "Ending value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"node_vlan_range_start": {
				Description: "Starting value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_kubernetes_clusters": {
				Description: "Number of k8s clusters currently using this ACI CNI profile.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"opflex_multicast_address_range": {
				Description: "Range of IP Multicast addresses to be used by the Opflex protocol for Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"pod_subnet_start": {
				Description: "Start of range of Kubernetes pod IP subnets to use by Kubernetes clusters using this ACI CNI policy This should be a /8 IP subnet so that multiple /16 subnets can be assigned for pod subnets of Kubernetes clusters using this profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"src_template": {
				Description: "A reference to a policyAbstractProfile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"svc_subnet_start": {
				Description: "Start of range of Kubernetes Service IP subnets to use by Kubernetes clusters using this ACI CNI policy Currently this is fixed internally and read-only.",
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
				Description: "Defines the type of the profile. Accepted value is instance.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "instance",
			},
			"vrf": {
				Description: "VRF (Virtual Routing and Forwarding) domain to be used within ACI fabric by all k8s clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceKubernetesAciCniProfileCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewKubernetesAciCniProfileWithDefaults()
	if v, ok := d.GetOk("aaep_name"); ok {
		x := (v.(string))
		o.SetAaepName(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("kubernetes.AciCniProfile")

	if v, ok := d.GetOk("cluster_aci_allocations"); ok {
		x := make([]models.KubernetesAciCniTenantClusterAllocationRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
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
			x = append(x, models.MoMoRefAsKubernetesAciCniTenantClusterAllocationRelationship(o))
		}
		if len(x) > 0 {
			o.SetClusterAciAllocations(x)
		}
	}

	if v, ok := d.GetOk("cluster_profiles"); ok {
		x := make([]models.KubernetesClusterProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
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
			x = append(x, models.MoMoRefAsKubernetesClusterProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetClusterProfiles(x)
		}
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("ext_svc_dyn_subnet_start"); ok {
		x := (v.(string))
		o.SetExtSvcDynSubnetStart(x)
	}

	if v, ok := d.GetOk("ext_svc_static_subnet_start"); ok {
		x := (v.(string))
		o.SetExtSvcStaticSubnetStart(x)
	}

	if v, ok := d.GetOk("infra_vlan_id"); ok {
		x := int64(v.(int))
		o.SetInfraVlanId(x)
	}

	if v, ok := d.GetOk("l3_out_network_name"); ok {
		x := (v.(string))
		o.SetL3OutNetworkName(x)
	}

	if v, ok := d.GetOk("l3_out_policy_name"); ok {
		x := (v.(string))
		o.SetL3OutPolicyName(x)
	}

	if v, ok := d.GetOk("l3_out_tenant"); ok {
		x := (v.(string))
		o.SetL3OutTenant(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	if v, ok := d.GetOk("nested_vmm_domain"); ok {
		x := (v.(string))
		o.SetNestedVmmDomain(x)
	}

	if v, ok := d.GetOk("node_svc_subnet_start"); ok {
		x := (v.(string))
		o.SetNodeSvcSubnetStart(x)
	}

	if v, ok := d.GetOk("node_vlan_range_end"); ok {
		x := int64(v.(int))
		o.SetNodeVlanRangeEnd(x)
	}

	if v, ok := d.GetOk("node_vlan_range_start"); ok {
		x := int64(v.(int))
		o.SetNodeVlanRangeStart(x)
	}

	if v, ok := d.GetOk("number_of_kubernetes_clusters"); ok {
		x := int64(v.(int))
		o.SetNumberOfKubernetesClusters(x)
	}

	o.SetObjectType("kubernetes.AciCniProfile")

	if v, ok := d.GetOk("opflex_multicast_address_range"); ok {
		x := (v.(string))
		o.SetOpflexMulticastAddressRange(x)
	}

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

	if v, ok := d.GetOk("pod_subnet_start"); ok {
		x := (v.(string))
		o.SetPodSubnetStart(x)
	}

	if v, ok := d.GetOk("registered_device"); ok {
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if v, ok := d.GetOk("src_template"); ok {
		p := make([]models.PolicyAbstractProfileRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsPolicyAbstractProfileRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSrcTemplate(x)
		}
	}

	if v, ok := d.GetOk("svc_subnet_start"); ok {
		x := (v.(string))
		o.SetSvcSubnetStart(x)
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

	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	if v, ok := d.GetOk("vrf"); ok {
		x := (v.(string))
		o.SetVrf(x)
	}

	r := conn.ApiClient.KubernetesApi.CreateKubernetesAciCniProfile(conn.ctx).KubernetesAciCniProfile(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating KubernetesAciCniProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceKubernetesAciCniProfileRead(c, d, meta)...)
}

func resourceKubernetesAciCniProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.KubernetesApi.GetKubernetesAciCniProfileByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "KubernetesAciCniProfile object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching KubernetesAciCniProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("aaep_name", (s.GetAaepName())); err != nil {
		return diag.Errorf("error occurred while setting property AaepName in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("cluster_aci_allocations", flattenListKubernetesAciCniTenantClusterAllocationRelationship(s.GetClusterAciAllocations(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ClusterAciAllocations in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("cluster_profiles", flattenListKubernetesClusterProfileRelationship(s.GetClusterProfiles(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ClusterProfiles in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("ext_svc_dyn_subnet_start", (s.GetExtSvcDynSubnetStart())); err != nil {
		return diag.Errorf("error occurred while setting property ExtSvcDynSubnetStart in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("ext_svc_static_subnet_start", (s.GetExtSvcStaticSubnetStart())); err != nil {
		return diag.Errorf("error occurred while setting property ExtSvcStaticSubnetStart in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("infra_vlan_id", (s.GetInfraVlanId())); err != nil {
		return diag.Errorf("error occurred while setting property InfraVlanId in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("l3_out_network_name", (s.GetL3OutNetworkName())); err != nil {
		return diag.Errorf("error occurred while setting property L3OutNetworkName in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("l3_out_policy_name", (s.GetL3OutPolicyName())); err != nil {
		return diag.Errorf("error occurred while setting property L3OutPolicyName in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("l3_out_tenant", (s.GetL3OutTenant())); err != nil {
		return diag.Errorf("error occurred while setting property L3OutTenant in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("nested_vmm_domain", (s.GetNestedVmmDomain())); err != nil {
		return diag.Errorf("error occurred while setting property NestedVmmDomain in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("node_svc_subnet_start", (s.GetNodeSvcSubnetStart())); err != nil {
		return diag.Errorf("error occurred while setting property NodeSvcSubnetStart in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("node_vlan_range_end", (s.GetNodeVlanRangeEnd())); err != nil {
		return diag.Errorf("error occurred while setting property NodeVlanRangeEnd in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("node_vlan_range_start", (s.GetNodeVlanRangeStart())); err != nil {
		return diag.Errorf("error occurred while setting property NodeVlanRangeStart in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("number_of_kubernetes_clusters", (s.GetNumberOfKubernetesClusters())); err != nil {
		return diag.Errorf("error occurred while setting property NumberOfKubernetesClusters in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("opflex_multicast_address_range", (s.GetOpflexMulticastAddressRange())); err != nil {
		return diag.Errorf("error occurred while setting property OpflexMulticastAddressRange in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Organization in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("pod_subnet_start", (s.GetPodSubnetStart())); err != nil {
		return diag.Errorf("error occurred while setting property PodSubnetStart in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
		return diag.Errorf("error occurred while setting property RegisteredDevice in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("src_template", flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)); err != nil {
		return diag.Errorf("error occurred while setting property SrcTemplate in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("svc_subnet_start", (s.GetSvcSubnetStart())); err != nil {
		return diag.Errorf("error occurred while setting property SvcSubnetStart in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("type", (s.GetType())); err != nil {
		return diag.Errorf("error occurred while setting property Type in KubernetesAciCniProfile object: %s", err.Error())
	}

	if err := d.Set("vrf", (s.GetVrf())); err != nil {
		return diag.Errorf("error occurred while setting property Vrf in KubernetesAciCniProfile object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceKubernetesAciCniProfileUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAciCniProfile{}
	if d.HasChange("aaep_name") {
		v := d.Get("aaep_name")
		x := (v.(string))
		o.SetAaepName(x)
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("kubernetes.AciCniProfile")

	if d.HasChange("cluster_aci_allocations") {
		v := d.Get("cluster_aci_allocations")
		x := make([]models.KubernetesAciCniTenantClusterAllocationRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
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
			x = append(x, models.MoMoRefAsKubernetesAciCniTenantClusterAllocationRelationship(o))
		}
		if len(x) > 0 {
			o.SetClusterAciAllocations(x)
		}
	}

	if d.HasChange("cluster_profiles") {
		v := d.Get("cluster_profiles")
		x := make([]models.KubernetesClusterProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
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
			x = append(x, models.MoMoRefAsKubernetesClusterProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetClusterProfiles(x)
		}
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("ext_svc_dyn_subnet_start") {
		v := d.Get("ext_svc_dyn_subnet_start")
		x := (v.(string))
		o.SetExtSvcDynSubnetStart(x)
	}

	if d.HasChange("ext_svc_static_subnet_start") {
		v := d.Get("ext_svc_static_subnet_start")
		x := (v.(string))
		o.SetExtSvcStaticSubnetStart(x)
	}

	if d.HasChange("infra_vlan_id") {
		v := d.Get("infra_vlan_id")
		x := int64(v.(int))
		o.SetInfraVlanId(x)
	}

	if d.HasChange("l3_out_network_name") {
		v := d.Get("l3_out_network_name")
		x := (v.(string))
		o.SetL3OutNetworkName(x)
	}

	if d.HasChange("l3_out_policy_name") {
		v := d.Get("l3_out_policy_name")
		x := (v.(string))
		o.SetL3OutPolicyName(x)
	}

	if d.HasChange("l3_out_tenant") {
		v := d.Get("l3_out_tenant")
		x := (v.(string))
		o.SetL3OutTenant(x)
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

	if d.HasChange("nested_vmm_domain") {
		v := d.Get("nested_vmm_domain")
		x := (v.(string))
		o.SetNestedVmmDomain(x)
	}

	if d.HasChange("node_svc_subnet_start") {
		v := d.Get("node_svc_subnet_start")
		x := (v.(string))
		o.SetNodeSvcSubnetStart(x)
	}

	if d.HasChange("node_vlan_range_end") {
		v := d.Get("node_vlan_range_end")
		x := int64(v.(int))
		o.SetNodeVlanRangeEnd(x)
	}

	if d.HasChange("node_vlan_range_start") {
		v := d.Get("node_vlan_range_start")
		x := int64(v.(int))
		o.SetNodeVlanRangeStart(x)
	}

	if d.HasChange("number_of_kubernetes_clusters") {
		v := d.Get("number_of_kubernetes_clusters")
		x := int64(v.(int))
		o.SetNumberOfKubernetesClusters(x)
	}

	o.SetObjectType("kubernetes.AciCniProfile")

	if d.HasChange("opflex_multicast_address_range") {
		v := d.Get("opflex_multicast_address_range")
		x := (v.(string))
		o.SetOpflexMulticastAddressRange(x)
	}

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

	if d.HasChange("pod_subnet_start") {
		v := d.Get("pod_subnet_start")
		x := (v.(string))
		o.SetPodSubnetStart(x)
	}

	if d.HasChange("registered_device") {
		v := d.Get("registered_device")
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if d.HasChange("src_template") {
		v := d.Get("src_template")
		p := make([]models.PolicyAbstractProfileRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsPolicyAbstractProfileRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSrcTemplate(x)
		}
	}

	if d.HasChange("svc_subnet_start") {
		v := d.Get("svc_subnet_start")
		x := (v.(string))
		o.SetSvcSubnetStart(x)
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

	if d.HasChange("type") {
		v := d.Get("type")
		x := (v.(string))
		o.SetType(x)
	}

	if d.HasChange("vrf") {
		v := d.Get("vrf")
		x := (v.(string))
		o.SetVrf(x)
	}

	r := conn.ApiClient.KubernetesApi.UpdateKubernetesAciCniProfile(conn.ctx, d.Id()).KubernetesAciCniProfile(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating KubernetesAciCniProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceKubernetesAciCniProfileRead(c, d, meta)...)
}

func resourceKubernetesAciCniProfileDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.KubernetesApi.DeleteKubernetesAciCniProfile(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting KubernetesAciCniProfile object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
