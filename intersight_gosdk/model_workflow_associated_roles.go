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

// WorkflowAssociatedRoles AssociatedRoles models the inferred tasks to the required roles mapping cached in the workflow definition.
type WorkflowAssociatedRoles struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Stores the identifier of the task definition for which the required roles are cached in the workflow definition. In the case of sub workflow tasks, this property stores the identifier of the workflow that is wrapped in the sub workflow task.
	Moid                 *string                   `json:"Moid,omitempty"`
	Roles                []string                  `json:"Roles,omitempty"`
	TaskNames            []string                  `json:"TaskNames,omitempty"`
	WorkflowRoles        []WorkflowAssociatedRoles `json:"WorkflowRoles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAssociatedRoles WorkflowAssociatedRoles

// NewWorkflowAssociatedRoles instantiates a new WorkflowAssociatedRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAssociatedRoles(classId string, objectType string) *WorkflowAssociatedRoles {
	this := WorkflowAssociatedRoles{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowAssociatedRolesWithDefaults instantiates a new WorkflowAssociatedRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAssociatedRolesWithDefaults() *WorkflowAssociatedRoles {
	this := WorkflowAssociatedRoles{}
	var classId string = "workflow.AssociatedRoles"
	this.ClassId = classId
	var objectType string = "workflow.AssociatedRoles"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowAssociatedRoles) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowAssociatedRoles) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowAssociatedRoles) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowAssociatedRoles) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowAssociatedRoles) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowAssociatedRoles) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *WorkflowAssociatedRoles) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAssociatedRoles) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *WorkflowAssociatedRoles) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *WorkflowAssociatedRoles) SetMoid(v string) {
	o.Moid = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowAssociatedRoles) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowAssociatedRoles) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *WorkflowAssociatedRoles) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *WorkflowAssociatedRoles) SetRoles(v []string) {
	o.Roles = v
}

// GetTaskNames returns the TaskNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowAssociatedRoles) GetTaskNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TaskNames
}

// GetTaskNamesOk returns a tuple with the TaskNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowAssociatedRoles) GetTaskNamesOk() (*[]string, bool) {
	if o == nil || o.TaskNames == nil {
		return nil, false
	}
	return &o.TaskNames, true
}

// HasTaskNames returns a boolean if a field has been set.
func (o *WorkflowAssociatedRoles) HasTaskNames() bool {
	if o != nil && o.TaskNames != nil {
		return true
	}

	return false
}

// SetTaskNames gets a reference to the given []string and assigns it to the TaskNames field.
func (o *WorkflowAssociatedRoles) SetTaskNames(v []string) {
	o.TaskNames = v
}

// GetWorkflowRoles returns the WorkflowRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowAssociatedRoles) GetWorkflowRoles() []WorkflowAssociatedRoles {
	if o == nil {
		var ret []WorkflowAssociatedRoles
		return ret
	}
	return o.WorkflowRoles
}

// GetWorkflowRolesOk returns a tuple with the WorkflowRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowAssociatedRoles) GetWorkflowRolesOk() (*[]WorkflowAssociatedRoles, bool) {
	if o == nil || o.WorkflowRoles == nil {
		return nil, false
	}
	return &o.WorkflowRoles, true
}

// HasWorkflowRoles returns a boolean if a field has been set.
func (o *WorkflowAssociatedRoles) HasWorkflowRoles() bool {
	if o != nil && o.WorkflowRoles != nil {
		return true
	}

	return false
}

// SetWorkflowRoles gets a reference to the given []WorkflowAssociatedRoles and assigns it to the WorkflowRoles field.
func (o *WorkflowAssociatedRoles) SetWorkflowRoles(v []WorkflowAssociatedRoles) {
	o.WorkflowRoles = v
}

func (o WorkflowAssociatedRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.TaskNames != nil {
		toSerialize["TaskNames"] = o.TaskNames
	}
	if o.WorkflowRoles != nil {
		toSerialize["WorkflowRoles"] = o.WorkflowRoles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowAssociatedRoles) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowAssociatedRolesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Stores the identifier of the task definition for which the required roles are cached in the workflow definition. In the case of sub workflow tasks, this property stores the identifier of the workflow that is wrapped in the sub workflow task.
		Moid          *string                   `json:"Moid,omitempty"`
		Roles         []string                  `json:"Roles,omitempty"`
		TaskNames     []string                  `json:"TaskNames,omitempty"`
		WorkflowRoles []WorkflowAssociatedRoles `json:"WorkflowRoles,omitempty"`
	}

	varWorkflowAssociatedRolesWithoutEmbeddedStruct := WorkflowAssociatedRolesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowAssociatedRolesWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowAssociatedRoles := _WorkflowAssociatedRoles{}
		varWorkflowAssociatedRoles.ClassId = varWorkflowAssociatedRolesWithoutEmbeddedStruct.ClassId
		varWorkflowAssociatedRoles.ObjectType = varWorkflowAssociatedRolesWithoutEmbeddedStruct.ObjectType
		varWorkflowAssociatedRoles.Moid = varWorkflowAssociatedRolesWithoutEmbeddedStruct.Moid
		varWorkflowAssociatedRoles.Roles = varWorkflowAssociatedRolesWithoutEmbeddedStruct.Roles
		varWorkflowAssociatedRoles.TaskNames = varWorkflowAssociatedRolesWithoutEmbeddedStruct.TaskNames
		varWorkflowAssociatedRoles.WorkflowRoles = varWorkflowAssociatedRolesWithoutEmbeddedStruct.WorkflowRoles
		*o = WorkflowAssociatedRoles(varWorkflowAssociatedRoles)
	} else {
		return err
	}

	varWorkflowAssociatedRoles := _WorkflowAssociatedRoles{}

	err = json.Unmarshal(bytes, &varWorkflowAssociatedRoles)
	if err == nil {
		o.MoBaseComplexType = varWorkflowAssociatedRoles.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Moid")
		delete(additionalProperties, "Roles")
		delete(additionalProperties, "TaskNames")
		delete(additionalProperties, "WorkflowRoles")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableWorkflowAssociatedRoles struct {
	value *WorkflowAssociatedRoles
	isSet bool
}

func (v NullableWorkflowAssociatedRoles) Get() *WorkflowAssociatedRoles {
	return v.value
}

func (v *NullableWorkflowAssociatedRoles) Set(val *WorkflowAssociatedRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAssociatedRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAssociatedRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAssociatedRoles(val *WorkflowAssociatedRoles) *NullableWorkflowAssociatedRoles {
	return &NullableWorkflowAssociatedRoles{value: val, isSet: true}
}

func (v NullableWorkflowAssociatedRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAssociatedRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
