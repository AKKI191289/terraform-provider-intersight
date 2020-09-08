/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowWaitTaskPrompt WaitTaskPrompt is used to create a customized prompts for wait control task. Wait task can be used in workflow for variety of reason, the prompts will help workflow designer to provide a customized set of prompts.
type WorkflowWaitTaskPrompt struct {
	MoBaseComplexType
	// Description that give more details about what it means to pick this prompt option for the wait task.
	Description *string `json:"Description,omitempty"`
	// User friendly label for the prompt. This label will be shown to the user as one of available options for the wait task.
	Label *string `json:"Label,omitempty"`
	// Name for the wait prompt.
	Name *string `json:"Name,omitempty"`
	// Task status for the wait task when this prompt option is selected. * `Scheduled` - The enum represents the status when task is in scheduled state. * `InProgress` - The enum represents the status when task is in-progress state. * `NoOp` - The enum represents the status when task is a noop. * `Timeout` - The enum represents the status when task has timed out. * `Completed` - The enum represents the status when task has completed. * `Failed` - The enum represents the status when task has failed.
	TaskStatus           *string `json:"TaskStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWaitTaskPrompt WorkflowWaitTaskPrompt

// NewWorkflowWaitTaskPrompt instantiates a new WorkflowWaitTaskPrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWaitTaskPrompt() *WorkflowWaitTaskPrompt {
	this := WorkflowWaitTaskPrompt{}
	var taskStatus string = "Scheduled"
	this.TaskStatus = &taskStatus
	return &this
}

// NewWorkflowWaitTaskPromptWithDefaults instantiates a new WorkflowWaitTaskPrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWaitTaskPromptWithDefaults() *WorkflowWaitTaskPrompt {
	this := WorkflowWaitTaskPrompt{}
	var taskStatus string = "Scheduled"
	this.TaskStatus = &taskStatus
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPrompt) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPrompt) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPrompt) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowWaitTaskPrompt) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPrompt) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPrompt) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPrompt) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowWaitTaskPrompt) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPrompt) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPrompt) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPrompt) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowWaitTaskPrompt) SetName(v string) {
	o.Name = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPrompt) GetTaskStatus() string {
	if o == nil || o.TaskStatus == nil {
		var ret string
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPrompt) GetTaskStatusOk() (*string, bool) {
	if o == nil || o.TaskStatus == nil {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPrompt) HasTaskStatus() bool {
	if o != nil && o.TaskStatus != nil {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given string and assigns it to the TaskStatus field.
func (o *WorkflowWaitTaskPrompt) SetTaskStatus(v string) {
	o.TaskStatus = &v
}

func (o WorkflowWaitTaskPrompt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.TaskStatus != nil {
		toSerialize["TaskStatus"] = o.TaskStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWaitTaskPrompt) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWaitTaskPromptWithoutEmbeddedStruct struct {
		// Description that give more details about what it means to pick this prompt option for the wait task.
		Description *string `json:"Description,omitempty"`
		// User friendly label for the prompt. This label will be shown to the user as one of available options for the wait task.
		Label *string `json:"Label,omitempty"`
		// Name for the wait prompt.
		Name *string `json:"Name,omitempty"`
		// Task status for the wait task when this prompt option is selected. * `Scheduled` - The enum represents the status when task is in scheduled state. * `InProgress` - The enum represents the status when task is in-progress state. * `NoOp` - The enum represents the status when task is a noop. * `Timeout` - The enum represents the status when task has timed out. * `Completed` - The enum represents the status when task has completed. * `Failed` - The enum represents the status when task has failed.
		TaskStatus *string `json:"TaskStatus,omitempty"`
	}

	varWorkflowWaitTaskPromptWithoutEmbeddedStruct := WorkflowWaitTaskPromptWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWaitTaskPromptWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWaitTaskPrompt := _WorkflowWaitTaskPrompt{}
		varWorkflowWaitTaskPrompt.Description = varWorkflowWaitTaskPromptWithoutEmbeddedStruct.Description
		varWorkflowWaitTaskPrompt.Label = varWorkflowWaitTaskPromptWithoutEmbeddedStruct.Label
		varWorkflowWaitTaskPrompt.Name = varWorkflowWaitTaskPromptWithoutEmbeddedStruct.Name
		varWorkflowWaitTaskPrompt.TaskStatus = varWorkflowWaitTaskPromptWithoutEmbeddedStruct.TaskStatus
		*o = WorkflowWaitTaskPrompt(varWorkflowWaitTaskPrompt)
	} else {
		return err
	}

	varWorkflowWaitTaskPrompt := _WorkflowWaitTaskPrompt{}

	err = json.Unmarshal(bytes, &varWorkflowWaitTaskPrompt)
	if err == nil {
		o.MoBaseComplexType = varWorkflowWaitTaskPrompt.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "TaskStatus")

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

type NullableWorkflowWaitTaskPrompt struct {
	value *WorkflowWaitTaskPrompt
	isSet bool
}

func (v NullableWorkflowWaitTaskPrompt) Get() *WorkflowWaitTaskPrompt {
	return v.value
}

func (v *NullableWorkflowWaitTaskPrompt) Set(val *WorkflowWaitTaskPrompt) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWaitTaskPrompt) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWaitTaskPrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWaitTaskPrompt(val *WorkflowWaitTaskPrompt) *NullableWorkflowWaitTaskPrompt {
	return &NullableWorkflowWaitTaskPrompt{value: val, isSet: true}
}

func (v NullableWorkflowWaitTaskPrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWaitTaskPrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
