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

func resourceCertificatemanagementPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCertificatemanagementPolicyCreate,
		ReadContext:   resourceCertificatemanagementPolicyRead,
		UpdateContext: resourceCertificatemanagementPolicyUpdate,
		DeleteContext: resourceCertificatemanagementPolicyDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"certificates": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"certificate": {
							Description: "Certificate that is used for verifying the authorization.",
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
									"issuer": {
										Description: "The X.509 distinguished name of the issuer of this certificate.",
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
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"pem_certificate": {
										Description: "The base64 encoded certificate in PEM format.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"sha256_fingerprint": {
										Description: "The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"signature_algorithm": {
										Description: "Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"subject": {
										Description: "The X.509 distinguished name of the subject of this certificate.",
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
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"enabled": {
							Description: "Enable/Disable the certificate in Certificate Management policy.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"is_privatekey_set": {
							Description: "Indicates whether the value of the 'privatekey' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"privatekey": {
							Description: "Private Key which is used to validate the certificate.",
							Type:        schema.TypeString,
							Optional:    true,
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
				Description: "Description of the policy.",
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
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
				Type:        schema.TypeList,
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

func resourceCertificatemanagementPolicyCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewCertificatemanagementPolicyWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("certificates"); ok {
		x := make([]models.CertificatemanagementCertificateBase, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCertificatemanagementCertificateBaseWithDefaults()
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
			if v, ok := l["certificate"]; ok {
				{
					p := make([]models.X509Certificate, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewX509CertificateWithDefaults()
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
						o.SetClassId("x509.Certificate")
						if v, ok := l["issuer"]; ok {
							{
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
									o.SetIssuer(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["pem_certificate"]; ok {
							{
								x := (v.(string))
								o.SetPemCertificate(x)
							}
						}
						if v, ok := l["sha256_fingerprint"]; ok {
							{
								x := (v.(string))
								o.SetSha256Fingerprint(x)
							}
						}
						if v, ok := l["signature_algorithm"]; ok {
							{
								x := (v.(string))
								o.SetSignatureAlgorithm(x)
							}
						}
						if v, ok := l["subject"]; ok {
							{
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
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetCertificate(x)
					}
				}
			}
			o.SetClassId("certificatemanagement.CertificateBase")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["is_privatekey_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsPrivatekeySet(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["privatekey"]; ok {
				{
					x := (v.(string))
					o.SetPrivatekey(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetCertificates(x)
		}
	}

	o.SetClassId("certificatemanagement.Policy")

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

	o.SetObjectType("certificatemanagement.Policy")

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

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
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
			x = append(x, models.MoMoRefAsPolicyAbstractConfigProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetProfiles(x)
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

	r := conn.ApiClient.CertificatemanagementApi.CreateCertificatemanagementPolicy(conn.ctx).CertificatemanagementPolicy(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating CertificatemanagementPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceCertificatemanagementPolicyRead(c, d, meta)...)
}
func detachCertificatemanagementPolicyProfiles(d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CertificatemanagementPolicy{}
	o.SetClassId("certificatemanagement.Policy")
	o.SetObjectType("certificatemanagement.Policy")
	o.SetProfiles([]models.PolicyAbstractConfigProfileRelationship{})

	r := conn.ApiClient.CertificatemanagementApi.UpdateCertificatemanagementPolicy(conn.ctx, d.Id()).CertificatemanagementPolicy(*o)
	_, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while detaching profile/profiles: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	return de
}

func resourceCertificatemanagementPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.CertificatemanagementApi.GetCertificatemanagementPolicyByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "CertificatemanagementPolicy object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching CertificatemanagementPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("certificates", flattenListCertificatemanagementCertificateBase(s.GetCertificates(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Certificates in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Organization in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Profiles in CertificatemanagementPolicy object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in CertificatemanagementPolicy object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceCertificatemanagementPolicyUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CertificatemanagementPolicy{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("certificates") {
		v := d.Get("certificates")
		x := make([]models.CertificatemanagementCertificateBase, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CertificatemanagementCertificateBase{}
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
			if v, ok := l["certificate"]; ok {
				{
					p := make([]models.X509Certificate, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewX509CertificateWithDefaults()
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
						o.SetClassId("x509.Certificate")
						if v, ok := l["issuer"]; ok {
							{
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
									o.SetIssuer(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["pem_certificate"]; ok {
							{
								x := (v.(string))
								o.SetPemCertificate(x)
							}
						}
						if v, ok := l["sha256_fingerprint"]; ok {
							{
								x := (v.(string))
								o.SetSha256Fingerprint(x)
							}
						}
						if v, ok := l["signature_algorithm"]; ok {
							{
								x := (v.(string))
								o.SetSignatureAlgorithm(x)
							}
						}
						if v, ok := l["subject"]; ok {
							{
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
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetCertificate(x)
					}
				}
			}
			o.SetClassId("certificatemanagement.CertificateBase")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["is_privatekey_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsPrivatekeySet(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["privatekey"]; ok {
				{
					x := (v.(string))
					o.SetPrivatekey(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetCertificates(x)
		}
	}

	o.SetClassId("certificatemanagement.Policy")

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
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

	o.SetObjectType("certificatemanagement.Policy")

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

	if d.HasChange("profiles") {
		v := d.Get("profiles")
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
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
			x = append(x, models.MoMoRefAsPolicyAbstractConfigProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetProfiles(x)
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

	r := conn.ApiClient.CertificatemanagementApi.UpdateCertificatemanagementPolicy(conn.ctx, d.Id()).CertificatemanagementPolicy(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating CertificatemanagementPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceCertificatemanagementPolicyRead(c, d, meta)...)
}

func resourceCertificatemanagementPolicyDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	if p, ok := d.GetOk("profiles"); ok {
		if len(p.([]interface{})) > 0 {
			e := detachCertificatemanagementPolicyProfiles(d, meta)
			if e.HasError() {
				return e
			}
		}
	}
	p := conn.ApiClient.CertificatemanagementApi.DeleteCertificatemanagementPolicy(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting CertificatemanagementPolicy object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
