package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMetaDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMetaDefinitionRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_concrete": {
				Description: "Boolean flag to specify whether the meta class is a concrete class or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"meta_type": {
				Description: "Indicates whether the meta class is a complex type or managed object.\n* `ManagedObject` - The meta.Definition object describes a managed object.\n* `ComplexType` - The meta.Definition object describes a nested complex type within a managed object.",
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
				Description: "The fully-qualified class name of the Managed Object or complex type. For example, \"compute:Blade\" where the Managed Object is \"Blade\" and the package is 'compute'.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"namespace": {
				Description: "The namespace of the meta.",
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
			"parent_class": {
				Description: "The fully-qualified name of the parent metaclass in the class inheritance hierarchy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_supported": {
				Description: "Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"rbac_resource": {
				Description: "Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource=yes. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"rest_path": {
				Description: "Restful URL path for the meta.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nr_version": {
				Description: "The version of the service that defines the meta-data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"access_privileges": {
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
							"method": {
								Description: "The type of CRUD operation (create, read, update, delete) for which an access privilege is required.\n* `Update` - The 'update' operation/state.\n* `Create` - The 'create' operation/state.\n* `Read` - The 'read' operation/state.\n* `Delete` - The 'delete' operation/state.",
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
							"privilege": {
								Description: "The name of the privilege which is required to invoke the specified CRUD method.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
						},
					},
					Computed: true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"ancestor_classes": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"display_name_metas": {
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
								"format": {
									Description: "A specification for constructing the displayname from the MO's properties.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"include_ancestor": {
									Description: "An indication of whether the displayname should be contructed 'recursively' including the displayname of the first ancestor with a similarly named displayname.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "The name of the displayname used as a key in the DisplayName map which is returned as part of an MO for a Rest request.",
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
						Computed: true,
					},
					"is_concrete": {
						Description: "Boolean flag to specify whether the meta class is a concrete class or not.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"meta_type": {
						Description: "Indicates whether the meta class is a complex type or managed object.\n* `ManagedObject` - The meta.Definition object describes a managed object.\n* `ComplexType` - The meta.Definition object describes a nested complex type within a managed object.",
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
						Description: "The fully-qualified class name of the Managed Object or complex type. For example, \"compute:Blade\" where the Managed Object is \"Blade\" and the package is 'compute'.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"namespace": {
						Description: "The namespace of the meta.",
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
					"parent_class": {
						Description: "The fully-qualified name of the parent metaclass in the class inheritance hierarchy.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"permission_supported": {
						Description: "Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"properties": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"api_access": {
									Description: "API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc.\n* `NoAccess` - The property is not accessible from the API.\n* `ReadOnly` - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property.\n* `CreateOnly` - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests.\n* `ReadWrite` - The property has read/write access.\n* `WriteOnly` - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords.\n* `ReadOnCreate` - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"name": {
									Description: "The name of the property.",
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
								"op_security": {
									Description: "The data-at-rest security protection applied to this property when a Managed Object is persisted.\nFor example, Encrypted or Cleartext.\n* `ClearText` - Data at rest is not encrypted using account specific keys.Note that data is always protected using volume encryption. ClearText propertiesare queryable and searchable.\n* `Encrypted` - The value of the property is encrypted with account-specific cryptographic keys.This setting is used for properties that contain sensitive data. Encrypted propertiesare not queryable and searchable.\n* `Pbkdf2` - The value of the property is hashed using the pbkdf2 key derivation functions thattakes a password, a salt, and a cost factor as inputs then generates a password hash.Its purpose is to make each password guessing trial by an attacker who has obtaineda password hash file expensive and therefore the cost of a guessing attack high or prohibitive.\n* `Bcrypt` - The value of the property is hashed using the bcrypt key derivation function.\n* `Sha512crypt` - The value of the property is hashed using the sha512crypt key derivation function.\n* `Argon2id` - The value of the property is hashed using the argon2id key derivation function.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"search_weight": {
									Description: "Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable.",
									Type:        schema.TypeFloat,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"rbac_resource": {
						Description: "Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource=yes. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"relationships": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"api_access": {
									Description: "API access definition for this relationship.\n* `NoAccess` - The property is not accessible from the API.\n* `ReadOnly` - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property.\n* `CreateOnly` - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests.\n* `ReadWrite` - The property has read/write access.\n* `WriteOnly` - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords.\n* `ReadOnCreate` - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"collection": {
									Description: "Specifies whether the relationship is a collection.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"export": {
									Description: "When turned off, the peer MO is not exported when the local MO is exported.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"export_with_peer": {
									Description: "When turned on, the local MO is exported when the peer is exported.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "The name of the relationship.",
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
								"peer_rel_name": {
									Description: "Name of relationship in peer managed object.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"peer_sync": {
									Description: "When turned on, peer MO corresponding to the reference provided in relation is updated with a reference to the current MO.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"type": {
									Description: "Fully qualified type of the peer managed object.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"rest_path": {
						Description: "Restful URL path for the meta.",
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
					"nr_version": {
						Description: "The version of the service that defines the meta-data.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceMetaDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.MetaDefinition{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("is_concrete"); ok {
		x := (v.(bool))
		o.SetIsConcrete(x)
	}
	if v, ok := d.GetOk("meta_type"); ok {
		x := (v.(string))
		o.SetMetaType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("namespace"); ok {
		x := (v.(string))
		o.SetNamespace(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("parent_class"); ok {
		x := (v.(string))
		o.SetParentClass(x)
	}
	if v, ok := d.GetOk("permission_supported"); ok {
		x := (v.(bool))
		o.SetPermissionSupported(x)
	}
	if v, ok := d.GetOk("rbac_resource"); ok {
		x := (v.(bool))
		o.SetRbacResource(x)
	}
	if v, ok := d.GetOk("rest_path"); ok {
		x := (v.(string))
		o.SetRestPath(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of MetaDefinition object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.MetaApi.GetMetaDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of MetaDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.MetaDefinitionList.GetCount()
	var i int32
	var metaDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.MetaApi.GetMetaDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching MetaDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.MetaDefinitionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for MetaDefinition data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["access_privileges"] = flattenListMetaAccessPrivilege(s.GetAccessPrivileges(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["ancestor_classes"] = (s.GetAncestorClasses())
				temp["class_id"] = (s.GetClassId())

				temp["display_name_metas"] = flattenListMetaDisplayNameDefinition(s.GetDisplayNameMetas(), d)
				temp["is_concrete"] = (s.GetIsConcrete())
				temp["meta_type"] = (s.GetMetaType())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["namespace"] = (s.GetNamespace())
				temp["object_type"] = (s.GetObjectType())
				temp["parent_class"] = (s.GetParentClass())
				temp["permission_supported"] = (s.GetPermissionSupported())

				temp["properties"] = flattenListMetaPropDefinition(s.GetProperties(), d)
				temp["rbac_resource"] = (s.GetRbacResource())

				temp["relationships"] = flattenListMetaRelationshipDefinition(s.GetRelationships(), d)
				temp["rest_path"] = (s.GetRestPath())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())
				metaDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(metaDefinitionResults))
	if err := d.Set("results", metaDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(metaDefinitionResults[0]["moid"].(string))
	return de
}
