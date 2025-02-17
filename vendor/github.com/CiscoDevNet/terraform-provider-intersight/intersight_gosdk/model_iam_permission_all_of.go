/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IamPermissionAllOf Definition of the list of properties defined in 'iam.Permission', excluding properties defined in parent classes.
type IamPermissionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The informative description about each permission.
	Description *string `json:"Description,omitempty"`
	// The name of the permission which has to be granted to user.
	Name    *string                 `json:"Name,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
	// An array of relationships to iamResourceRoles resources.
	ResourceRoles []IamResourceRolesRelationship `json:"ResourceRoles,omitempty"`
	// An array of relationships to iamRole resources.
	Roles         []IamRoleRelationship         `json:"Roles,omitempty"`
	SessionLimits *IamSessionLimitsRelationship `json:"SessionLimits,omitempty"`
	// An array of relationships to iamUserGroup resources.
	UserGroups []IamUserGroupRelationship `json:"UserGroups,omitempty"`
	// An array of relationships to iamUser resources.
	Users                []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPermissionAllOf IamPermissionAllOf

// NewIamPermissionAllOf instantiates a new IamPermissionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermissionAllOf(classId string, objectType string) *IamPermissionAllOf {
	this := IamPermissionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPermissionAllOfWithDefaults instantiates a new IamPermissionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionAllOfWithDefaults() *IamPermissionAllOf {
	this := IamPermissionAllOf{}
	var classId string = "iam.Permission"
	this.ClassId = classId
	var objectType string = "iam.Permission"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPermissionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPermissionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPermissionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPermissionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPermissionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPermissionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamPermissionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamPermissionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamPermissionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamPermissionAllOf) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamPermissionAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamPermissionAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetEndPointRoles returns the EndPointRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionAllOf) GetEndPointRoles() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRoles
}

// GetEndPointRolesOk returns a tuple with the EndPointRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionAllOf) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool) {
	if o == nil || o.EndPointRoles == nil {
		return nil, false
	}
	return &o.EndPointRoles, true
}

// HasEndPointRoles returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasEndPointRoles() bool {
	if o != nil && o.EndPointRoles != nil {
		return true
	}

	return false
}

// SetEndPointRoles gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRoles field.
func (o *IamPermissionAllOf) SetEndPointRoles(v []IamEndPointRoleRelationship) {
	o.EndPointRoles = v
}

// GetPrivilegeSets returns the PrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionAllOf) GetPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.PrivilegeSets
}

// GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionAllOf) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool) {
	if o == nil || o.PrivilegeSets == nil {
		return nil, false
	}
	return &o.PrivilegeSets, true
}

// HasPrivilegeSets returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasPrivilegeSets() bool {
	if o != nil && o.PrivilegeSets != nil {
		return true
	}

	return false
}

// SetPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the PrivilegeSets field.
func (o *IamPermissionAllOf) SetPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.PrivilegeSets = v
}

// GetResourceRoles returns the ResourceRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionAllOf) GetResourceRoles() []IamResourceRolesRelationship {
	if o == nil {
		var ret []IamResourceRolesRelationship
		return ret
	}
	return o.ResourceRoles
}

// GetResourceRolesOk returns a tuple with the ResourceRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionAllOf) GetResourceRolesOk() (*[]IamResourceRolesRelationship, bool) {
	if o == nil || o.ResourceRoles == nil {
		return nil, false
	}
	return &o.ResourceRoles, true
}

// HasResourceRoles returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasResourceRoles() bool {
	if o != nil && o.ResourceRoles != nil {
		return true
	}

	return false
}

// SetResourceRoles gets a reference to the given []IamResourceRolesRelationship and assigns it to the ResourceRoles field.
func (o *IamPermissionAllOf) SetResourceRoles(v []IamResourceRolesRelationship) {
	o.ResourceRoles = v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionAllOf) GetRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionAllOf) GetRolesOk() (*[]IamRoleRelationship, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
func (o *IamPermissionAllOf) SetRoles(v []IamRoleRelationship) {
	o.Roles = v
}

// GetSessionLimits returns the SessionLimits field value if set, zero value otherwise.
func (o *IamPermissionAllOf) GetSessionLimits() IamSessionLimitsRelationship {
	if o == nil || o.SessionLimits == nil {
		var ret IamSessionLimitsRelationship
		return ret
	}
	return *o.SessionLimits
}

// GetSessionLimitsOk returns a tuple with the SessionLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionAllOf) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool) {
	if o == nil || o.SessionLimits == nil {
		return nil, false
	}
	return o.SessionLimits, true
}

// HasSessionLimits returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasSessionLimits() bool {
	if o != nil && o.SessionLimits != nil {
		return true
	}

	return false
}

// SetSessionLimits gets a reference to the given IamSessionLimitsRelationship and assigns it to the SessionLimits field.
func (o *IamPermissionAllOf) SetSessionLimits(v IamSessionLimitsRelationship) {
	o.SessionLimits = &v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionAllOf) GetUserGroups() []IamUserGroupRelationship {
	if o == nil {
		var ret []IamUserGroupRelationship
		return ret
	}
	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionAllOf) GetUserGroupsOk() (*[]IamUserGroupRelationship, bool) {
	if o == nil || o.UserGroups == nil {
		return nil, false
	}
	return &o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasUserGroups() bool {
	if o != nil && o.UserGroups != nil {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []IamUserGroupRelationship and assigns it to the UserGroups field.
func (o *IamPermissionAllOf) SetUserGroups(v []IamUserGroupRelationship) {
	o.UserGroups = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionAllOf) GetUsers() []IamUserRelationship {
	if o == nil {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionAllOf) GetUsersOk() (*[]IamUserRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamPermissionAllOf) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamPermissionAllOf) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamPermissionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.EndPointRoles != nil {
		toSerialize["EndPointRoles"] = o.EndPointRoles
	}
	if o.PrivilegeSets != nil {
		toSerialize["PrivilegeSets"] = o.PrivilegeSets
	}
	if o.ResourceRoles != nil {
		toSerialize["ResourceRoles"] = o.ResourceRoles
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.SessionLimits != nil {
		toSerialize["SessionLimits"] = o.SessionLimits
	}
	if o.UserGroups != nil {
		toSerialize["UserGroups"] = o.UserGroups
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPermissionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamPermissionAllOf := _IamPermissionAllOf{}

	if err = json.Unmarshal(bytes, &varIamPermissionAllOf); err == nil {
		*o = IamPermissionAllOf(varIamPermissionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "EndPointRoles")
		delete(additionalProperties, "PrivilegeSets")
		delete(additionalProperties, "ResourceRoles")
		delete(additionalProperties, "Roles")
		delete(additionalProperties, "SessionLimits")
		delete(additionalProperties, "UserGroups")
		delete(additionalProperties, "Users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamPermissionAllOf struct {
	value *IamPermissionAllOf
	isSet bool
}

func (v NullableIamPermissionAllOf) Get() *IamPermissionAllOf {
	return v.value
}

func (v *NullableIamPermissionAllOf) Set(val *IamPermissionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPermissionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPermissionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPermissionAllOf(val *IamPermissionAllOf) *NullableIamPermissionAllOf {
	return &NullableIamPermissionAllOf{value: val, isSet: true}
}

func (v NullableIamPermissionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPermissionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
