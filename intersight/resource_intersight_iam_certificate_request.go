package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIamCertificateRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIamCertificateRequestCreate,
		ReadContext:   resourceIamCertificateRequestRead,
		UpdateContext: resourceIamCertificateRequestUpdate,
		DeleteContext: resourceIamCertificateRequestDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"certificate": {
				Description: "A reference to a iamCertificate resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"email_address": {
				Description: "User input email address, an optional part of the subject of the certificate request.",
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
				Description: "Name of the certificate request.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"private_key_spec": {
				Description: "A reference to a iamPrivateKeySpec resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"request": {
				Description: "Generated certificate signing request (CSR) in PEM format.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"self_signed": {
				Description: "Whether the user wants the generated CSR to be self-signed by the appliance.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"subject": {
				Description: "The x.509 distinguished name of the subject of the certificate request.",
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
						"common_name": {
							Description: "A required component that identifies a person or an object.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"country": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"locality": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"organization": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"organizational_unit": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"state": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"subject_alternate_name": {
				Description: "The x.509 subject alternate name values of the certificate request.",
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
						"dns_name": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"email_address": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"ip_address": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"uri": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
		},
	}
}

func resourceIamCertificateRequestCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewIamCertificateRequestWithDefaults()
	if v, ok := d.GetOk("account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccount(x)
		}
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		p := make([]models.IamCertificateRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamCertificateRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCertificate(x)
		}
	}

	o.SetClassId("iam.CertificateRequest")

	if v, ok := d.GetOk("email_address"); ok {
		x := (v.(string))
		o.SetEmailAddress(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("iam.CertificateRequest")

	if v, ok := d.GetOk("private_key_spec"); ok {
		p := make([]models.IamPrivateKeySpecRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamPrivateKeySpecRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPrivateKeySpec(x)
		}
	}

	if v, ok := d.GetOk("request"); ok {
		x := (v.(string))
		o.SetRequest(x)
	}

	if v, ok := d.GetOkExists("self_signed"); ok {
		x := v.(bool)
		o.SetSelfSigned(x)
	}

	if v, ok := d.GetOk("subject"); ok {
		p := make([]models.PkixDistinguishedName, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewPkixDistinguishedNameWithDefaults()
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
			o.SetClassId("pkix.DistinguishedName")
			if v, ok := l["common_name"]; ok {
				{
					x := (v.(string))
					o.SetCommonName(x)
				}
			}
			if v, ok := l["country"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetCountry(x)
					}
				}
			}
			if v, ok := l["locality"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetLocality(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["organization"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetOrganization(x)
					}
				}
			}
			if v, ok := l["organizational_unit"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetOrganizationalUnit(x)
					}
				}
			}
			if v, ok := l["state"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetState(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSubject(x)
		}
	}

	if v, ok := d.GetOk("subject_alternate_name"); ok {
		p := make([]models.PkixSubjectAlternateName, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewPkixSubjectAlternateNameWithDefaults()
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
			o.SetClassId("pkix.SubjectAlternateName")
			if v, ok := l["dns_name"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetDnsName(x)
					}
				}
			}
			if v, ok := l["email_address"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetEmailAddress(x)
					}
				}
			}
			if v, ok := l["ip_address"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetIpAddress(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["uri"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetUri(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSubjectAlternateName(x)
		}
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

	r := conn.ApiClient.IamApi.CreateIamCertificateRequest(conn.ctx).IamCertificateRequest(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating IamCertificateRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceIamCertificateRequestRead(c, d, meta)...)
}

func resourceIamCertificateRequestRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.IamApi.GetIamCertificateRequestByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "IamCertificateRequest object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching IamCertificateRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Account in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("certificate", flattenMapIamCertificateRelationship(s.GetCertificate(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Certificate in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("email_address", (s.GetEmailAddress())); err != nil {
		return diag.Errorf("error occurred while setting property EmailAddress in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("private_key_spec", flattenMapIamPrivateKeySpecRelationship(s.GetPrivateKeySpec(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PrivateKeySpec in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("request", (s.GetRequest())); err != nil {
		return diag.Errorf("error occurred while setting property Request in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("self_signed", (s.GetSelfSigned())); err != nil {
		return diag.Errorf("error occurred while setting property SelfSigned in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("subject", flattenMapPkixDistinguishedName(s.GetSubject(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Subject in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("subject_alternate_name", flattenMapPkixSubjectAlternateName(s.GetSubjectAlternateName(), d)); err != nil {
		return diag.Errorf("error occurred while setting property SubjectAlternateName in IamCertificateRequest object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in IamCertificateRequest object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceIamCertificateRequestUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamCertificateRequest{}
	if d.HasChange("account") {
		v := d.Get("account")
		p := make([]models.IamAccountRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccount(x)
		}
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

	if d.HasChange("certificate") {
		v := d.Get("certificate")
		p := make([]models.IamCertificateRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamCertificateRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCertificate(x)
		}
	}

	o.SetClassId("iam.CertificateRequest")

	if d.HasChange("email_address") {
		v := d.Get("email_address")
		x := (v.(string))
		o.SetEmailAddress(x)
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

	o.SetObjectType("iam.CertificateRequest")

	if d.HasChange("private_key_spec") {
		v := d.Get("private_key_spec")
		p := make([]models.IamPrivateKeySpecRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamPrivateKeySpecRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPrivateKeySpec(x)
		}
	}

	if d.HasChange("request") {
		v := d.Get("request")
		x := (v.(string))
		o.SetRequest(x)
	}

	if d.HasChange("self_signed") {
		v := d.Get("self_signed")
		x := (v.(bool))
		o.SetSelfSigned(x)
	}

	if d.HasChange("subject") {
		v := d.Get("subject")
		p := make([]models.PkixDistinguishedName, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.PkixDistinguishedName{}
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
			o.SetClassId("pkix.DistinguishedName")
			if v, ok := l["common_name"]; ok {
				{
					x := (v.(string))
					o.SetCommonName(x)
				}
			}
			if v, ok := l["country"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetCountry(x)
					}
				}
			}
			if v, ok := l["locality"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetLocality(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["organization"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetOrganization(x)
					}
				}
			}
			if v, ok := l["organizational_unit"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetOrganizationalUnit(x)
					}
				}
			}
			if v, ok := l["state"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetState(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSubject(x)
		}
	}

	if d.HasChange("subject_alternate_name") {
		v := d.Get("subject_alternate_name")
		p := make([]models.PkixSubjectAlternateName, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.PkixSubjectAlternateName{}
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
			o.SetClassId("pkix.SubjectAlternateName")
			if v, ok := l["dns_name"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetDnsName(x)
					}
				}
			}
			if v, ok := l["email_address"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetEmailAddress(x)
					}
				}
			}
			if v, ok := l["ip_address"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetIpAddress(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["uri"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetUri(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSubjectAlternateName(x)
		}
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

	r := conn.ApiClient.IamApi.UpdateIamCertificateRequest(conn.ctx, d.Id()).IamCertificateRequest(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating IamCertificateRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceIamCertificateRequestRead(c, d, meta)...)
}

func resourceIamCertificateRequestDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.IamApi.DeleteIamCertificateRequest(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting IamCertificateRequest object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
