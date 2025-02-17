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

// WorkflowParameterSetAllOf Definition of the list of properties defined in 'workflow.ParameterSet', excluding properties defined in parent classes.
type WorkflowParameterSetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The condition to be evaluated. * `eq` - Checks if the values of the two parameters are equal. * `ne` - Checks if the values of the two parameters are not equal. * `contains` - Checks if the second parameter string value is a substring of the first parameter string value. * `matchesPattern` - Checks if a string matches a regular expression.
	Condition *string `json:"Condition,omitempty"`
	// Name of the controlling entity, whose value will be used for evaluating the parameter set.
	ControlParameter *string  `json:"ControlParameter,omitempty"`
	EnableParameters []string `json:"EnableParameters,omitempty"`
	// Name for the parameter set.  Limited to 64 alphanumeric characters (upper and lower case), and special characters '-' and '_'.
	Name *string `json:"Name,omitempty"`
	// The controlling parameter will be evaluated against this 'value'.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowParameterSetAllOf WorkflowParameterSetAllOf

// NewWorkflowParameterSetAllOf instantiates a new WorkflowParameterSetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowParameterSetAllOf(classId string, objectType string) *WorkflowParameterSetAllOf {
	this := WorkflowParameterSetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var condition string = "eq"
	this.Condition = &condition
	return &this
}

// NewWorkflowParameterSetAllOfWithDefaults instantiates a new WorkflowParameterSetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowParameterSetAllOfWithDefaults() *WorkflowParameterSetAllOf {
	this := WorkflowParameterSetAllOf{}
	var classId string = "workflow.ParameterSet"
	this.ClassId = classId
	var objectType string = "workflow.ParameterSet"
	this.ObjectType = objectType
	var condition string = "eq"
	this.Condition = &condition
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowParameterSetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowParameterSetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowParameterSetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowParameterSetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *WorkflowParameterSetAllOf) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSetAllOf) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *WorkflowParameterSetAllOf) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *WorkflowParameterSetAllOf) SetCondition(v string) {
	o.Condition = &v
}

// GetControlParameter returns the ControlParameter field value if set, zero value otherwise.
func (o *WorkflowParameterSetAllOf) GetControlParameter() string {
	if o == nil || o.ControlParameter == nil {
		var ret string
		return ret
	}
	return *o.ControlParameter
}

// GetControlParameterOk returns a tuple with the ControlParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSetAllOf) GetControlParameterOk() (*string, bool) {
	if o == nil || o.ControlParameter == nil {
		return nil, false
	}
	return o.ControlParameter, true
}

// HasControlParameter returns a boolean if a field has been set.
func (o *WorkflowParameterSetAllOf) HasControlParameter() bool {
	if o != nil && o.ControlParameter != nil {
		return true
	}

	return false
}

// SetControlParameter gets a reference to the given string and assigns it to the ControlParameter field.
func (o *WorkflowParameterSetAllOf) SetControlParameter(v string) {
	o.ControlParameter = &v
}

// GetEnableParameters returns the EnableParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowParameterSetAllOf) GetEnableParameters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EnableParameters
}

// GetEnableParametersOk returns a tuple with the EnableParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowParameterSetAllOf) GetEnableParametersOk() (*[]string, bool) {
	if o == nil || o.EnableParameters == nil {
		return nil, false
	}
	return &o.EnableParameters, true
}

// HasEnableParameters returns a boolean if a field has been set.
func (o *WorkflowParameterSetAllOf) HasEnableParameters() bool {
	if o != nil && o.EnableParameters != nil {
		return true
	}

	return false
}

// SetEnableParameters gets a reference to the given []string and assigns it to the EnableParameters field.
func (o *WorkflowParameterSetAllOf) SetEnableParameters(v []string) {
	o.EnableParameters = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowParameterSetAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSetAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowParameterSetAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowParameterSetAllOf) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkflowParameterSetAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSetAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowParameterSetAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WorkflowParameterSetAllOf) SetValue(v string) {
	o.Value = &v
}

func (o WorkflowParameterSetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}
	if o.ControlParameter != nil {
		toSerialize["ControlParameter"] = o.ControlParameter
	}
	if o.EnableParameters != nil {
		toSerialize["EnableParameters"] = o.EnableParameters
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowParameterSetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowParameterSetAllOf := _WorkflowParameterSetAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowParameterSetAllOf); err == nil {
		*o = WorkflowParameterSetAllOf(varWorkflowParameterSetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Condition")
		delete(additionalProperties, "ControlParameter")
		delete(additionalProperties, "EnableParameters")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowParameterSetAllOf struct {
	value *WorkflowParameterSetAllOf
	isSet bool
}

func (v NullableWorkflowParameterSetAllOf) Get() *WorkflowParameterSetAllOf {
	return v.value
}

func (v *NullableWorkflowParameterSetAllOf) Set(val *WorkflowParameterSetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowParameterSetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowParameterSetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowParameterSetAllOf(val *WorkflowParameterSetAllOf) *NullableWorkflowParameterSetAllOf {
	return &NullableWorkflowParameterSetAllOf{value: val, isSet: true}
}

func (v NullableWorkflowParameterSetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowParameterSetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
