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

// WorkflowExpectPromptAllOf Definition of the list of properties defined in 'workflow.ExpectPrompt', excluding properties defined in parent classes.
type WorkflowExpectPromptAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The regex of the expect prompt of the interactive command.
	Expect *string `json:"Expect,omitempty"`
	// The answer string to the expect prompt.
	Send                 *string `json:"Send,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowExpectPromptAllOf WorkflowExpectPromptAllOf

// NewWorkflowExpectPromptAllOf instantiates a new WorkflowExpectPromptAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowExpectPromptAllOf(classId string, objectType string) *WorkflowExpectPromptAllOf {
	this := WorkflowExpectPromptAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowExpectPromptAllOfWithDefaults instantiates a new WorkflowExpectPromptAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowExpectPromptAllOfWithDefaults() *WorkflowExpectPromptAllOf {
	this := WorkflowExpectPromptAllOf{}
	var classId string = "workflow.ExpectPrompt"
	this.ClassId = classId
	var objectType string = "workflow.ExpectPrompt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowExpectPromptAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowExpectPromptAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowExpectPromptAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowExpectPromptAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowExpectPromptAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowExpectPromptAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExpect returns the Expect field value if set, zero value otherwise.
func (o *WorkflowExpectPromptAllOf) GetExpect() string {
	if o == nil || o.Expect == nil {
		var ret string
		return ret
	}
	return *o.Expect
}

// GetExpectOk returns a tuple with the Expect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExpectPromptAllOf) GetExpectOk() (*string, bool) {
	if o == nil || o.Expect == nil {
		return nil, false
	}
	return o.Expect, true
}

// HasExpect returns a boolean if a field has been set.
func (o *WorkflowExpectPromptAllOf) HasExpect() bool {
	if o != nil && o.Expect != nil {
		return true
	}

	return false
}

// SetExpect gets a reference to the given string and assigns it to the Expect field.
func (o *WorkflowExpectPromptAllOf) SetExpect(v string) {
	o.Expect = &v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *WorkflowExpectPromptAllOf) GetSend() string {
	if o == nil || o.Send == nil {
		var ret string
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExpectPromptAllOf) GetSendOk() (*string, bool) {
	if o == nil || o.Send == nil {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *WorkflowExpectPromptAllOf) HasSend() bool {
	if o != nil && o.Send != nil {
		return true
	}

	return false
}

// SetSend gets a reference to the given string and assigns it to the Send field.
func (o *WorkflowExpectPromptAllOf) SetSend(v string) {
	o.Send = &v
}

func (o WorkflowExpectPromptAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Expect != nil {
		toSerialize["Expect"] = o.Expect
	}
	if o.Send != nil {
		toSerialize["Send"] = o.Send
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowExpectPromptAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowExpectPromptAllOf := _WorkflowExpectPromptAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowExpectPromptAllOf); err == nil {
		*o = WorkflowExpectPromptAllOf(varWorkflowExpectPromptAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Expect")
		delete(additionalProperties, "Send")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowExpectPromptAllOf struct {
	value *WorkflowExpectPromptAllOf
	isSet bool
}

func (v NullableWorkflowExpectPromptAllOf) Get() *WorkflowExpectPromptAllOf {
	return v.value
}

func (v *NullableWorkflowExpectPromptAllOf) Set(val *WorkflowExpectPromptAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowExpectPromptAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowExpectPromptAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowExpectPromptAllOf(val *WorkflowExpectPromptAllOf) *NullableWorkflowExpectPromptAllOf {
	return &NullableWorkflowExpectPromptAllOf{value: val, isSet: true}
}

func (v NullableWorkflowExpectPromptAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowExpectPromptAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
