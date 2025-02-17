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

// WorkflowDefaultValueAllOf Definition of the list of properties defined in 'workflow.DefaultValue', excluding properties defined in parent classes.
type WorkflowDefaultValueAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.
	IsValueSet *bool `json:"IsValueSet,omitempty"`
	// Override the default value provided for the data type. When true, allow the user to enter value for the data type.
	Override *bool `json:"Override,omitempty"`
	// Default value for the data type. If default value was provided and the input was required the default value will be used as the input.
	Value                interface{} `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDefaultValueAllOf WorkflowDefaultValueAllOf

// NewWorkflowDefaultValueAllOf instantiates a new WorkflowDefaultValueAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDefaultValueAllOf(classId string, objectType string) *WorkflowDefaultValueAllOf {
	this := WorkflowDefaultValueAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowDefaultValueAllOfWithDefaults instantiates a new WorkflowDefaultValueAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDefaultValueAllOfWithDefaults() *WorkflowDefaultValueAllOf {
	this := WorkflowDefaultValueAllOf{}
	var classId string = "workflow.DefaultValue"
	this.ClassId = classId
	var objectType string = "workflow.DefaultValue"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowDefaultValueAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValueAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowDefaultValueAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowDefaultValueAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValueAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowDefaultValueAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsValueSet returns the IsValueSet field value if set, zero value otherwise.
func (o *WorkflowDefaultValueAllOf) GetIsValueSet() bool {
	if o == nil || o.IsValueSet == nil {
		var ret bool
		return ret
	}
	return *o.IsValueSet
}

// GetIsValueSetOk returns a tuple with the IsValueSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValueAllOf) GetIsValueSetOk() (*bool, bool) {
	if o == nil || o.IsValueSet == nil {
		return nil, false
	}
	return o.IsValueSet, true
}

// HasIsValueSet returns a boolean if a field has been set.
func (o *WorkflowDefaultValueAllOf) HasIsValueSet() bool {
	if o != nil && o.IsValueSet != nil {
		return true
	}

	return false
}

// SetIsValueSet gets a reference to the given bool and assigns it to the IsValueSet field.
func (o *WorkflowDefaultValueAllOf) SetIsValueSet(v bool) {
	o.IsValueSet = &v
}

// GetOverride returns the Override field value if set, zero value otherwise.
func (o *WorkflowDefaultValueAllOf) GetOverride() bool {
	if o == nil || o.Override == nil {
		var ret bool
		return ret
	}
	return *o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefaultValueAllOf) GetOverrideOk() (*bool, bool) {
	if o == nil || o.Override == nil {
		return nil, false
	}
	return o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *WorkflowDefaultValueAllOf) HasOverride() bool {
	if o != nil && o.Override != nil {
		return true
	}

	return false
}

// SetOverride gets a reference to the given bool and assigns it to the Override field.
func (o *WorkflowDefaultValueAllOf) SetOverride(v bool) {
	o.Override = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowDefaultValueAllOf) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowDefaultValueAllOf) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowDefaultValueAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *WorkflowDefaultValueAllOf) SetValue(v interface{}) {
	o.Value = v
}

func (o WorkflowDefaultValueAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsValueSet != nil {
		toSerialize["IsValueSet"] = o.IsValueSet
	}
	if o.Override != nil {
		toSerialize["Override"] = o.Override
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDefaultValueAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowDefaultValueAllOf := _WorkflowDefaultValueAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowDefaultValueAllOf); err == nil {
		*o = WorkflowDefaultValueAllOf(varWorkflowDefaultValueAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsValueSet")
		delete(additionalProperties, "Override")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowDefaultValueAllOf struct {
	value *WorkflowDefaultValueAllOf
	isSet bool
}

func (v NullableWorkflowDefaultValueAllOf) Get() *WorkflowDefaultValueAllOf {
	return v.value
}

func (v *NullableWorkflowDefaultValueAllOf) Set(val *WorkflowDefaultValueAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDefaultValueAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDefaultValueAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDefaultValueAllOf(val *WorkflowDefaultValueAllOf) *NullableWorkflowDefaultValueAllOf {
	return &NullableWorkflowDefaultValueAllOf{value: val, isSet: true}
}

func (v NullableWorkflowDefaultValueAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDefaultValueAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
