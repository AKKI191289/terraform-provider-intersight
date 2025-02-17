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

// WorkflowArrayDataTypeAllOf Definition of the list of properties defined in 'workflow.ArrayDataType', excluding properties defined in parent classes.
type WorkflowArrayDataTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                    `json:"ObjectType"`
	ArrayItemType NullableWorkflowArrayItem `json:"ArrayItemType,omitempty"`
	// Specify the maximum value of the array.
	Max *int64 `json:"Max,omitempty"`
	// Specify the minimum value of the array.
	Min                  *int64 `json:"Min,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowArrayDataTypeAllOf WorkflowArrayDataTypeAllOf

// NewWorkflowArrayDataTypeAllOf instantiates a new WorkflowArrayDataTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowArrayDataTypeAllOf(classId string, objectType string) *WorkflowArrayDataTypeAllOf {
	this := WorkflowArrayDataTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowArrayDataTypeAllOfWithDefaults instantiates a new WorkflowArrayDataTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowArrayDataTypeAllOfWithDefaults() *WorkflowArrayDataTypeAllOf {
	this := WorkflowArrayDataTypeAllOf{}
	var classId string = "workflow.ArrayDataType"
	this.ClassId = classId
	var objectType string = "workflow.ArrayDataType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowArrayDataTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowArrayDataTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowArrayDataTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowArrayDataTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArrayItemType returns the ArrayItemType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowArrayDataTypeAllOf) GetArrayItemType() WorkflowArrayItem {
	if o == nil || o.ArrayItemType.Get() == nil {
		var ret WorkflowArrayItem
		return ret
	}
	return *o.ArrayItemType.Get()
}

// GetArrayItemTypeOk returns a tuple with the ArrayItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowArrayDataTypeAllOf) GetArrayItemTypeOk() (*WorkflowArrayItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrayItemType.Get(), o.ArrayItemType.IsSet()
}

// HasArrayItemType returns a boolean if a field has been set.
func (o *WorkflowArrayDataTypeAllOf) HasArrayItemType() bool {
	if o != nil && o.ArrayItemType.IsSet() {
		return true
	}

	return false
}

// SetArrayItemType gets a reference to the given NullableWorkflowArrayItem and assigns it to the ArrayItemType field.
func (o *WorkflowArrayDataTypeAllOf) SetArrayItemType(v WorkflowArrayItem) {
	o.ArrayItemType.Set(&v)
}

// SetArrayItemTypeNil sets the value for ArrayItemType to be an explicit nil
func (o *WorkflowArrayDataTypeAllOf) SetArrayItemTypeNil() {
	o.ArrayItemType.Set(nil)
}

// UnsetArrayItemType ensures that no value is present for ArrayItemType, not even an explicit nil
func (o *WorkflowArrayDataTypeAllOf) UnsetArrayItemType() {
	o.ArrayItemType.Unset()
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *WorkflowArrayDataTypeAllOf) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataTypeAllOf) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *WorkflowArrayDataTypeAllOf) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *WorkflowArrayDataTypeAllOf) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *WorkflowArrayDataTypeAllOf) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataTypeAllOf) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *WorkflowArrayDataTypeAllOf) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *WorkflowArrayDataTypeAllOf) SetMin(v int64) {
	o.Min = &v
}

func (o WorkflowArrayDataTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArrayItemType.IsSet() {
		toSerialize["ArrayItemType"] = o.ArrayItemType.Get()
	}
	if o.Max != nil {
		toSerialize["Max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["Min"] = o.Min
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowArrayDataTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowArrayDataTypeAllOf := _WorkflowArrayDataTypeAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowArrayDataTypeAllOf); err == nil {
		*o = WorkflowArrayDataTypeAllOf(varWorkflowArrayDataTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayItemType")
		delete(additionalProperties, "Max")
		delete(additionalProperties, "Min")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowArrayDataTypeAllOf struct {
	value *WorkflowArrayDataTypeAllOf
	isSet bool
}

func (v NullableWorkflowArrayDataTypeAllOf) Get() *WorkflowArrayDataTypeAllOf {
	return v.value
}

func (v *NullableWorkflowArrayDataTypeAllOf) Set(val *WorkflowArrayDataTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowArrayDataTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowArrayDataTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowArrayDataTypeAllOf(val *WorkflowArrayDataTypeAllOf) *NullableWorkflowArrayDataTypeAllOf {
	return &NullableWorkflowArrayDataTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowArrayDataTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowArrayDataTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
