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

func dataSourceIamUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamUserRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_ip_address": {
				Description: "IP address from which the user last logged in to Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"email": {
				Description: "Email of the user. Users are added to Intersight using the email configured in the IdP.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"first_name": {
				Description: "First name of the user. This field is populated from the IdP attributes received after authentication.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_login_time": {
				Description: "Last successful login time for user.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_name": {
				Description: "Last name of the user. This field is populated from the IdP attributes received after authentication.",
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
				Description: "Name as configured in the IdP.",
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
			"user_id_or_email": {
				Description: "UserID or email as configured in the IdP.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"user_type": {
				Description: "Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"user_unique_identifier": {
				Description: "Unique id of the user used by the identity provider to store the user.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamUser().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamUserRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamUser{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("client_ip_address"); ok {
		x := (v.(string))
		o.SetClientIpAddress(x)
	}
	if v, ok := d.GetOk("email"); ok {
		x := (v.(string))
		o.SetEmail(x)
	}
	if v, ok := d.GetOk("first_name"); ok {
		x := (v.(string))
		o.SetFirstName(x)
	}
	if v, ok := d.GetOk("last_login_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastLoginTime(x)
	}
	if v, ok := d.GetOk("last_name"); ok {
		x := (v.(string))
		o.SetLastName(x)
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
	if v, ok := d.GetOk("user_id_or_email"); ok {
		x := (v.(string))
		o.SetUserIdOrEmail(x)
	}
	if v, ok := d.GetOk("user_type"); ok {
		x := (v.(string))
		o.SetUserType(x)
	}
	if v, ok := d.GetOk("user_unique_identifier"); ok {
		x := (v.(string))
		o.SetUserUniqueIdentifier(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamUser object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamUserList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamUser: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamUserList.GetCount()
	var i int32
	var iamUserResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamUserList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamUser: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamUserList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamUser data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["api_keys"] = flattenListIamApiKeyRelationship(s.GetApiKeys(), d)

				temp["app_registrations"] = flattenListIamAppRegistrationRelationship(s.GetAppRegistrations(), d)
				temp["class_id"] = (s.GetClassId())
				temp["client_ip_address"] = (s.GetClientIpAddress())
				temp["email"] = (s.GetEmail())
				temp["first_name"] = (s.GetFirstName())

				temp["idp"] = flattenMapIamIdpRelationship(s.GetIdp(), d)

				temp["idpreference"] = flattenMapIamIdpReferenceRelationship(s.GetIdpreference(), d)

				temp["last_login_time"] = (s.GetLastLoginTime()).String()
				temp["last_name"] = (s.GetLastName())

				temp["local_user_password"] = flattenMapIamLocalUserPasswordRelationship(s.GetLocalUserPassword(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["oauth_tokens"] = flattenListIamOAuthTokenRelationship(s.GetOauthTokens(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["permissions"] = flattenListIamPermissionRelationship(s.GetPermissions(), d)

				temp["sessions"] = flattenListIamSessionRelationship(s.GetSessions(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["user_id_or_email"] = (s.GetUserIdOrEmail())
				temp["user_type"] = (s.GetUserType())
				temp["user_unique_identifier"] = (s.GetUserUniqueIdentifier())
				iamUserResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamUserResults))
	if err := d.Set("results", iamUserResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamUserResults[0]["moid"].(string))
	return de
}
