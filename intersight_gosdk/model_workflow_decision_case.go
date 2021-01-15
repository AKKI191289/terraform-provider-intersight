/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-06T06:42:37Z.
 *
 * API version: 1.0.9-3181
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowDecisionCase A Decision Case is a condition for a Decision Task. It is the equivalent of case statement in the programming language paradigm of switch and case.
type WorkflowDecisionCase struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of this decision case.
	Description *string `json:"Description,omitempty"`
	// The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node.
	NextTask *string `json:"NextTask,omitempty"`
	// Value for the decision case.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDecisionCase WorkflowDecisionCase

// NewWorkflowDecisionCase instantiates a new WorkflowDecisionCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDecisionCase(classId string, objectType string) *WorkflowDecisionCase {
	this := WorkflowDecisionCase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowDecisionCaseWithDefaults instantiates a new WorkflowDecisionCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDecisionCaseWithDefaults() *WorkflowDecisionCase {
	this := WorkflowDecisionCase{}
	var classId string = "workflow.DecisionCase"
	this.ClassId = classId
	var objectType string = "workflow.DecisionCase"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowDecisionCase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowDecisionCase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowDecisionCase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowDecisionCase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowDecisionCase) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCase) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowDecisionCase) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowDecisionCase) SetDescription(v string) {
	o.Description = &v
}

// GetNextTask returns the NextTask field value if set, zero value otherwise.
func (o *WorkflowDecisionCase) GetNextTask() string {
	if o == nil || o.NextTask == nil {
		var ret string
		return ret
	}
	return *o.NextTask
}

// GetNextTaskOk returns a tuple with the NextTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCase) GetNextTaskOk() (*string, bool) {
	if o == nil || o.NextTask == nil {
		return nil, false
	}
	return o.NextTask, true
}

// HasNextTask returns a boolean if a field has been set.
func (o *WorkflowDecisionCase) HasNextTask() bool {
	if o != nil && o.NextTask != nil {
		return true
	}

	return false
}

// SetNextTask gets a reference to the given string and assigns it to the NextTask field.
func (o *WorkflowDecisionCase) SetNextTask(v string) {
	o.NextTask = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkflowDecisionCase) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionCase) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowDecisionCase) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WorkflowDecisionCase) SetValue(v string) {
	o.Value = &v
}

func (o WorkflowDecisionCase) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.NextTask != nil {
		toSerialize["NextTask"] = o.NextTask
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDecisionCase) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowDecisionCaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of this decision case.
		Description *string `json:"Description,omitempty"`
		// The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node.
		NextTask *string `json:"NextTask,omitempty"`
		// Value for the decision case.
		Value *string `json:"Value,omitempty"`
	}

	varWorkflowDecisionCaseWithoutEmbeddedStruct := WorkflowDecisionCaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowDecisionCaseWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowDecisionCase := _WorkflowDecisionCase{}
		varWorkflowDecisionCase.ClassId = varWorkflowDecisionCaseWithoutEmbeddedStruct.ClassId
		varWorkflowDecisionCase.ObjectType = varWorkflowDecisionCaseWithoutEmbeddedStruct.ObjectType
		varWorkflowDecisionCase.Description = varWorkflowDecisionCaseWithoutEmbeddedStruct.Description
		varWorkflowDecisionCase.NextTask = varWorkflowDecisionCaseWithoutEmbeddedStruct.NextTask
		varWorkflowDecisionCase.Value = varWorkflowDecisionCaseWithoutEmbeddedStruct.Value
		*o = WorkflowDecisionCase(varWorkflowDecisionCase)
	} else {
		return err
	}

	varWorkflowDecisionCase := _WorkflowDecisionCase{}

	err = json.Unmarshal(bytes, &varWorkflowDecisionCase)
	if err == nil {
		o.MoBaseComplexType = varWorkflowDecisionCase.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "NextTask")
		delete(additionalProperties, "Value")

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

type NullableWorkflowDecisionCase struct {
	value *WorkflowDecisionCase
	isSet bool
}

func (v NullableWorkflowDecisionCase) Get() *WorkflowDecisionCase {
	return v.value
}

func (v *NullableWorkflowDecisionCase) Set(val *WorkflowDecisionCase) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDecisionCase) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDecisionCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDecisionCase(val *WorkflowDecisionCase) *NullableWorkflowDecisionCase {
	return &NullableWorkflowDecisionCase{value: val, isSet: true}
}

func (v NullableWorkflowDecisionCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDecisionCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
