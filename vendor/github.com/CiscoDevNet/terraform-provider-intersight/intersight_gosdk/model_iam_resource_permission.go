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
	"reflect"
	"strings"
)

// IamResourcePermission ResourcePermission represents the permissions defined on a resource like organization.
type IamResourcePermission struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                 `json:"ObjectType"`
	PermissionRoles []IamPermissionToRoles `json:"PermissionRoles,omitempty"`
	// Name of the service owning the resource.
	TargetApp            *string                        `json:"TargetApp,omitempty"`
	Holder               *IamSecurityHolderRelationship `json:"Holder,omitempty"`
	Resource             *MoBaseMoRelationship          `json:"Resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamResourcePermission IamResourcePermission

// NewIamResourcePermission instantiates a new IamResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamResourcePermission(classId string, objectType string) *IamResourcePermission {
	this := IamResourcePermission{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamResourcePermissionWithDefaults instantiates a new IamResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamResourcePermissionWithDefaults() *IamResourcePermission {
	this := IamResourcePermission{}
	var classId string = "iam.ResourcePermission"
	this.ClassId = classId
	var objectType string = "iam.ResourcePermission"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamResourcePermission) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamResourcePermission) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamResourcePermission) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamResourcePermission) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamResourcePermission) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamResourcePermission) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPermissionRoles returns the PermissionRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamResourcePermission) GetPermissionRoles() []IamPermissionToRoles {
	if o == nil {
		var ret []IamPermissionToRoles
		return ret
	}
	return o.PermissionRoles
}

// GetPermissionRolesOk returns a tuple with the PermissionRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamResourcePermission) GetPermissionRolesOk() (*[]IamPermissionToRoles, bool) {
	if o == nil || o.PermissionRoles == nil {
		return nil, false
	}
	return &o.PermissionRoles, true
}

// HasPermissionRoles returns a boolean if a field has been set.
func (o *IamResourcePermission) HasPermissionRoles() bool {
	if o != nil && o.PermissionRoles != nil {
		return true
	}

	return false
}

// SetPermissionRoles gets a reference to the given []IamPermissionToRoles and assigns it to the PermissionRoles field.
func (o *IamResourcePermission) SetPermissionRoles(v []IamPermissionToRoles) {
	o.PermissionRoles = v
}

// GetTargetApp returns the TargetApp field value if set, zero value otherwise.
func (o *IamResourcePermission) GetTargetApp() string {
	if o == nil || o.TargetApp == nil {
		var ret string
		return ret
	}
	return *o.TargetApp
}

// GetTargetAppOk returns a tuple with the TargetApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourcePermission) GetTargetAppOk() (*string, bool) {
	if o == nil || o.TargetApp == nil {
		return nil, false
	}
	return o.TargetApp, true
}

// HasTargetApp returns a boolean if a field has been set.
func (o *IamResourcePermission) HasTargetApp() bool {
	if o != nil && o.TargetApp != nil {
		return true
	}

	return false
}

// SetTargetApp gets a reference to the given string and assigns it to the TargetApp field.
func (o *IamResourcePermission) SetTargetApp(v string) {
	o.TargetApp = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *IamResourcePermission) GetHolder() IamSecurityHolderRelationship {
	if o == nil || o.Holder == nil {
		var ret IamSecurityHolderRelationship
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourcePermission) GetHolderOk() (*IamSecurityHolderRelationship, bool) {
	if o == nil || o.Holder == nil {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *IamResourcePermission) HasHolder() bool {
	if o != nil && o.Holder != nil {
		return true
	}

	return false
}

// SetHolder gets a reference to the given IamSecurityHolderRelationship and assigns it to the Holder field.
func (o *IamResourcePermission) SetHolder(v IamSecurityHolderRelationship) {
	o.Holder = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *IamResourcePermission) GetResource() MoBaseMoRelationship {
	if o == nil || o.Resource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourcePermission) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *IamResourcePermission) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMoRelationship and assigns it to the Resource field.
func (o *IamResourcePermission) SetResource(v MoBaseMoRelationship) {
	o.Resource = &v
}

func (o IamResourcePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PermissionRoles != nil {
		toSerialize["PermissionRoles"] = o.PermissionRoles
	}
	if o.TargetApp != nil {
		toSerialize["TargetApp"] = o.TargetApp
	}
	if o.Holder != nil {
		toSerialize["Holder"] = o.Holder
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamResourcePermission) UnmarshalJSON(bytes []byte) (err error) {
	type IamResourcePermissionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType      string                 `json:"ObjectType"`
		PermissionRoles []IamPermissionToRoles `json:"PermissionRoles,omitempty"`
		// Name of the service owning the resource.
		TargetApp *string                        `json:"TargetApp,omitempty"`
		Holder    *IamSecurityHolderRelationship `json:"Holder,omitempty"`
		Resource  *MoBaseMoRelationship          `json:"Resource,omitempty"`
	}

	varIamResourcePermissionWithoutEmbeddedStruct := IamResourcePermissionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamResourcePermissionWithoutEmbeddedStruct)
	if err == nil {
		varIamResourcePermission := _IamResourcePermission{}
		varIamResourcePermission.ClassId = varIamResourcePermissionWithoutEmbeddedStruct.ClassId
		varIamResourcePermission.ObjectType = varIamResourcePermissionWithoutEmbeddedStruct.ObjectType
		varIamResourcePermission.PermissionRoles = varIamResourcePermissionWithoutEmbeddedStruct.PermissionRoles
		varIamResourcePermission.TargetApp = varIamResourcePermissionWithoutEmbeddedStruct.TargetApp
		varIamResourcePermission.Holder = varIamResourcePermissionWithoutEmbeddedStruct.Holder
		varIamResourcePermission.Resource = varIamResourcePermissionWithoutEmbeddedStruct.Resource
		*o = IamResourcePermission(varIamResourcePermission)
	} else {
		return err
	}

	varIamResourcePermission := _IamResourcePermission{}

	err = json.Unmarshal(bytes, &varIamResourcePermission)
	if err == nil {
		o.MoBaseMo = varIamResourcePermission.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PermissionRoles")
		delete(additionalProperties, "TargetApp")
		delete(additionalProperties, "Holder")
		delete(additionalProperties, "Resource")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamResourcePermission struct {
	value *IamResourcePermission
	isSet bool
}

func (v NullableIamResourcePermission) Get() *IamResourcePermission {
	return v.value
}

func (v *NullableIamResourcePermission) Set(val *IamResourcePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableIamResourcePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableIamResourcePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamResourcePermission(val *IamResourcePermission) *NullableIamResourcePermission {
	return &NullableIamResourcePermission{value: val, isSet: true}
}

func (v NullableIamResourcePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamResourcePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
